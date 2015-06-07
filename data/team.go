package data

// Team represents an NBA team.
type Team struct {
	ID                 int     `json:"id"`
	City               string  `json:"city"`
	Name               string  `json:"name"`
	StartYear          string  `json:"start_year"`
	EndYear            string  `json:"end_year"`
	Games              int     `json:"games"`
	Wins               int     `json:"wins"`
	Losses             int     `json:"losses"`
	WinPercentage      float64 `json:"win_percentage"`
	PlayOffAppearances int     `json:"playoff_appearances"`
	DivisionTitles     int     `json:"division_titles"`
	ConferenceTitles   int     `json:"conference_titles"`
	LeagueTitles       int     `json:"league_titles"`
}
