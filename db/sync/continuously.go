package sync

import (
	"fmt"
	"os"
	"time"

	"github.com/jbowens/nbagame/data"
)

var defaultContinuousConfig = continuousSyncConfig{
	allGamesPeriod:       24 * 3 * time.Hour,
	newGamesPeriod:       6 * time.Hour,
	liveGamesPeriod:      3 * time.Minute,
	scheduledGamesPeriod: 15 * time.Minute,
	teamsPeriod:          6 * time.Hour,
	errorFn:              func(e error) {},
}

type continuousSyncConfig struct {
	allGamesPeriod       time.Duration
	newGamesPeriod       time.Duration
	liveGamesPeriod      time.Duration
	scheduledGamesPeriod time.Duration
	teamsPeriod          time.Duration
	errorFn              func(e error)
}

// ContinuousOption defines options to the Continuously function
// that change its behavior.
type ContinuousOption func(*continuousSyncConfig)

// PrintErrors returns a ContinuousOption that will cause all errors
// to be printed to standard error.
func PrintErrors() ContinuousOption {
	return func(c *continuousSyncConfig) {
		c.errorFn = func(e error) {
			fmt.Fprintf(os.Stderr, "[sync] Continuous sync error: %s\n", e)
		}
	}
}

// Continuously will continuously sync database data. Typically,
// it's invoked via a new goroutine:
//
//     go sync.Continuously(s)
//
func Continuously(s *Syncer, opts ...ContinuousOption) {
	var c continuousSyncConfig
	c = defaultContinuousConfig

	for _, opt := range opts {
		opt(&c)
	}

	allC := time.Tick(c.allGamesPeriod)
	newC := time.Tick(c.newGamesPeriod)
	liveC := time.Tick(c.liveGamesPeriod)
	scheduledC := time.Tick(c.scheduledGamesPeriod)
	teamsC := time.Tick(c.teamsPeriod)

	for {
		select {
		case <-allC:
			_, err := s.SyncAllGames(data.CurrentSeason)
			if err != nil {
				c.errorFn(err)
				continue
			}

		case <-liveC:
			var liveGames []data.GameID
			const gamesWithStatusQ = `SELECT id FROM games WHERE status = ?`
			err := s.db.DB.Select(&liveGames, gamesWithStatusQ, data.Live)
			if err != nil {
				c.errorFn(err)
				continue
			}

			_, err = s.SyncGamesWithIDs(data.CurrentSeason, liveGames)
			if err != nil {
				c.errorFn(err)
				continue
			}

		case <-scheduledC:
			var scheduledGames []data.GameID
			const gamesWithStatusQ = `SELECT id FROM games WHERE status = ?`
			err := s.db.DB.Select(&scheduledGames, gamesWithStatusQ, data.Scheduled)
			if err != nil {
				c.errorFn(err)
				continue
			}

			_, err = s.SyncGamesWithIDs(data.CurrentSeason, scheduledGames)
			if err != nil {
				c.errorFn(err)
				continue
			}

		case <-newC:
			allIDs, err := s.allGameIDs(data.CurrentSeason)
			if err != nil {
				c.errorFn(err)
				continue
			}

			const knownIDsQ = `SELECT id FROM games WHERE season = $1`
			var knownIDs []data.GameID
			err = s.db.DB.Select(&knownIDs, knownIDsQ, data.CurrentSeason)
			if err != nil {
				c.errorFn(err)
				continue
			}
			knownSet := map[data.GameID]struct{}{}
			for _, knownID := range knownIDs {
				knownSet[knownID] = struct{}{}
			}

			var newIDs []data.GameID
			for _, id := range allIDs {
				if _, ok := knownSet[id]; !ok {
					newIDs = append(newIDs, id)
				}
			}

			_, err = s.SyncGamesWithIDs(data.CurrentSeason, newIDs)
			if err != nil {
				c.errorFn(err)
				continue
			}

		case <-teamsC:
			_, err := s.SyncAllTeams()
			if err != nil {
				c.errorFn(err)
				continue
			}
		}
	}
}
