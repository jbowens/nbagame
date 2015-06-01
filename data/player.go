package data

import "time"

// RosterStatus indicates whether a player is active on a roster.
type RosterStatus int

const (
	// Inactive indicates a player's roster status is inactive.
	Inactive RosterStatus = iota
	// Active indicates a player's roster status is active.
	Active
)

// Player holds basic, identifying information about an NBA player.
type Player struct {
	ID              int          `json:"id"`
	FirstName       string       `json:"first_name"`
	LastName        string       `json:"last_name"`
	RosterStatus    RosterStatus `json:"roster_status"`
	CareerStartYear string       `json:"career_start_year"`
	CareerEndYear   string       `json:"career_end_year"`
	PlayerCode      string       `json:"player_code"`
}

// PlayerDescription summarizes a player.
type PlayerDescription struct {
	ID     int    `json:"id"`
	Name   string `json:"name"`
	TeamID int    `json:"team_id"`
}

// PlayerDetails contains detailed information about an NBA player.
type PlayerDetails struct {
	PlayerID         int          `json:"id"`
	FirstName        string       `json:"first_name"`
	LastName         string       `json:"last_name"`
	Birthdate        *time.Time   `json:"birth_date"`
	School           string       `json:"school"`
	Country          string       `json:"country"`
	Height           int          `json:"height"`
	Weight           int          `json:"weight"`
	SeasonExperience int          `json:"season_experience"`
	Jersey           string       `json:"jersey"`
	Position         string       `json:"position"`
	RosterStatus     RosterStatus `json:"roster_status"`
	TeamID           int          `json:"team_id"`
	TeamName         string       `json:"team_name"`
	TeamAbbreviation string       `json:"team_abbreviation"`
	TeamCity         string       `json:"team_city"`
	CareerStartYear  string       `json:"career_start_year"`
	CareerEndYear    string       `json:"career_end_year"`
	DLeague          bool         `json:"dleague"`
}
