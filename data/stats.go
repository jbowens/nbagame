package data

// TeamStats contains a summary stats line about a team's performance. The time
// period of the stats depends on the context.
type TeamStats struct {
	TeamID           int    `json:"id" db:"team_id"`
	TeamName         string `json:"team_name" db:"-"`
	TeamAbbreviation string `json:"team_abbreviation" db:"-"`
	TeamCity         string `json:"team_city" db:"-"`
	Stats
}

// TeamGameStats joins between the team, game and stats tables.
type TeamGameStats struct {
	TeamID  int    `json:"team_id" db:"team_id"`
	GameID  GameID `json:"game_id" db:"game_id"`
	StatsID int    `json:"stats_id" db:"stats_id"`
}

// PlayerStats contains a summary stats line about a player's performance. The time
// period of the stats depends on the context.
type PlayerStats struct {
	PlayerID   int    `json:"id" db:"player_id"`
	PlayerName string `json:"player_name" db:"-"`
	TeamID     int    `json:"team_id" db:"-"`
	Stats
}

// PlayerGameStats joins between the player, game and stats tables.
type PlayerGameStats struct {
	PlayerID int    `json:"player_id" db:"player_id"`
	GameID   GameID `json:"game_id" db:"game_id"`
	TeamID   int    `json:"team_id" db:"-"`
	StatsID  int    `json:"stats_id" db:"stats_id"`
}

// Stats contains a stat line. These are sometimes incorporated into other structs
// that provide additional context about who the stats apply to and over what
// duration.
type Stats struct {
	ID                     int     `json:"-" db:"id"`
	SecondsPlayed          int     `json:"seconds_played" db:"seconds_played"`
	FieldGoalsMade         int     `json:"field_goals_made" db:"field_goals_made"`
	FieldGoalsAttempted    int     `json:"field_goals_attempted" db:"field_goals_attempted"`
	FieldGoalPercentage    float64 `json:"field_goal_percentage" db:"-"`
	ThreePointersMade      int     `json:"three_pointers_made" db:"three_pointers_made"`
	ThreePointersAttempted int     `json:"three_pointers_attempted" db:"three_pointers_attempted"`
	ThreePointPercentage   float64 `json:"three_point_percentage" db:"-"`
	FreeThrowsMade         int     `json:"free_throws_made" db:"free_throws_made"`
	FreeThrowsAttempted    int     `json:"free_throws_attempted" db:"free_throws_attempted"`
	FreeThrowPercentage    float64 `json:"free_throw_percentage" db:"-"`
	OffensiveRebounds      int     `json:"offensive_rebounds" db:"offensive_rebounds"`
	DefensiveRebounds      int     `json:"defensive_rebounds" db:"defensive_rebounds"`
	Rebounds               int     `json:"rebounds" db:"-"`
	Assists                int     `json:"assists" db:"assists"`
	Steals                 int     `json:"steals" db:"steals"`
	Blocks                 int     `json:"blocks" db:"blocks"`
	Turnovers              int     `json:"turnovers" db:"turnovers"`
	PersonalFouls          int     `json:"personal_fouls" db:"personal_fouls"`
	Points                 int     `json:"points" db:"points"`
	PlusMinus              int     `json:"plus_minus" db:"plus_minus"`
}
