package endpoints

import "github.com/jbowens/nbagame/data"

// ScoreboardParams defines parameters for a ScoreboardV2 request.
// http://stats.nba.com/stats/scoreboardV2?DayOffset=0&LeagueID=00&gameDate=04%2F27%2F2015
type ScoreboardParams struct {
	LeagueID  string `json:"LeagueID"`
	DayOffset int    `json:"DayOffset"`
	GameDate  string `json:"gameDate"`
}

// ScoreboardResponse is the type for all result sets returned by the
// 'scoreboardV2' resource.
type ScoreboardResponse struct {
	GameHeader  []*GameHeaderRow  `nbagame:"GameHeader"`
	LineScore   []*LineScoreRow   `nbagame:"LineScore"`
	LastMeeting []*LastMeetingRow `nbagame:"LastMeeting"`

	// TODO(jackson): Add support for EastConfStandingsByDay and
	// WestConfStandingsByDay result sets.
}

// ToData converts a ScoreboardResponse to a slice of data.Games.
func (resp *ScoreboardResponse) ToData() (games []*data.Game, err error) {
	lastMeetingGameIDs := make(map[string]string)
	for _, row := range resp.LastMeeting {
		lastMeetingGameIDs[row.GameID] = row.LastGameID
	}

	for _, row := range resp.GameHeader {
		season, err := row.ParseSeason()
		if err != nil {
			return nil, err
		}

		games = append(games, &data.Game{
			ID:                row.ParseGameID(),
			HomeTeamID:        row.HomeTeamID,
			VisitorTeamID:     row.VisitorTeamID,
			Season:            season,
			Status:            row.ParseStatus(),
			LastMeetingGameID: data.GameID(lastMeetingGameIDs[row.GameID]),
		})
	}

	return games, err
}

// GameHeaderRow represents the schema returned for 'GameHeader' result
// sets, returned from the 'scoreboardV2' resource. It is identical to
// the GameSummaryRow, so this is just a type alias of that type.
type GameHeaderRow struct {
	GameSummaryRow
}
