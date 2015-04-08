package endpoints

// TeamGameLogParams defines parameters for a TeamGameLog request.
type TeamGameLogParams struct {
	LeagueID   string `json:"LeagueID"`
	Season     string `json:"Season"`
	SeasonType string `json:"SeasonType"`
	TeamID     int    `json:"TeamID"`
}

// TeamGameLogResponse is the type for all result sets returned by the
// 'teamgamelog' resource.
type TeamGameLogResponse struct {
	TeamGameLog []*TeamGameLogRow `nbagame:"TeamGameLog"`
}

// TeamGameLogRow represents the schema returned for 'TeamGameLog' result
// sets, returned from the 'teamgamelog' resource.
//
// Example URL:
// http://stats.nba.com/stats/teamgamelog?LeagueID=00&Season=2014-15&SeasonType=Regular+Season&TeamID=1610612737
type TeamGameLogRow struct {
	TeamID                 int     `nbagame:"Team_ID"`
	GameID                 string  `nbagame:"Game_ID"`
	Date                   string  `nbagame:"GAME_DATE"`
	Matchup                string  `nbagame:"MATCHUP"`
	WinOrLoss              string  `nbagame:"WL"`
	MinutesPlayed          int     `nbagame:"MIN"`
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
	Turnovers              int     `nbagame:"TOV"`
	PersonalFouls          int     `nbagame:"PF"`
	Points                 int     `nbagame:"PTS"`
}
