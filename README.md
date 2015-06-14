# nbagame
An API to retrieve NBA statistics from stats.nba.com.

[![GoDoc](https://godoc.org/github.com/jbowens/nbagame?status.svg)](https://godoc.org/github.com/jbowens/nbagame)

![Russ](https://github.com/jbowens/nbagame/blob/master/russ.jpg)

## Overview

The endpoints exposed by nba.com are not intended for public consumption, and no public documentation exists. This package attempts to wrap these endpoints in a clean, well-documented interface.

```go
teams, err := API.Teams.All()
if err != nil {
  panic(err)
}

for _, team := range teams {
  fmt.Printf("%s, %s - %v\n", team.Name, team.City, team.WinPercentage)
}
```

## Database Syncing

NBAGame is most useful as a means to populate a MySQL database with up-to-date NBA statistics. The [nbagame/db/sync](https://godoc.org/github.com/jbowens/nbagame/db/sync) package provides a programmatic interface for syncing data. If you don't need the programmatic interface or will be using a language other than go, you can use the command-line tool in the [nbagame/cli](https://github.com/jbowens/nbagame/tree/master/cli) package. First, follow the directions in the [nbagame/db README](https://github.com/jbowens/nbagame/tree/master/db) to setup your MySQL database and your goose dbconf.yml configuration file. Running `goose up` will create the nbagame schema.

Once your database is constructed, `cd cli`. You may run any or all of the following commands to sync the data that you care about:

```bash
go run main.go teams
go run main.go players
go run main.go games
go run main.go shots
```

By default, the command-line tool only loads data from the current season (except for players, which will load all historical players too). Once you've loaded the data, open a MySQL client and try querying. Here's a sample query that calculates average blocks per game by team.

```sql
SELECT teams.id, teams.name, AVG(stats.blocks) AS avg_blocks_per_game
FROM teams
LEFT JOIN team_game_stats ON teams.id = team_game_stats.team_id
LEFT JOIN stats ON team_game_stats.stats_id = stats.id
GROUP BY teams.id ORDER BY avg_blocks_per_game DESC;
```
