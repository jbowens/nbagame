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
	GameHeader  []*GameHeaderRow     `nbagame:"GameHeader"`
	LineScore   []*LineScoreboardRow `nbagame:"LineScore"`
	LastMeeting []*LastMeetingRow    `nbagame:"LastMeeting"`

	// TODO(jackson): Add support for EastConfStandingsByDay and
	// WestConfStandingsByDay result sets.
}

// LineScoreboardRow represents the schema returned for 'LineScore' result
// sets, returned from the 'scoreboardv2' resource. It contains slightly more
// information than the 'LineScore' result set for the box score summary
// endpoint.
type LineScoreboardRow struct {
	GameDateEST          string  `nbagame:"GAME_DATE_EST"`
	GameSequence         int     `nbagame:"GAME_SEQUENCE"`
	GameID               string  `nbagame:"GAME_ID"`
	TeamID               int     `nbagame:"TEAM_ID"`
	TeamAbbreviation     string  `nbagame:"TEAM_ABBREVIATION"`
	TeamCityName         string  `nbagame:"TEAM_CITY_NAME"`
	TeamWinsLosses       string  `nbagame:"TEAM_WINS_LOSSES"`
	Q1                   int     `nbagame:"PTS_QTR1"`
	Q2                   int     `nbagame:"PTS_QTR2"`
	Q3                   int     `nbagame:"PTS_QTR3"`
	Q4                   int     `nbagame:"PTS_QTR4"`
	OT1                  int     `nbagame:"PTS_OT1"`
	OT2                  int     `nbagame:"PTS_OT2"`
	OT3                  int     `nbagame:"PTS_OT3"`
	OT4                  int     `nbagame:"PTS_OT4"`
	OT5                  int     `nbagame:"PTS_OT5"`
	OT6                  int     `nbagame:"PTS_OT6"`
	OT7                  int     `nbagame:"PTS_OT7"`
	OT8                  int     `nbagame:"PTS_OT8"`
	OT9                  int     `nbagame:"PTS_OT9"`
	OT10                 int     `nbagame:"PTS_OT10"`
	Total                int     `nbagame:"PTS"`
	FieldGoalPercentage  float64 `nbagame:"PG_PCT"`
	FreeThrowPercentage  float64 `nbagame:"PT_PCT"`
	ThreePointPercentage float64 `nbagame:"FG3_PCT"`
	Assists              int     `nbagame:"AST"`
	Rebounds             int     `nbagame:"REB"`
	Turnovers            int     `nbagame:"TOV"`
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
			Playoffs:          row.ParseGameID().IsPlayoff(),
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
