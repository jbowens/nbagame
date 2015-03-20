package nbagame

import (
	"github.com/jbowens/nbagame/data"
	"github.com/jbowens/nbagame/endpoints"
	"github.com/jbowens/nbagame/results"
)

var (
	// API is the default APIClient with default parameters.
	API = APIClient{&endpoints.DefaultRequester}
)

// APIClient is the master API object for interating with the API.
type APIClient struct {
	Requester *endpoints.Requester
}

// AllTeams returns a slice of all the current NBA teams.
func (c *APIClient) AllTeams() ([]*data.Team, error) {
	var resp results.FranchiseHistoryResponse
	err := c.Requester.Request("franchisehistory", &endpoints.FranchiseHistoryParams{
		LeagueID: "00",
	}, &resp)
	return resp.Present(), err
}

// PlayersForSeason retrieves a slice of all players in the NBA in the current
// season.
func (c *APIClient) PlayersForCurrentSeason() ([]*data.Player, error) {
	params := endpoints.CommonAllPlayersParams{
		LeagueID:            "00",
		Season:              string(data.CurrentSeason),
		IsOnlyCurrentSeason: 1,
	}
	var resp results.CommonAllPlayersResponse
	if err := c.Requester.Request("commonallplayers", &params, &resp); err != nil {
		return nil, err
	}
	return resp.Present(), nil
}

// AllPlayers returns a slice of all players from all time.
func (c *APIClient) AllPlayers() ([]*data.Player, error) {
	params := endpoints.CommonAllPlayersParams{
		LeagueID:            "00",
		Season:              string(data.CurrentSeason),
		IsOnlyCurrentSeason: 0,
	}
	var resp results.CommonAllPlayersResponse
	if err := c.Requester.Request("commonallplayers", &params, &resp); err != nil {
		return nil, err
	}
	return resp.Present(), nil
}
