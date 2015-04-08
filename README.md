# nbagame
An API to retrieve NBA statistics from stats.nba.com.

[![GoDoc](https://godoc.org/github.com/jbowens/nbagame?status.svg)](https://godoc.org/github.com/jbowens/nbagame)

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
