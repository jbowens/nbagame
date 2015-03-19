package results

import "github.com/jbowens/nbagame/data"

// CommonAllPlayersRow represents the schema returned for 'CommonAllPlayers' result
// sets, from the 'commonallplayers' resource.
//
// Example URL:
// http://stats.nba.com/stats/commonallplayers?IsOnlyCurrentSeason=0&LeagueID=00&Season=2014-15
type CommonAllPlayersRow struct {
	PersonID     int    `nbagame:"PERSON_ID"`
	FullName     string `nbagame:"DISPLAY_LAST_COMMA_FIRST"`
	RosterStatus int    `nbagame:"ROSTERSTATUS"`
	FromYear     string `nbagame:"FROM_YEAR"`
	ToYear       string `nbagame:"TO_YEAR"`
	PlayerCode   string `nbagame:"PLAYERCODE"`
}

// ToPlayer converts a CommonAllPlayersRow to a Player data struct.
func (row *CommonAllPlayersRow) ToPlayer() *data.Player {
	return &data.Player{
		ID:              row.PersonID,
		RosterStatus:    data.RosterStatus(row.RosterStatus),
		CareerStartYear: row.FromYear,
		CareerEndYear:   row.ToYear,
		PlayerCode:      row.PlayerCode,
	}
}
