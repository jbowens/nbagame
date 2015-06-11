package sync

import (
	"fmt"
	"log"

	"github.com/jbowens/nbagame"
	"github.com/jbowens/nbagame/db"
)

const (
	maximumConcurrentRequests = 10
)

// Syncer handles syncing data from the NBA API to a MySQL database.
type Syncer struct {
	Logger *log.Logger
	db     *db.DB
}

// New constructs a Syncer from the goose dbconf.yml configuration file. It
// takes one paramter, the name of the environment to use for the configuration.
func New(env string, dbconfLocation string) (*Syncer, error) {
	db, err := db.New(env, dbconfLocation)
	if err != nil {
		return nil, err
	}

	return &Syncer{
		db: db,
	}, nil
}

// SyncAllTeams syncs all teams to the database. Running this after teams have already
// been synced will update teams already in the database.
func (s *Syncer) SyncAllTeams() (int, error) {
	teams, err := nbagame.API.Teams.All()
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
	players, err := nbagame.API.Players.Historical()
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
