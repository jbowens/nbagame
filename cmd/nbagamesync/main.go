package main

import (
	"flag"
	"fmt"
	"os"
	"strings"

	"github.com/jbowens/nbagame/data"
	"github.com/jbowens/nbagame/db/sync"
)

const (
	defaultDatabaseEnvironment = "development"
)

var (
	seasonFlag = flag.String("season", data.CurrentSeason.String(), "the season to sync")
)

func main() {
	flag.Parse()
	season := data.Season(*seasonFlag)

	// Figure out what we should sync based on the arguments.
	var syncTeams, syncPlayers, syncGames, syncPlays bool
	if flag.NArg() == 0 {
		// Default to syncing everything if the flag is omitted.
		syncTeams, syncPlayers, syncGames, syncPlays = true, true, true, true
	}
	for _, arg := range flag.Args() {
		switch strings.TrimSpace(strings.ToLower(arg)) {
		case "teams":
			syncTeams = true
		case "players":
			syncPlayers = true
		case "games":
			syncGames = true
		case "plays":
			syncPlays = true
		default:
			fatal(fmt.Errorf("unrecognized argument: `%s`", arg))
		}
	}

	syncer, err := sync.New(defaultDatabaseEnvironment, "./db")
	if err != nil {
		fatal(err)
	}

	if syncTeams {
		count, err := syncer.SyncAllTeams()
		if err != nil {
			fatal(err)
		}
		fmt.Println("Synced", count, "teams to the database.")
	}

	if syncPlayers {
		count, err := syncer.SyncAllPlayers()
		if err != nil {
			fatal(err)
		}
		fmt.Println("Synced", count, "players to the database.")
	}

	if syncGames {
		count, err := syncer.SyncAllGames(season)
		if err != nil {
			fatal(err)
		}
		fmt.Println("Synced", count, "games to the database.")
	}

	if syncPlays {
		count, err := syncer.SyncAllGamesPlayByPlay(season)
		if err != nil {
			fatal(err)
		}
		fmt.Println("Synced", count, "games' play-by-plays to the database.")
	}
}

func fatal(err error) {
	fmt.Fprintf(os.Stderr, "error: %s\n", err)
	os.Exit(1)
}
