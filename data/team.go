package data

// Team represents an NBA team.
type Team struct {
	ID                 int
	City               string
	Name               string
	StartYear          string
	EndYear            string
	Games              int
	Wins               int
	Losses             int
	WinPercentage      float64
	PlayOffAppearances int
	DivisionTitles     int
	ConferenceTitles   int
	LeagueTitles       int
}
