package data

// TeamStats contains a summary stats line about a team's performance. The time
// period of the stats depends on the context.
type TeamStats struct {
	TeamID           int    `json:"id"`
	TeamName         string `json:"team_name"`
	TeamAbbreviation string `json:"team_abbreviation"`
	TeamCity         string `json:"team_city"`
	Stats
}

// PlayerStats contains a summary stats line about a player's performance. The time
// period of the stats depends on the context.
type PlayerStats struct {
	PlayerID   int    `json:"id"`
	PlayerName string `json:"player_name"`
	TeamID     int    `json:"team_id"`
	Stats
}

// Stats contains a stat line. These are sometimes incorporated into other structs
// that provide additional context about who the stats apply to and over what
// duration.
type Stats struct {
	MinutesPlayed          string  `json:"minutes_played"`
	FieldGoalsMade         int     `json:"field_goals_made"`
	FieldGoalsAttempted    int     `json:"field_goals_attempted"`
	FieldGoalPercentage    float64 `json:"field_goal_percentage"`
	ThreePointersMade      int     `json:"three_pointers_made"`
	ThreePointersAttempted int     `json:"three_pointers_attempted"`
	ThreePointPercentage   float64 `json:"three_point_percentage"`
	FreeThrowsMade         int     `json:"free_throws_made"`
	FreeThrowsAttempted    int     `json:"free_throws_attempted"`
	FreeThrowPercentage    float64 `json:"free_throw_percentage"`
	OffensiveRebounds      int     `json:"offensive_rebounds"`
	DefensiveRebounds      int     `json:"defensive_rebounds"`
	Rebounds               int     `json:"rebounds"`
	Assists                int     `json:"assists"`
	Steals                 int     `json:"steals"`
	Blocks                 int     `json:"blocks"`
	Turnovers              int     `json:"turnovers"`
	PersonalFouls          int     `json:"personal_fouls"`
	Points                 int     `json:"points"`
	PlusMinus              int     `json:"plus_minus"`
}
