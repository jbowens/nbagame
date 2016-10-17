package sync

import (
	"fmt"
	"log"
	"sync"

	"github.com/jbowens/nbagame"
	"github.com/jbowens/nbagame/data"
	"github.com/jbowens/nbagame/db"
)

const (
	maximumConcurrentRequests = 20
)

// Syncer handles syncing data from the NBA API to a MySQL database.
type Syncer struct {
	Logger *log.Logger
	api    *nbagame.Client
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

// Client returns the Client used by this syncer. This may be the default,
// so be careful with mutating the client.
func (s *Syncer) Client() *nbagame.Client {
	// If the API to use is configured on this syncer, return that.
	if s.api != nil {
		return s.api
	}
	// Otherwise, use the DefaultClient.
	return nbagame.DefaultClient
}

// SyncAllTeams syncs all teams to the database. Running this after teams have already
// been synced will update teams already in the database.
func (s *Syncer) SyncAllTeams() (int, error) {
	teams, err := s.Client().Teams()
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
	players, err := s.Client().HistoricalPlayers()
	if err != nil {
		return 0, err
	}

	// Submit requests, but keeping the count under maximumConcurrentRequests
	errs := make(chan error)

	funcs := make([]func(), len(players))
	for i, p := range players {
		player := p
		funcs[i] = func() {
			playerDetails, err := s.Client().PlayerDetails(player.ID)
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

func (s *Syncer) allGameIDs(season data.Season) ([]data.GameID, error) {
	teams, err := s.Client().Teams()
	if err != nil {
		return nil, err
	}

	// First retireve all of the game IDs for this season. Unfortunately, we have
	// to make 1 request per team to retrieve all of the games.
	gameIDSet := make(map[data.GameID]struct{})
	var mu sync.Mutex
	throttler := newThrottler(maximumConcurrentRequests)
	for _, team := range teams {
		t := team
		throttler.run(func() error {
			gameIDs, err := s.Client().GamesPlayedBy(season, t.ID)
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
		return nil, err
	}

	var gameIDs []data.GameID
	for gID := range gameIDSet {
		gameIDs = append(gameIDs, gID)
	}
	return gameIDs, nil
}

// SyncGamesWithIDs syncs all games with the provided game IDs.
func (s *Syncer) SyncGamesWithIDs(season data.Season, gameIDs []data.GameID) (int, error) {
	s.log("going to start syncing details for %v games", len(gameIDs))

	// Now, we retrieve each individual game concurrently, and insert it into the database.
	throttler := newThrottler(maximumConcurrentRequests)
	for _, gameID := range gameIDs {
		id := gameID
		throttler.run(func() error {
			details, err := s.Client().GameDetails(string(id))
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
			if details.Status == data.Final {
				// Box score is only available after the game :(
				boxscore, err := s.Client().BoxScore(season, string(id))
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
			}

			s.log("processed game %s from %v", id, details.Date)
			return nil
		})
	}
	if err := throttler.wait(); err != nil {
		return 0, err
	}
	return len(gameIDs), nil
}

// SyncPlaysForGames syncs play-by-play histories for games
// with the provided game IDs.
func (s *Syncer) SyncPlaysForGames(season data.Season, gameIDs []data.GameID) (int, error) {
	s.log("going to start syncing play-by-play for %v games", len(gameIDs))

	// Now, we retrieve each individual game concurrently, and insert it into the database.
	throttler := newThrottler(maximumConcurrentRequests)
	for _, gameID := range gameIDs {
		id := gameID
		throttler.run(func() error {
			events, err := s.Client().GamePlayByPlay(season, string(id))
			if err != nil {
				s.log("err retrieving game play-by-play: %s", err)
				return err
			}

			for _, evt := range events {
				err := s.db.RecordGameEvent(id, evt)
				if err != nil {
					s.log("err recording game event: %s", err)
					return err
				}
			}
			return nil
		})
	}
	if err := throttler.wait(); err != nil {
		return 0, err
	}
	return len(gameIDs), nil
}

// SyncAllGames syncs all the games for the given season to the database.
// Running twice will update the games and find any new games. Note that
// this function does not optimize and try to predict which data may need
// updating. It will re-fetch all games, including games that may have
// happened several months ago.
func (s *Syncer) SyncAllGames(season data.Season) (int, error) {
	gameIDs, err := s.allGameIDs(season)
	if err != nil {
		return 0, err
	}
	return s.SyncGamesWithIDs(season, gameIDs)
}

func (s *Syncer) SyncAllGamesPlayByPlay(season data.Season) (int, error) {
	gameIDs, err := s.allGameIDs(season)
	if err != nil {
		return 0, err
	}
	return s.SyncPlaysForGames(season, gameIDs)
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
