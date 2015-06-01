package endpoints

import (
	"fmt"
	"strconv"
	"strings"
	"time"

	"github.com/jbowens/nbagame/data"
)

var (
	EastCoast *time.Location
)

var knownGameStatus = []data.GameStatus{
	data.Live,
	data.Final,
}

func init() {
	var err error
	EastCoast, err = time.LoadLocation("America/New_York")
	if err != nil {
		EastCoast = time.FixedZone("EST", 60*60*5)
	}
}

// BoxScoreSummaryParams defines parameters for a BoxScoreSummary request.
// http://stats.nba.com/stats/boxscoresummaryv2?GameID=0021401185
type BoxScoreSummaryParams struct {
	GameID string `json:"GameID"`
}

// BoxScoreSummaryResponse is the type for all result sets returned by the
// 'boxscoresummary' resource.
type BoxScoreSummaryResponse struct {
	GameSummary     []*GameSummaryRow     `nbagame:"GameSummary"`
	OtherStats      []*OtherStatsRow      `nbagame:"OtherStats"`
	Officials       []*OfficialsRow       `nbagame:"Officials"`
	InactivePlayers []*InactivePlayersRow `nbagame:"InactivePlayers"`
	GameInfo        []*GameInfoRow        `nbagame:"GameInfo"`
	LineScore       []*LineScoreRow       `nbagame:"LineScore"`
	LastMeeting     []*LastMeetingRow     `nbagame:"LastMeeting"`
	SeasonSeries    []*SeasonSeriesRow    `nbagame:"SeasonSeries"`
}

func (r *BoxScoreSummaryResponse) ToData() (*data.GameDetails, error) {
	if len(r.GameSummary) < 1 {
		return nil, ErrBadResponse("no game summary data")
	} else if len(r.LastMeeting) < 1 {
		return nil, ErrBadResponse("no last meeting data")
	} else if len(r.GameInfo) < 1 {
		return nil, ErrBadResponse("no game info data")
	} else if len(r.OtherStats) < 2 {
		return nil, ErrBadResponse("no other stats data")
	} else if len(r.LineScore) < 2 {
		return nil, ErrBadResponse("no line score data")
	}
	summary := r.GameSummary[0]

	season, err := summary.ParseSeason()
	if err != nil {
		return nil, err
	}

	gameDate, err := summary.ParseDate()
	if err != nil {
		return nil, err
	}

	var homeLineScore, visitorLineScore *LineScoreRow
	var homeOtherStats, visitorOtherStats *OtherStatsRow
	if r.LineScore[0].TeamID == summary.HomeTeamID {
		homeLineScore, visitorLineScore = r.LineScore[0], r.LineScore[1]
	} else {
		homeLineScore, visitorLineScore = r.LineScore[1], r.LineScore[0]
	}
	if r.OtherStats[0].TeamID == summary.HomeTeamID {
		homeOtherStats, visitorOtherStats = r.OtherStats[0], r.OtherStats[1]
	} else {
		homeOtherStats, visitorOtherStats = r.OtherStats[1], r.OtherStats[0]
	}

	details := &data.GameDetails{
		Game: data.Game{
			ID:                summary.ParseGameID(),
			HomeTeamID:        summary.HomeTeamID,
			VisitorTeamID:     summary.VisitorTeamID,
			Season:            season,
			Status:            summary.ParseStatus(),
			LastMeetingGameID: data.GameID(r.LastMeeting[0].LastGameID),
		},
		Date:          gameDate,
		LengthMinutes: HourMinuteStringToMinutes(r.GameInfo[0].GameTime),
		Attendance:    r.GameInfo[0].Attendance,
		HomePoints: &data.PointSummary{
			InPaint:       homeOtherStats.PointsInPaint,
			SecondChance:  homeOtherStats.SecondChancePoints,
			FromBench:     homeOtherStats.PointsFromBench,
			FirstQuarter:  homeLineScore.Q1,
			SecondQuarter: homeLineScore.Q2,
			ThirdQuarter:  homeLineScore.Q3,
			FourthQuarter: homeLineScore.Q4,
		},
		VisitorPoints: &data.PointSummary{
			InPaint:       visitorOtherStats.PointsInPaint,
			SecondChance:  visitorOtherStats.SecondChancePoints,
			FromBench:     visitorOtherStats.PointsFromBench,
			FirstQuarter:  visitorLineScore.Q1,
			SecondQuarter: visitorLineScore.Q2,
			ThirdQuarter:  visitorLineScore.Q3,
			FourthQuarter: visitorLineScore.Q4,
		},
		LeadChanges: r.OtherStats[0].LeadChanges,
		TimesTied:   r.OtherStats[0].TimesTied,
	}

	hls, vls := homeLineScore, visitorLineScore
	details.HomePoints.Overtime = nonZero(hls.OT1, hls.OT2, hls.OT3, hls.OT4,
		hls.OT5, hls.OT6, hls.OT7, hls.OT8, hls.OT9, hls.OT10)
	details.VisitorPoints.Overtime = nonZero(vls.OT1, vls.OT2, vls.OT3, vls.OT4,
		vls.OT5, vls.OT6, vls.OT7, vls.OT8, vls.OT9, vls.OT10)

	for _, row := range r.Officials {
		details.Officials = append(details.Officials, row.ToData())
	}
	return details, nil
}

