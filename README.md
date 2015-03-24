# nbagame
An API to retrieve NBA statistics from stats.nba.com.

[![GoDoc](https://godoc.org/github.com/jbowens/nbagame?status.svg)](https://godoc.org/github.com/jbowens/nbagame)

## Overview

The endpoints exposed by nba.com are not intended for public consumption, and no public documentation exists. This package attempts to wrap these endpoints in a clean, well-documented interface. I chose Go because many of the values returned by the API are inconsistent, unclear and undocumented. Forcing them to adhere to a typed API will force some consistency and clarity. Eventually, maybe the learnings and documentation from this project can be used to create an API client in a dynamic language like Python.

```go
teams, err := API.Teams.All()
if err != nil {
  panic(err)
}

for _, team := range teams {
  fmt.Printf("%s, %s - %v\n", team.Name, team.City, team.WinPercentage)
}
```
