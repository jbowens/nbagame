package results

// PlayerCareerStatsResponse is the type for all result sets returned by the
// 'playercareerstats' resource.
type PlayerCareerStatsResponse struct {
	SeasonTotalsRegularSeason []*SeasonStatsRow `nbagame:"SeasonTotalsRegularSeason"`
}

// SeasonStatsRow contains a set of stats for a season.
type SeasonStatsRow struct {
	PlayerID               int     `nbagame:"PLAYER_ID"`
	SeasonID               string  `nbagame:"SEASON_ID"`
	LeagueID               string  `nbagame:"LEAGUE_ID"`
	TeamID                 string  `nbagame:"TEAM_ID"`
	TeamAbbreviation       string  `nbagame:"TEAM_ABBREVIATION"`
	PlayerAge              int     `nbagame:"PLAYER_AGE"`
	GamesPlayed            int     `nbagame:"GP"`
	GamesStarted           int     `nbagame:"GS"`
	Minutes                float64 `nbagame:"MIN"`
	FieldGoalsMade         int     `nbagame:"FGM"`
	FieldGoalsAttempted    int     `nbagame:"FGA"`
	FieldGoalPercentage    float64 `nbagame:"FG_PCT"`
	ThreePointersMade      int     `nbagame:"FG3M"`
	ThreePointersAttempted int     `nbagame:"FG3A"`
	ThreePointPercentage   float64 `nbagame:FG3_PCT"`
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
