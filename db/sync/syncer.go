package sync

import (
	"fmt"
	"log"

	"github.com/jbowens/nbagame"
	"github.com/jbowens/nbagame/data"
	"github.com/jbowens/nbagame/db"
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
				gameIDSet[gameID] = struct{}{}
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
				return err
			}

			if err := s.db.DB.Replace(details); err != nil {
				return err
			}
			for _, official := range details.Officials {
				officiated := &data.Officiated{GameID: id, OfficialID: official.ID}
				if err := s.db.DB.Replace(official, officiated); err != nil {
					return err
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