// GameSummaryRow represents the schema returned for 'GameSummary' result
// sets, returned from the 'boxscoresummary' resource.
type GameSummaryRow struct {
	GameDateEST             string `nbagame:"GAME_DATE_EST"`
	GameSequence            int    `nbagame:"GAME_SEQUENCE"`
	GameID                  string `nbagame:"GAME_ID"`
	GameStatusID            int    `nbagame:"GAME_STATUS_ID"`
	GameStatusText          string `nbagame:"GAME_STATUS_TEXT"`
	GameCode                string `nbagame:"GAMECODE"`
	HomeTeamID              int    `nbagame:"HOME_TEAM_ID"`
	VisitorTeamID           int    `nbagame:"VISITOR_TEAM_ID"`
	Season                  string `nbagame:"SEASON"`
	LivePeriod              int    `nbagame:"LIVE_PERIOD"`
	LivePCTime              string `nbagame:"LIVE_PC_TIME"`
	LivePeriodTimeBroadcast string `nbagame:"LIVE_PERIOD_TIME_BCAST"`
	WHStatus                int    `nbagame:"WH_STATUS"`
	// NationalTVBroadcaster   *string `nbagame:"NATL_TV_BROADCASTER_ABBREVIATION"`
}

// ParseGameID returns a data.GameID for the game.
func (row *GameSummaryRow) ParseGameID() data.GameID {
	return data.GameID(row.GameID)
}

// ParseSeason returns a data.Season representing the season in this row.
func (row *GameSummaryRow) ParseSeason() (data.Season, error) {
	seasonYear, err := strconv.Atoi(row.Season)
	if err != nil {
		return "", ErrBadResponse(err.Error())
	}
	season := fmt.Sprintf("%s-%s", strconv.Itoa(seasonYear), strconv.Itoa(seasonYear + 1)[2:])
	return data.Season(season), nil
}

// ParseDate returns the date when the game occurred.
func (row *GameSummaryRow) ParseDate() (t time.Time, err error) {
	t, err = time.ParseInLocation("2006-01-02T15:04:05", row.GameDateEST, EastCoast)
	if err != nil {
		return t, ErrBadResponse(fmt.Sprintf("unable to parse state time: %s", err.Error()))
	}
	return t, nil
}

// ParseStatus returns the status of the game.
func (row *GameSummaryRow) ParseStatus() data.GameStatus {
	return ConvertGameStatus(row.GameStatusID)
}

// OtherStatsRow represents the schema returned for 'OtherStats' result
// sets, returned from the 'boxscoresummary' resource.
type OtherStatsRow struct {
	LeagueID           string `nbagame:"LEAGUE_ID"`
	TeamID             int    `nbagame:"TEAM_ID"`
	TeamAbbreviation   string `nbagame:"TEAM_ABBREVIATION"`
	TeamCity           string `nbagame:"TEAM_CITY"`
	PointsInPaint      int    `nbagame:"PTS_PAINT"`
	SecondChancePoints int    `nbagame:"PTS_2ND_CHANCE"`
	PointsFromBench    int    `nbagame:"PTS_FB"`
	LargestLead        int    `nbagame:"LARGEST_LEAD"`
	LeadChanges        int    `nbagame:"LEAD_CHANGES"`
	TimesTied          int    `nbagame:"TIMES_TIED"`
}

// OfficialsRow represents the schema returned for 'Officials' result
// sets, returned from the 'boxscoresummary' resource.
type OfficialsRow struct {
	OfficialID   int    `nbagame:"OFFICIAL_ID"`
	FirstName    string `nbagame:"FIRST_NAME"`
	LastName     string `nbagame:"LAST_NAME"`
	JerseyNumber string `nbagame:"JERSEY_NUM"`
}

// ToData returns a data.Official struct representing this result.
func (r *OfficialsRow) ToData() *data.Official {
	return &data.Official{
		ID:           r.OfficialID,
		FirstName:    r.FirstName,
		LastName:     r.LastName,
		JerseyNumber: strings.TrimSpace(r.JerseyNumber),
	}
}

