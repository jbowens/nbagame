package nbagame

import (
	"github.com/jbowens/nbagame/data"
	"github.com/jbowens/nbagame/endpoints"
)

var (
	// API is the default APIClient with default parameters.
	API APIClient
)

func init() {
	API = APIClient{
		Requester: &endpoints.DefaultRequester,
	}
	API.Teams = &Teams{client: &API}
	API.Players = &Players{client: &API}
	API.Games = &Games{client: &API}
}

// APIClient is the master API object for interating with the API.
type APIClient struct {
	Requester *endpoints.Requester
	Teams     *Teams
	Players   *Players
	Games     *Games
}

type Teams struct {
	client *APIClient
}

// All returns a slice of all the current NBA teams.
func (c *Teams) All() ([]*data.Team, error) {
	var resp endpoints.FranchiseHistoryResponse
	err := c.client.Requester.Request("franchisehistory", &endpoints.FranchiseHistoryParams{
		LeagueID: "00",
	}, &resp)
	return resp.Present(), err
}

type Players struct {
	client *APIClient
}

// All retrieves a slice of all players in the NBA in the current season.
func (c *Players) All() ([]*data.Player, error) {
	params := endpoints.CommonAllPlayersParams{
		LeagueID:            "00",
		Season:              data.CurrentSeason.String(),
		IsOnlyCurrentSeason: 1,
	}
	var resp endpoints.CommonAllPlayersResponse
	if err := c.client.Requester.Request("commonallplayers", &params, &resp); err != nil {
		return nil, err
	}
	return resp.Present(), nil
}

// Historical returns a slice of all players from all time.
func (c *Players) Historical() ([]*data.Player, error) {
	params := endpoints.CommonAllPlayersParams{
		LeagueID:            "00",
		Season:              data.CurrentSeason.String(),
		IsOnlyCurrentSeason: 0,
	}
	var resp endpoints.CommonAllPlayersResponse
	if err := c.client.Requester.Request("commonallplayers", &params, &resp); err != nil {
		return nil, err
	}
	return resp.Present(), nil
}

// Details returns detailed information about a player. It does not include
// stats about the player's performance.
func (c *Players) Details(playerID int) (*data.PlayerDetails, error) {
	var resp endpoints.CommonPlayerInfoResponse
	if err := c.client.Requester.Request("commonplayerinfo", &endpoints.CommonPlayerInfoParams{
		LeagueID: "00",
		PlayerID: playerID,
	}, &resp); err != nil {
		return nil, err
	}

	if len(resp.CommonPlayerInfo) == 0 {
		return nil, nil
	}
	return resp.CommonPlayerInfo[0].ToPlayerDetails(), nil
}

type Games struct {
	client *APIClient
}

// BoxScore returns the box score for the given game.
func (c *Games) BoxScore(gameID string) ([]*data.TeamStats, []*data.PlayerStats, error) {
	var resp endpoints.BoxScoreTraditionalResponse
	if err := c.client.Requester.Request("boxscoretraditionalv2", &endpoints.BoxScoreTraditionalParams{
		GameID:      gameID,
		Season:      data.CurrentSeason.String(),
		SeasonType:  "Regular Season",
		StartPeriod: 1,
		EndPeriod:   10,
		StartRange:  0,
		EndRange:    28800,
		RangeType:   2,
	}, &resp); err != nil {
		return nil, nil, err
	}

	if len(resp.TeamStats) == 0 {
		return nil, nil, nil
	}

	teamStats, playerStats := resp.ToData()
	return teamStats, playerStats, nil
}
