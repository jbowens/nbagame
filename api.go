package nbagame

import (
	"time"

	"github.com/jbowens/nbagame/data"
	"github.com/jbowens/nbagame/endpoints"
)

var (
	// API is the default APIClient with default parameters.
	API APIClient

	// teamIDs is a slice of all the current NBA teams' team IDs.
	teamIDs = []int{
		1610612737, 1610612738, 1610612739, 1610612740, 1610612741, 1610612742,
		1610612743, 1610612744, 1610612745, 1610612746, 1610612747, 1610612748,
		1610612749, 1610612750, 1610612751, 1610612752, 1610612753, 1610612754,
		1610612755, 1610612756, 1610612757, 1610612758, 1610612759, 1610612760,
		1610612761, 1610612762, 1610612763, 1610612764, 1610612765, 1610612766,
	}
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

// Season retrieves the season that this API client is configured to. Not all
// functions are constrained by the season.
func (api *APIClient) Season() data.Season {
	return api.season
}

// SetSeason sets the season that this API client is configured to. Not all
// functions are constrained by the season.
func (api *APIClient) SetSeason(season data.Season) {
	api.season = season
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
		Season:              c.client.season.String(),
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
		Season:              c.client.season.String(),
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

// All returns the game IDs of all of the games played in the season,
// including playoff games.
func (c *Games) All() ([]data.GameID, error) {
	var gameIDs []data.GameID
	for _, teamID := range teamIDs {
		teamGameIDs, err := c.PlayedBy(teamID)
		if err != nil {
			return nil, err
		}

		gameIDs = append(gameIDs, teamGameIDs...)
	}
	return gameIDs, nil
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
	seasonType := "Regular Season"
	if data.GameID(gameID).IsPlayoff() {
		seasonType = "Playoffs"
	}

	var resp endpoints.BoxScoreTraditionalResponse
	if err := c.client.Requester.Request("boxscoretraditionalv2", &endpoints.BoxScoreTraditionalParams{
		GameID:      gameID,
		Season:      c.client.season.String(),
		SeasonType:  seasonType,
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
	seasonType := "Regular Season"
	if data.GameID(gameID).IsPlayoff() {
		seasonType = "Playoffs"
	}

	var resp endpoints.PlayByPlayResponse
	if err := c.client.Requester.Request("playbyplayv2", &endpoints.PlayByPlayParams{
		GameID:      gameID,
		Season:      c.client.season.String(),
		SeasonType:  seasonType,
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
	gameIDSet := map[data.GameID]struct{}{}

	var resp endpoints.TeamGameLogResponse

	// Regular season games
	if err := c.client.Requester.Request("teamgamelog", &endpoints.TeamGameLogParams{
		LeagueID:   "00",
		TeamID:     teamID,
		Season:     c.client.season.String(),
		SeasonType: "Regular Season",
	}, &resp); err != nil {
		return nil, err
	}
	for _, game := range resp.TeamGameLog {
		gid := data.GameID(game.GameID)
		gameIDSet[gid] = struct{}{}
	}

	// Playoff games
	if err := c.client.Requester.Request("teamgamelog", &endpoints.TeamGameLogParams{
		LeagueID:   "00",
		TeamID:     teamID,
		Season:     c.client.season.String(),
		SeasonType: "Playoffs",
	}, &resp); err != nil {
		return nil, err
	}
	for _, game := range resp.TeamGameLog {
		gid := data.GameID(game.GameID)
		gameIDSet[gid] = struct{}{}
	}

	var gameIDs []data.GameID
	for gID := range gameIDSet {
		gameIDs = append(gameIDs, gID)
	}
	return gameIDs, nil
}
