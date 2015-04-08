package endpoints

import "github.com/jbowens/nbagame/data"

// BoxScoreTraditionalParams defines parameters for a BoxScoreTraditional request.
// http://stats.nba.com/stats/boxscoretraditionalv2?EndPeriod=10&EndRange=28800&GameID=0021401147&RangeType=2&Season=2014-15&SeasonType=Regular+Season&StartPeriod=1&StartRange=0
type BoxScoreTraditionalParams struct {
	GameID      string `json:"GameID"`
	Season      string `json:"Season"`
	SeasonType  string `json:"SeasonType"`
	StartPeriod int    `json:"StartPeriod"`
	EndPeriod   int    `json:"EndPeriod"`

	// What the fuck are these?
	StartRange int `json:"StartRange"`
	EndRange   int `json:"EndRange"`
	RangeType  int `json:"RangeType"`
}

// BoxScoreTraditionalResponse is the type for all result sets returned by the
// 'boxscore' resource.
type BoxScoreTraditionalResponse struct {
	PlayerStats []*PlayerStatsRow `nbagame:"PlayerStats"`
	TeamStats   []*TeamStatsRow   `nbagame:"TeamStats"`
}

// ToData returns a nbagame.data representation of this response.
func (resp *BoxScoreTraditionalResponse) ToData() ([]*data.TeamStats, []*data.PlayerStats) {
	var teamStats []*data.TeamStats
	for _, r := range resp.TeamStats {
		teamStats = append(teamStats, r.ToTeamStats())
	}

	var playerStats []*data.PlayerStats
	for _, r := range resp.PlayerStats {
		playerStats = append(playerStats, r.ToPlayerStats())
	}
	return teamStats, playerStats
}

// PlayerStatsRow represents the schema returned for 'PlayerStats' result
// sets, returned from the 'boxscore' resource.
type PlayerStatsRow struct {
	GameID           string `nbagame:"GAME_ID"`
	TeamID           int    `nbagame:"TEAM_ID"`
	TeamAbbreviation string `nbagame:"TEAM_ABBREVIATION"`
	TeamCity         string `nbagame:"TEAM_CITY"`
	PlayerID         int    `nbagame:"PLAYER_ID"`
	PlayerName       string `nbagame:"PLAYER_NAME"`
	StartPosition    string `nbagame:"START_POSITION"`
	Comment          string `nbagame:"COMMENT"`
	StatLine
}

// ToPlayerStats converts this row into a PlayerStats data struct.
func (r *PlayerStatsRow) ToPlayerStats() *data.PlayerStats {
	return &data.PlayerStats{
		PlayerID:   r.PlayerID,
		PlayerName: r.PlayerName,
		TeamID:     r.TeamID,
		Stats:      *r.StatLine.ToStats(),
	}
}

// TeamStatsRow represents the schema returned for 'TeamStats' result
// sets, returned from the 'boxscore' resource.
type TeamStatsRow struct {
	GameID           string `nbagame:"GAME_ID"`
	TeamID           int    `nbagame:"TEAM_ID"`
	TeamName         string `nbagame:"TEAM_NAME"`
	TeamAbbreviation string `nbagame:"TEAM_ABBREVIATION"`
	TeamCity         string `nbagame:"TEAM_CITY"`
	StatLine
}

// ToTeamStats converts this row into a TeamStats data struct.
func (r *TeamStatsRow) ToTeamStats() *data.TeamStats {
	return &data.TeamStats{
		TeamID:           r.TeamID,
		TeamName:         r.TeamName,
		TeamAbbreviation: r.TeamAbbreviation,
		TeamCity:         r.TeamCity,
		Stats:            *r.StatLine.ToStats(),
	}
}

// StatLine contains a summary of statistics.
type StatLine struct {
	MinutesPlayed          string  `nbagame:"MIN"`
	FieldGoalsMade         int     `nbagame:"FGM"`
	FieldGoalsAttempted    int     `nbagame:"FGA"`
	FieldGoalPercentage    float64 `nbagame:"FG_PCT"`
	ThreePointersMade      int     `nbagame:"FG3M"`
	ThreePointersAttempted int     `nbagame:"FG3A"`
	ThreePointPercentage   float64 `nbagame:"FG3_PCT"`
	FreeThrowsMade         int     `nbagame:"FTM"`
	FreeThrowsAttempted    int     `nbagame:"FTA"`
	FreeThrowPercentage    float64 `nbagame:"FT_PCT"`
	OffensiveRebounds      int     `nbagame:"OREB"`
	DefensiveRebounds      int     `nbagame:"DREB"`
	Rebounds               int     `nbagame:"REB"`
	Assists                int     `nbagame:"AST"`
	Steals                 int     `nbagame:"STL"`
	Blocks                 int     `nbagame:"BLK"`
	Turnovers              int     `nbagame:"TO"`
	PersonalFouls          int     `nbagame:"PF"`
	Points                 int     `nbagame:"PTS"`
	PlusMinus              float64 `nbagame:"PLUS_MINUS"`
}

// ToStats converts a StatLine into a data Stats struct.
func (sl *StatLine) ToStats() *data.Stats {
	return &data.Stats{
		MinutesPlayed:          sl.MinutesPlayed,
		FieldGoalsMade:         sl.FieldGoalsMade,
		FieldGoalsAttempted:    sl.FieldGoalsAttempted,
		FieldGoalPercentage:    sl.FieldGoalPercentage,
		ThreePointersMade:      sl.ThreePointersMade,
		ThreePointersAttempted: sl.ThreePointersAttempted,
		ThreePointPercentage:   sl.ThreePointPercentage,
		FreeThrowsMade:         sl.FreeThrowsMade,
		FreeThrowsAttempted:    sl.FreeThrowsAttempted,
		FreeThrowPercentage:    sl.FreeThrowPercentage,
		OffensiveRebounds:      sl.OffensiveRebounds,
		DefensiveRebounds:      sl.DefensiveRebounds,
		Rebounds:               sl.Rebounds,
		Assists:                sl.Assists,
		Steals:                 sl.Steals,
		Blocks:                 sl.Blocks,
		Turnovers:              sl.Turnovers,
		PersonalFouls:          sl.PersonalFouls,
		Points:                 sl.Points,
		PlusMinus:              sl.PlusMinus,
	}
}
