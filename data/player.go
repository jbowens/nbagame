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
	ID              int
	FirstName       string
	LastName        string
	RosterStatus    RosterStatus
	CareerStartYear string
	CareerEndYear   string
	PlayerCode      string
}

// PlayerDetails contains detailed information about an NBA player.
type PlayerDetails struct {
	PlayerID         int
	FirstName        string
	LastName         string
	Birthdate        *time.Time
	School           string
	Country          string
	Height           int
	Weight           int
	SeasonExperience int
	Jersey           string
	Position         string
	RosterStatus     RosterStatus
	TeamID           int
	TeamName         string
	TeamAbbreviation string
	TeamCity         string
	CareerStartYear  string
	CareerEndYear    string
	DLeague          bool
}