// InactivePlayersRow represents the schema returned for 'InactivePlayers' result
// sets, returned from the 'boxscoresummary' resource.
type InactivePlayersRow struct {
	PlayerID         int    `nbagame:"PLAYER_ID"`
	FirstName        string `nbagame:"FIRST_NAME"`
	LastName         string `nbagame:"LAST_NAME"`
	JerseyNumber     string `nbagame:"JERSEY_NUM"`
	TeamID           int    `nbagame:"TEAM_ID"`
	TeamCity         string `nbagame:"TEAM_CITY"`
	TeamName         string `nbagame:"TEAM_NAME"`
	TeamAbbreviation string `nbagame:"TEAM_ABBREVIATION"`
}

// GameInfoRow represents the schema returned for 'GameInfo' result
// sets, returned from the 'boxscoresummary' resource.
type GameInfoRow struct {
	GameDate   string `nbagame:"GAME_DATE"`
	Attendance int    `nbagame:"ATTENDANCE"`
	GameTime   string `nbagame:"GAME_TIME"`
}

// LineScoreRow represents the schema returned for 'LineScore' result
// sets, returned from the 'boxscoresummary' resource.
type LineScoreRow struct {
	GameDateEST      string `nbagame:"GAME_DATE_EST"`
	GameSequence     int    `nbagame:"GAME_SEQUENCE"`
	GameID           string `nbagame:"GAME_ID"`
	TeamID           int    `nbagame:"TEAM_ID"`
	TeamAbbreviation string `nbagame:"TEAM_ABBREVIATION"`
	TeamCityName     string `nbagame:"TEAM_CITY_NAME"`
	TeamNickname     string `nbagame:"TEAM_NICKNAME"`
	TeamWinsLosses   string `nbagame:"TEAM_WINS_LOSSES"`
	Q1               int    `nbagame:"PTS_QTR1"`
	Q2               int    `nbagame:"PTS_QTR2"`
	Q3               int    `nbagame:"PTS_QTR3"`
	Q4               int    `nbagame:"PTS_QTR4"`
	OT1              int    `nbagame:"PTS_OT1"`
	OT2              int    `nbagame:"PTS_OT2"`
	OT3              int    `nbagame:"PTS_OT3"`
	OT4              int    `nbagame:"PTS_OT4"`
	OT5              int    `nbagame:"PTS_OT5"`
	OT6              int    `nbagame:"PTS_OT6"`
	OT7              int    `nbagame:"PTS_OT7"`
	OT8              int    `nbagame:"PTS_OT8"`
	OT9              int    `nbagame:"PTS_OT9"`
	OT10             int    `nbagame:"PTS_OT10"`
	Total            int    `nbagame:"PTS"`
}

// LastMeetingRow represents the schema returned for 'LastMeeting' result
// sets, returned from the 'boxscoresummary' resource.
type LastMeetingRow struct {
	GameID                          string `nbagame:"GAME_ID"`
	LastGameID                      string `nbagame:"LAST_GAME_ID"`
	LastGameDateEST                 string `nbagame:"LAST_GAME_DATE_EST"`
	LastGameHomeTeamID              int    `nbagame:"LAST_GAME_HOME_TEAM_ID"`
	LastGameHomeTeamCity            string `nbagame:"LAST_GAME_HOME_TEAM_CITY"`
	LastGameHomeTeamName            string `nbagame:"LAST_GAME_HOME_TEAM_NAME"`
	LastGameHomeTeamAbbreviation    string `nbagame:"LAST_GAME_HOME_TEAM_ABBREVIATION"`
	LastGameHomeTeamPoints          int    `nbagame:"LAST_GAME_HOME_TEAM_POINTS"`
	LastGameVisitorTeamID           int    `nbagame:"LAST_GAME_VISITOR_TEAM_ID"`
	LastGameVisitorTeamCity         string `nbagame:"LAST_GAME_VISITOR_TEAM_CITY"`
	LastGameVisitorTeamName         string `nbagame:"LAST_GAME_VISITOR_TEAM_NAME"`
	LastGameVisitorTeamAbbreviation string `nbagame:"LAST_GAME_VISITOR_TEAM_CITY1"` // wat
	LastGameVisitorTeamPoints       int    `nbagame:"LAST_GAME_VISITOR_TEAM_POINTS"`
}

// SeasonSeriesRow represents the schema returned for 'SeasonSeries' result
// sets, returned from the 'boxscoresummary' resource.
type SeasonSeriesRow struct {
	GameID         string `nbagame:"GAME_ID"`
	HomeTeamID     int    `nbagame:"HOME_TEAM_ID"`
	VisitorTeamID  int    `nbagame:"VISITOR_TEAM_ID"`
	GameDateEST    string `nbagame:"GAME_DATE_EST"`
	HomeTeamWins   int    `nbagame:"HOME_TEAM_WINS"`
	HomeTeamLosses int    `nbagame:"HOME_TEAM_LOSSES"`
	SeriesLeader   string `nbagame:"SERIES_LEADER"`
}
