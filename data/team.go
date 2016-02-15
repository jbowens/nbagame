package data

import "fmt"

// Team represents an NBA team.
type Team struct {
	ID                 int    `json:"id" db:"id"`
	City               string `json:"city" db:"city"`
	Name               string `json:"name" db:"name"`
	StartYear          string `json:"start_year" db:"start_year"`
	EndYear            string `json:"end_year" db:"end_year"`
	Games              int    `json:"games" db:"games"`
	Wins               int    `json:"wins" db:"wins"`
	Losses             int    `json:"losses" db:"losses"`
	PlayOffAppearances int    `json:"playoff_appearances" db:"playoff_appearances"`
	DivisionTitles     int    `json:"division_titles" db:"division_titles"`
	ConferenceTitles   int    `json:"conference_titles" db:"conference_titles"`
	LeagueTitles       int    `json:"league_titles" db:"league_titles"`
}

func (t *Team) String() string {
	return fmt.Sprintf("%v - %s %s", t.ID, t.City, t.Name)
}
