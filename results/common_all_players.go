package results

import (
	"strings"

	"github.com/jbowens/nbagame/data"
)

// CommonAllPlayersResponse is the type for all result sets returned by the
// 'commonallplayers' resource.
type CommonAllPlayersResponse struct {
	CommonAllPlayers []*CommonAllPlayersRow `nbagame:"CommonAllPlayers"`
}

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
	player := data.Player{
		ID:              row.PersonID,
		RosterStatus:    data.RosterStatus(row.RosterStatus),
		CareerStartYear: row.FromYear,
		CareerEndYear:   row.ToYear,
		PlayerCode:      row.PlayerCode,
	}

	namePieces := strings.Split(row.FullName, ",")
	if len(namePieces) < 2 {
		// Sometimes names are missing commas...
		namePieces = strings.SplitN(row.FullName, " ", 2)
	}
	if len(namePieces) < 2 {
		// If there are no commas or spaces, there's not much we can do.
		namePieces = []string{row.FullName, ""}
	}

	player.FirstName = strings.TrimSpace(namePieces[1])
	player.LastName = strings.TrimSpace(namePieces[0])
	return &player
}
