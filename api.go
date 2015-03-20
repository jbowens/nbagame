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
	response, err := c.Requester.Request("commonallplayers", &params)
	if err != nil {
		return nil, err
	}

	var playerRows []*results.CommonAllPlayersRow
	if err := response.ResultSets[0].Decode(&playerRows); err != nil {
		return nil, err
	}
	players := make([]*data.Player, len(playerRows))
	for idx, row := range playerRows {
		players[idx] = row.ToPlayer()
	}

	return players, err
}
