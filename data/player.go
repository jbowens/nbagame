package data

import "fmt"

// RosterStatus indicates whether a player is active on a roster.
type RosterStatus int

const (
	// Inactive indicates a player's roster status is inactive.
	Inactive RosterStatus = iota
	// Active indicates a player's roster status is active.
	Active
)

func (rs RosterStatus) MarshalText() ([]byte, error) {
	if rs == Active {
		return []byte("Active"), nil
	}
	return []byte("Inactive"), nil
}

// Player holds basic, identifying information about an NBA player.
type Player struct {
	ID              int          `json:"id" db:"id"`
	FirstName       string       `json:"first_name" db:"first_name"`
	LastName        string       `json:"last_name" db:"last_name"`
	RosterStatus    RosterStatus `json:"roster_status" db:"roster_status"`
	CareerStartYear string       `json:"career_start_year" db:"career_start"`
	CareerEndYear   string       `json:"career_end_year" db:"career_end"`
	PlayerCode      string       `json:"player_code" db:"-"`
}

func (p *Player) String() string {
	return fmt.Sprintf("%v - %s, %s", p.ID, p.LastName, p.FirstName)
}

// PlayerDescription summarizes a player.
type PlayerDescription struct {
	ID     int    `json:"id"`
	Name   string `json:"name"`
	TeamID int    `json:"team_id"`
}

// PlayerDetails contains detailed information about an NBA player.
type PlayerDetails struct {
	PlayerID         int          `json:"id" db:"id"`
	FirstName        string       `json:"first_name" db:"first_name"`
	LastName         string       `json:"last_name" db:"last_name"`
	Birthdate        *Date        `json:"birthdate,omitempty" db:"birthdate"`
	School           string       `json:"school,omitempty" db:"school"`
	Country          string       `json:"country,omitempty" db:"country"`
	Height           int          `json:"height,omitempty" db:"height"`
	Weight           int          `json:"weight,omitempty" db:"weight"`
	SeasonExperience int          `json:"season_experience" db:"season_experience"`
	Jersey           string       `json:"jersey,omitempty" db:"jersey"`
	Position         string       `json:"position,omitempty" db:"position"`
	RosterStatus     RosterStatus `json:"roster_status" db:"roster_status"`
	TeamID           int          `json:"team_id" db:"team_id"`
	TeamName         string       `json:"team_name,omitempty" db:"-"`
	TeamAbbreviation string       `json:"team_abbreviation,omitempty" db:"team_abbreviation"`
	TeamCity         string       `json:"team_city,omitempty" db:"-"`
	CareerStartYear  string       `json:"career_start_year" db:"career_start"`
	CareerEndYear    string       `json:"career_end_year" db:"career_end"`
	DLeague          bool         `json:"dleague" db:"dleague"`
}
