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

// AllPlayers returns a slice of all players.
func (c *APIClient) AllPlayers() ([]*data.Player, error) {
	params := endpoints.CommonAllPlayersParams{
		LeagueID:            "00",
		Season:              "2014-15",
		IsOnlyCurrentSeason: 0,
	}
	var resp results.CommonAllPlayersResponse
	if err := c.Requester.Request("commonallplayers", &params, &resp); err != nil {
		return nil, err
	}

	players := make([]*data.Player, len(resp.CommonAllPlayers))
	for idx, row := range resp.CommonAllPlayers {
		players[idx] = row.ToPlayer()
	}
	return players, nil
}
