package nbagame

import (
	"time"

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
		season:    data.CurrentSeason,
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

	season data.Season
}

// Season creates an API client for the given season.
func Season(season data.Season) *APIClient {
	api := &APIClient{
		Requester: &endpoints.DefaultRequester,
		season:    season,
	}
	api.Teams, api.Players, api.Games = &Teams{api}, &Players{api}, &Games{api}
	return api
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
	return resp.CommonPlayerInfo[0].ToPlayerDetails()
}

type Games struct {
	client *APIClient
}

// Details returns detailed information about the given game.
func (c *Games) Details(gameID string) (*data.GameDetails, error) {
	var resp endpoints.BoxScoreSummaryResponse
	if err := c.client.Requester.Request("boxscoresummaryv2", &endpoints.BoxScoreSummaryParams{
		GameID: gameID,
	}, &resp); err != nil {
		return nil, err
	}

	return resp.ToData()
}

// BoxScore returns the box score for the given game.
func (c *Games) BoxScore(gameID string) (*data.BoxScore, error) {
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
		return nil, err
	}

	if len(resp.TeamStats) == 0 {
		return nil, nil
	}

	teamStats, playerStats := resp.ToData()
	return &data.BoxScore{teamStats, playerStats}, nil
}

// PlayByPlay returns a play-by-play list of events for a game.
func (c *Games) PlayByPlay(gameID string) ([]*data.Event, error) {
	var resp endpoints.PlayByPlayResponse
	if err := c.client.Requester.Request("playbyplayv2", &endpoints.PlayByPlayParams{
		GameID:      gameID,
		Season:      data.CurrentSeason.String(),
		SeasonType:  "RegularSeason",
		StartPeriod: 1,
		EndPeriod:   10,
		StartRange:  0,
		EndRange:    55800,
		RangeType:   2,
	}, &resp); err != nil {
		return nil, err
	}

	return resp.ToData(), nil
}

// ByDate retrieves all the NBA games happening on the given date.
func (c *Games) ByDate(date time.Time) ([]*data.Game, error) {
	var resp endpoints.ScoreboardResponse
	if err := c.client.Requester.Request("scoreboardV2", &endpoints.ScoreboardParams{
		LeagueID:  "00",
		DayOffset: 0,
		GameDate:  date.Format("01/02/2006"),
	}, &resp); err != nil {
		return nil, err
	}

	return resp.ToData()
}

// PlayedBy returns the IDs of all games played by the given team so far this
// year. Unfortunately, the stats.nba.com API does not provide upcoming games.
func (c *Games) PlayedBy(teamID int) ([]data.GameID, error) {
	var resp endpoints.TeamGameLogResponse
	if err := c.client.Requester.Request("teamgamelog", &endpoints.TeamGameLogParams{
		LeagueID:   "00",
		TeamID:     teamID,
		Season:     c.client.season.String(),
		SeasonType: "Regular Season",
	}, &resp); err != nil {
		return nil, err
	}

	gameIDs := []data.GameID{}
	for _, game := range resp.TeamGameLog {
		gameIDs = append(gameIDs, data.GameID(game.GameID))
	}
	return gameIDs, nil
}
