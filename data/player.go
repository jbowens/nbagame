package data

// RosterStatus indicates whether a player is active on a roster.
type RosterStatus int

const (
	// Inactive indicates a player's roster status is inactive.
	Inactive RosterStatus = iota
	// Active indicates a player's roster status is active.
	Active
)

// Player holds information about an NBA player.
type Player struct {
	ID              int
	FirstName       string
	LastName        string
	RosterStatus    RosterStatus
	CareerStartYear string
	CareerEndYear   string
	PlayerCode      string
}
