package results

import "github.com/jbowens/nbagame/data"

// FranchiseHistoryResponse is the type for all result sets returned by the
// 'franchisehistory' resource.
type FranchiseHistoryResponse struct {
	FranchiseHistory []*FranchiseHistoryRow `nbagame:"FranchiseHistory"`
}

// Present returns all of the current teams, created from the franchise history.
func (r *FranchiseHistoryResponse) Present() []*data.Team {
	var teams []*data.Team
	teamsSeen := make(map[int]struct{})

	// The first row for a given TeamID will contain the cumulative stats for all of
	// the franchise's history.
	for _, row := range r.FranchiseHistory {
		if _, ok := teamsSeen[row.TeamID]; ok {
			continue // Already seen
		}
		teams = append(teams, row.ToTeam())
		teamsSeen[row.TeamID] = struct{}{}
	}

	return teams
}

// FranchiseHistoryRow represents the schema returned for 'FranchiseHistory'
// result sets, from the 'franchisehistory' resource.
//
// Example URL:
// http://stats.nba.com/stats/franchisehistory?LeagueID=00
type FranchiseHistoryRow struct {
	LeagueID           string  `nbagame:"LEAGUE_ID"`
	TeamID             int     `nbagame:"TEAM_ID"`
	TeamCity           string  `nbagame:"TEAM_CITY"`
	TeamName           string  `nbagame:"TEAM_NAME"`
	StartYear          string  `nbagame:"START_YEAR"`
	EndYear            string  `nbagame:"END_YEAR"`
	Years              int     `nbagame:"YEARS"`
	Games              int     `nbagame:"GAMES"`
	Wins               int     `nbagame:"WINS"`
	Losses             int     `nbagame:"LOSSES"`
	WinPercentage      float64 `nbagame:"WIN_PCT"`
	PlayOffAppearances int     `nbagame:"PO_APPEARANCES"`
	DivisionTitles     int     `nbagame:"DIV_TITLES"`
	ConferenceTitles   int     `nbagame:"CONF_TITLES"`
	LeagueTitles       int     `nbagame:"LEAGUE_TITLES"`
}

// ToTeam converts the row to a Team data struct.
func (r *FranchiseHistoryRow) ToTeam() *data.Team {
	return &data.Team{
		ID:                 r.TeamID,
		City:               r.TeamCity,
		Name:               r.TeamName,
		StartYear:          r.StartYear,
		EndYear:            r.EndYear,
		Games:              r.Games,
		Wins:               r.Wins,
		Losses:             r.Losses,
		WinPercentage:      r.WinPercentage,
		PlayOffAppearances: r.PlayOffAppearances,
		DivisionTitles:     r.DivisionTitles,
		ConferenceTitles:   r.ConferenceTitles,
		LeagueTitles:       r.LeagueTitles,
	}
}
