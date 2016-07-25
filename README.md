# nbagame
A Go client to retrieve NBA statistics from stats.nba.com. The client also supports syncing data to a mysql database.

[![GoDoc](https://godoc.org/github.com/jbowens/nbagame?status.svg)](https://godoc.org/github.com/jbowens/nbagame)

![Russ](https://github.com/jbowens/nbagame/blob/master/russ.jpg)

## nbaapi.com

If you're just looking for an API for the NBA, check out the [nbaapi.com repo](https://github.com/jbowens/nbaapi.com) that uses nbagame to sync its database.

## Overview

The endpoints exposed by nba.com are not intended for public consumption, and no public documentation exists. This package attempts to wrap these endpoints in a clean, well-documented interface.

```go
teams, err := API.Teams.All()
if err != nil {
  panic(err)
}

for _, team := range teams {
  fmt.Printf("%s, %s - %v\n", team.Name, team.City, team.Wins)
}
```

## Database Syncing

NBAGame is most useful as a means to populate a MySQL database with up-to-date NBA statistics. The [nbagame/db/sync](https://godoc.org/github.com/jbowens/nbagame/db/sync) package provides a programmatic interface for syncing data. If you don't need the programmatic interface or will be using a language other than go, you can use the command-line tool in the [nbagame/cmd](https://github.com/jbowens/nbagame/tree/master/cmd) package. First, follow the directions in the [nbagame/db README](https://github.com/jbowens/nbagame/tree/master/db) to setup your MySQL database and your goose dbconf.yml configuration file. If you have permissions to create new MySQL databases and access them without credentials, you can create the database by running

```
mysql -e "CREATE DATABASE nbagame"
go get bitbucket.org/liamstask/goose/cmd/goose
goose up
```

Once your database is constructed, `go install ./cmd/...` to install the command-line utilities.

To sync data to the entire database, run the following command:

```bash
nbagamesync -season="2015-16"
```

By default, the command-line tool only loads data from the current season (except for players, which will load all historical players too). If you want to load data from a particular season, add the season flag.

```bash
nbagamesync -season="2015-16"
```

If you don't want to sync everything, specify which entities you want to sync as arguments, ex:

```
nbagamesync teams games
```

Once you've loaded the data, open a MySQL client and try querying. Here's a sample query that calculates average blocks per game by team.

```sql
SELECT teams.id, teams.name, AVG(stats.blocks) AS avg_blocks_per_game
FROM teams
LEFT JOIN team_game_stats ON teams.id = team_game_stats.team_id
LEFT JOIN stats ON team_game_stats.stats_id = stats.id
GROUP BY teams.id ORDER BY avg_blocks_per_game DESC;
```
