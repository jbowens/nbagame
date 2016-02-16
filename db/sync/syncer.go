package sync

import (
	"fmt"
	"log"
	"sort"
	"sync"

	"github.com/jbowens/nbagame"
	"github.com/jbowens/nbagame/data"
	"github.com/jbowens/nbagame/db"
	"github.com/jbowens/nbagame/endpoints"
)

const (
	maximumConcurrentRequests = 10
)

// Syncer handles syncing data from the NBA API to a MySQL database.
type Syncer struct {
	Logger *log.Logger
	api    *nbagame.APIClient
	db     *db.DB
}

// WithDB constructs a new Syncer from an existing DB reference.
func WithDB(db *db.DB) *Syncer {
	return &Syncer{
		db: db,
	}
}

// New constructs a Syncer from the goose dbconf.yml configuration file. It
// takes one parameter, the name of the environment to use for the configuration.
func New(env string, dbconfLocation string) (*Syncer, error) {
	db, err := db.New(env, dbconfLocation)
	if err != nil {
		return nil, err
	}

	return &Syncer{
		db: db,
	}, nil
}

// API returns the APIClient used by this syncer. This may be the default,
// so be careful with mutating the client.
func (s *Syncer) API() *nbagame.APIClient {
	// If the API to use is configured on this syncer, return that.
	if s.api != nil {
		return s.api
	}
	// Otherwise, use the default APIClient.
	return &nbagame.API
}

// SetAPI sets the nbagame APIClient to use when syncing.
func (s *Syncer) SetAPI(api *nbagame.APIClient) {
	s.api = api
}

// SyncAllTeams syncs all teams to the database. Running this after teams have already
// been synced will update teams already in the database.
func (s *Syncer) SyncAllTeams() (int, error) {
	teams, err := s.API().Teams.All()
	if err != nil {
		return 0, err
	}

	// Convert to []interface{}
	objs := make([]interface{}, len(teams))
	for i := range teams {
		objs[i] = teams[i]
		s.log("processed %s", teams[i])
	}

	return len(teams), s.db.DB.Replace(objs...)
}

// SyncAllPlayers syncs all the players to the database. Running twice will update players.
func (s *Syncer) SyncAllPlayers() (int, error) {
	players, err := s.API().Players.Historical()
	if err != nil {
		return 0, err
	}

	// Submit requests, but keeping the count under maximumConcurrentRequests
	errs := make(chan error)

	funcs := make([]func(), len(players))
	for i, p := range players {
		player := p
		funcs[i] = func() {
			playerDetails, err := nbagame.API.Players.Details(player.ID)
			if err != nil {
				s.logError(err)
				errs <- err
				return
			}

			err = s.db.DB.Replace(playerDetails)
			if err == nil {
				s.log("processed %s", player)
			}
			errs <- err
		}
	}
	go throttle(funcs, maximumConcurrentRequests)

	// Wait for all the goroutines to finish for every player. Record the last error
	// to occur.
	var retError error
	for i := 0; i < len(players); i++ {
		if err := <-errs; err != nil {
			s.logError(err)
			retError = err
		}
	}

	return len(players), retError
}

// SyncAllGames syncs all the games for the given season to the database. Running twice will update the
// games and find any new games.
func (s *Syncer) SyncAllGames(season data.Season) (int, error) {
	api := nbagame.Season(season)

	teams, err := s.API().Teams.All()
	if err != nil {
		return 0, err
	}

	// First retireve all of the game IDs for this season. Unfortunately, we have
	// to make 1 request per team to retrieve all of the games.
	gameIDSet := make(map[data.GameID]struct{})
	var mu sync.Mutex
	throttler := newThrottler(maximumConcurrentRequests)
	for _, team := range teams {
		t := team
		throttler.run(func() error {
			gameIDs, err := api.Games.PlayedBy(t.ID)
			if err != nil {
				return err
			}
			s.log("found %v games for %s %s", len(gameIDs), t.City, t.Name)

			for _, gameID := range gameIDs {
				mu.Lock()
				gameIDSet[gameID] = struct{}{}
				mu.Unlock()
			}
			return nil
		})
	}
	if err := throttler.wait(); err != nil {
		return 0, err
	}

	s.log("going to start syncing details for %v games", len(gameIDSet))

	// Now, we retrieve each individual game concurrently, and insert it into the database.
	throttler = newThrottler(maximumConcurrentRequests)
	for gameID := range gameIDSet {
		id := gameID
		throttler.run(func() error {
			details, err := api.Games.Details(string(id))
			if err != nil {
				s.log("err retrieving game details: %s", err)
				return err
			}

			if err := s.db.DB.Replace(details); err != nil {
				s.log("err recording game details: %s", err)
				return err
			}
			for _, official := range details.Officials {
				officiated := &data.Officiated{GameID: id, OfficialID: official.ID}
				if err := s.db.DB.Replace(official, officiated); err != nil {
					s.log("err saving officiated details: %s", err)
					return err
				}
			}

			// Sync the box score too
			boxscore, err := api.Games.BoxScore(string(id))
			if err != nil {
				s.log("err retrieving boxscore: %s", err)
				return err
			}
			if boxscore != nil {
				for _, ts := range boxscore.TeamStats {
					if err := s.db.RecordTeamGameStats(ts.TeamID, id, &ts.Stats); err != nil {
						s.log("err recording team game stats: %s", err)
						return err
					}
				}
				for _, ps := range boxscore.PlayerStats {
					if err := s.db.RecordPlayerGameStats(ps.PlayerID, id, ps.TeamID, &ps.Stats); err != nil {
						s.log("err recording player game stats: %s", err)
						return err
					}
				}
			}

			s.log("processed game %s from %v", id, details.Date)
			return nil
		})
	}
	if err := throttler.wait(); err != nil {
		return 0, err
	}
	return len(gameIDSet), nil
}

