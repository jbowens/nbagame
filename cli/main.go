package main

import (
	"fmt"
	"log"
	"os"

	"github.com/codegangsta/cli"
	"github.com/jbowens/nbagame/data"
	"github.com/jbowens/nbagame/db/sync"
)

const (
	defualtDatabaseEnvironment = "development"
)

var (
	logger = log.New(os.Stdout, "[sync] ", 0)
)

func env(c *cli.Context) string {
	return defualtDatabaseEnvironment
}

func newSyncer(c *cli.Context) (*sync.Syncer, error) {
	return sync.New(env(c), "../db")
}

var syncer *sync.Syncer

func before(c *cli.Context) (err error) {
	syncer, err = newSyncer(c)
	if syncer != nil {
		syncer.Logger = logger
	}
	return err
}

func main() {
	app := cli.NewApp()
	app.Name = "nbagame"
	app.Usage = "NBAGame API command line interface"

	app.Commands = []cli.Command{
		{
			Name:    "sync",
			Aliases: []string{"s"},
			Usage:   "sync a data type to the database",
			Subcommands: []cli.Command{
				{
					Name:   "games",
					Usage:  "sync all nba games for a season to the database",
					Before: before,
					Action: func(c *cli.Context) {
						count, err := syncer.SyncAllGames(data.CurrentSeason)
						if err != nil {
							fmt.Println("error syncing games: ", err)
							return
						}

						fmt.Println("Synced", count, "games to the database.")
					},
				},
				{
					Name:   "teams",
					Usage:  "sync all nba teams to the database",
					Before: before,
					Action: func(c *cli.Context) {
						count, err := syncer.SyncAllTeams()
						if err != nil {
							fmt.Println("error syncing teams: ", err)
							return
						}

						fmt.Println("Synced", count, "teams to the database.")
					},
				},
				{
					Name:   "players",
					Usage:  "sync all nba players to the database",
					Before: before,
					Action: func(c *cli.Context) {
						count, err := syncer.SyncAllPlayers()
						if err != nil {
							fmt.Println("error syncing players:", err)
							return
						}

						fmt.Println("Synced", count, "players to the database.")
					},
				},
				{
					Name:   "shots",
					Usage:  "sync all shots in a season to the database",
					Before: before,
					Action: func(c *cli.Context) {
						err := syncer.SyncAllShots()
						if err != nil {
							fmt.Println("error syncing shots:", err)
							return
						}

						fmt.Println("Synced all shots to the database.")
					},
				},
			},
		},
	}

	if err := app.Run(os.Args); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
