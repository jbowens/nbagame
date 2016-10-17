package nbagame

import (
	"time"

	"github.com/jbowens/nbagame/data"
	"github.com/jbowens/nbagame/endpoints"
)

var (
	// teamIDs is a slice of all the current NBA teams' team IDs.
	teamIDs = []int{
		1610612737, 1610612738, 1610612739, 1610612740, 1610612741, 1610612742,
		1610612743, 1610612744, 1610612745, 1610612746, 1610612747, 1610612748,
		1610612749, 1610612750, 1610612751, 1610612752, 1610612753, 1610612754,
		1610612755, 1610612756, 1610612757, 1610612758, 1610612759, 1610612760,
		1610612761, 1610612762, 1610612763, 1610612764, 1610612765, 1610612766,
	}
)

var (
	DefaultClient *Client = &Client{
		requester: &endpoints.DefaultRequester,
	}
)

type Client struct {
	requester *endpoints.Requester
}

// Teams returns a slice of all the current NBA teams.
func (c *Client) Teams() ([]*data.Team, error) {
	var resp endpoints.FranchiseHistoryResponse
	err := c.requester.Request("franchisehistory", &endpoints.FranchiseHistoryParams{
		LeagueID: "00",
	}, &resp)
	return resp.Present(), err
}

// Players retrieves a slice of all players in the NBA in the provided
// season.
func (c *Client) Players(season data.Season) ([]*data.Player, error) {
	params := endpoints.CommonAllPlayersParams{
		LeagueID:            "00",
		Season:              season.String(),
		IsOnlyCurrentSeason: 1,
	}
	var resp endpoints.CommonAllPlayersResponse
	if err := c.requester.Request("commonallplayers", &params, &resp); err != nil {
		return nil, err
	}
	return resp.Present(), nil
}

// HistoricalPlayers returns a slice of all players from all time.
func (c *Client) HistoricalPlayers() ([]*data.Player, error) {
	params := endpoints.CommonAllPlayersParams{
		LeagueID:            "00",
		Season:              "2014-15", // arbitrary
		IsOnlyCurrentSeason: 0,
	}
	var resp endpoints.CommonAllPlayersResponse
	if err := c.requester.Request("commonallplayers", &params, &resp); err != nil {
		return nil, err
	}
	return resp.Present(), nil
}

// Details returns detailed information about a player. It does not include
// stats about the player's performance.
func (c *Client) PlayerDetails(playerID int) (*data.PlayerDetails, error) {
	var resp endpoints.CommonPlayerInfoResponse
	if err := c.requester.Request("commonplayerinfo", &endpoints.CommonPlayerInfoParams{
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

// Games returns the game IDs of all of the games played in the season,
// including playoff games.
func (c *Client) Games(season data.Season) ([]data.GameID, error) {
	var gameIDs []data.GameID
	for _, teamID := range teamIDs {
		teamGameIDs, err := c.GamesPlayedBy(season, teamID)
		if err != nil {
			return nil, err
		}

		gameIDs = append(gameIDs, teamGameIDs...)
	}
	return gameIDs, nil
}

// GameDetails returns detailed information about the given game.
func (c *Client) GameDetails(gameID string) (*data.GameDetails, error) {
	var resp endpoints.BoxScoreSummaryResponse
	if err := c.requester.Request("boxscoresummaryv2", &endpoints.BoxScoreSummaryParams{
		GameID: gameID,
	}, &resp); err != nil {
		return nil, err
	}
	return resp.ToData()
}

// BoxScore returns the box score for the given game.
func (c *Client) BoxScore(season data.Season, gameID string) (*data.BoxScore, error) {
	seasonType := "Regular Season"
	if data.GameID(gameID).IsPlayoff() {
		seasonType = "Playoffs"
	}

	var resp endpoints.BoxScoreTraditionalResponse
	if err := c.requester.Request("boxscoretraditionalv2", &endpoints.BoxScoreTraditionalParams{
		GameID:      gameID,
		Season:      season.String(),
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

// GamePlayByPlay returns a play-by-play list of events for a game.
func (c *Client) GamePlayByPlay(season data.Season, gameID string) ([]*data.Event, error) {
	seasonType := "Regular Season"
	if data.GameID(gameID).IsPlayoff() {
		seasonType = "Playoffs"
	}

	var resp endpoints.PlayByPlayResponse
	if err := c.requester.Request("playbyplayv2", &endpoints.PlayByPlayParams{
		GameID:      gameID,
		Season:      season.String(),
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

// GamesByDate retrieves all the NBA games happening on the given date.
func (c *Client) GamesByDate(date time.Time) ([]*data.Game, error) {
	var resp endpoints.ScoreboardResponse
	if err := c.requester.Request("scoreboardV2", &endpoints.ScoreboardParams{
		LeagueID:  "00",
		DayOffset: 0,
		GameDate:  date.Format("01/02/2006"),
	}, &resp); err != nil {
		return nil, err
	}
	return resp.ToData()
}

// GamesPlayedBy returns the IDs of all games played by the given team so far
// in the provided season. Unfortunately, the stats.nba.com API does not
// provide upcoming games.
func (c *Client) GamesPlayedBy(season data.Season, teamID int) ([]data.GameID, error) {
	gameIDSet := map[data.GameID]struct{}{}

	var resp endpoints.TeamGameLogResponse

	// Regular season games
	if err := c.requester.Request("teamgamelog", &endpoints.TeamGameLogParams{
		LeagueID:   "00",
		TeamID:     teamID,
		Season:     season.String(),
		SeasonType: "Regular Season",
	}, &resp); err != nil {
		return nil, err
	}
	for _, game := range resp.TeamGameLog {
		gid := data.GameID(game.GameID)
		gameIDSet[gid] = struct{}{}
	}

	// Playoff games
	if err := c.requester.Request("teamgamelog", &endpoints.TeamGameLogParams{
		LeagueID:   "00",
		TeamID:     teamID,
		Season:     season.String(),
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