// SyncAllShots syncs all the shots in the entire season to the database. Running twice will
// update shots and add any shots that have occurred since the last sync.
func (s *Syncer) SyncAllShots() error {
	api := s.API()

	players, err := api.Players.All()
	if err != nil {
		return err
	}

	throttler := newThrottler(maximumConcurrentRequests)
	for _, player := range players {
		player := player
		throttler.run(func() error {
			shots, err := api.Players.Shots(player.ID)
			if err != nil {
				s.log("error: %s", err)
				return err
			}

			toInsert := make([]interface{}, len(shots))
			for idx, shot := range shots {
				shot.PlayerID = player.ID
				toInsert[idx] = shot
			}

			if err := s.db.DB.Replace(toInsert...); err != nil {
				s.log("error: %s", err)
				return err
			}
			s.log("synced shots for %s %s", player.FirstName, player.LastName)
			return nil
		})
	}
	if err := throttler.wait(); err != nil {
		return err
	}

	return nil
}

// SyncShotDetails syncs additional details about shots like the location of the player
// on the court when they took the shot. This syncing function requires that games and
// shots already be synced.
func (s *Syncer) SyncShotDetails() error {
	api := s.API()

	var playerGames []struct {
		data.PlayerGameStats
		Season string `db:"season"`
	}
	if err := s.db.DB.Select(&playerGames, "SELECT `player_game_stats`.*, games.season FROM `player_game_stats` LEFT JOIN `games` ON `player_game_stats`.`game_id` = `games`.`id`"); err != nil {
		return err
	}

	// TODO: Handle "Regular Season" vs playoffs
	throttler := newThrottler(maximumConcurrentRequests)
	for _, playerGame := range playerGames {
		playerGame := playerGame
		throttler.run(func() error {
			// Query for this player's shot chart in this game.
			var resp endpoints.ShotChartDetailResponse
			if err := api.Requester.Request("shotchartdetail", &endpoints.ShotChartDetailParams{
				ContextMeasure: "FGA",
				EndPeriod:      10,
				EndRange:       28800,
				GameID:         string(playerGame.GameID),
				LeagueID:       "00",
				PlayerID:       playerGame.PlayerID,
				Season:         playerGame.Season,
				SeasonType:     "Regular Season",
				StartPeriod:    1,
				TeamID:         playerGame.TeamID,
			}, &resp); err != nil {
				s.log("error for %v, %s: %s", playerGame.PlayerID, playerGame.GameID, err)
				return err
			}

			// Sort the shots by when they occurred in the game. This lets us uniquely determine the shot
			// by mapping it to the shot_number column.
			sort.Sort(&resp)

			for i, shotDetail := range resp.ShotDetails {
				updateQuery := "UPDATE shots SET shot_type = ?, description = ?, zone = ?, location_x = ?, location_y = ? WHERE game_id = ? AND player_id = ? AND shot_number = ?"
				res, err := s.db.DB.Exec(updateQuery, shotDetail.ActionType, shotDetail.ShotZoneBasic,
					shotDetail.ShotZoneArea, shotDetail.LocationX, shotDetail.LocationY, playerGame.GameID,
					playerGame.PlayerID, i+1)
				if err != nil {
					s.log("error: %s", err)
					return err
				}

				rowsAffected, err := res.RowsAffected()
				if err != nil {
					s.log("error: %s", err)
					return err
				}
				if rowsAffected != 1 {
					s.log("player %v, game %s, shot #%v --- %v rows affected", playerGame.PlayerID, playerGame.GameID, i+1, rowsAffected)
				}
			}

			s.log("Synced %v shot details for player %v and game %s", len(resp.ShotDetails), playerGame.PlayerID, playerGame.GameID)
			return nil
		})
	}
	if err := throttler.wait(); err != nil {
		return err
	}

	return nil
}

func (s *Syncer) logError(err error) {
	if err != nil {
		s.log("error: %s", err)
	}
}

func (s *Syncer) log(format string, args ...interface{}) {
	if s.Logger != nil {
		s.Logger.Printf("%s\n", fmt.Sprintf(format, args...))
	}
}
