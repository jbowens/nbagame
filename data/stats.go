package data

// TeamStats contains a summary stats line about a team's performance. The time
// period of the stats depends on the context.
type TeamStats struct {
	TeamID           int
	TeamName         string
	TeamAbbreviation string
	TeamCity         string
	Stats
}

// PlayerStats contains a summary stats line about a player's performance. The time
// period of the stats depends on the context.
type PlayerStats struct {
	PlayerID   int
	PlayerName string
	TeamID     int
	Stats
}

// Stats contains a stat line. These are sometimes incorporated into other structs
// that provide additional context about who the stats apply to and over what
// duration.
type Stats struct {
	MinutesPlayed          string
	FieldGoalsMade         int
	FieldGoalsAttempted    int
	FieldGoalPercentage    float64
	ThreePointersMade      int
	ThreePointersAttempted int
	ThreePointPercentage   float64
	FreeThrowsMade         int
	FreeThrowsAttempted    int
	FreeThrowPercentage    float64
	OffensiveRebounds      int
	DefensiveRebounds      int
	Rebounds               int
	Assists                int
	Steals                 int
	Blocks                 int
	Turnovers              int
	PersonalFouls          int
	Points                 int
	PlusMinus              int
}
