package endpoints

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

// TeamStatsRow represents the schema returned for 'GameSummary' result
// sets, returned from the 'boxscoresummary' resource.
type GameSummaryRow struct {
	GameDateEST             string  `nbagame:"GAME_DATE_EST"`
	GameSequence            int     `nbagame:"GAME_SEQUENCE"`
	GameID                  string  `nbagame:"GAME_ID"`
	GameStatusID            int     `nbagame:"GAME_STATUS_ID"`
	GameStatusText          string  `nbagame:"GAME_STATUS_TEXT"`
	GameCode                string  `nbagame:"GAMECODE"`
	HomeTeamID              int     `nbagame:"HOME_TEAM_ID"`
	VisitorTeamID           int     `nbagame:"VISITOR_TEAM_ID"`
	Season                  string  `nbagame:"SEASON"`
	LivePeriod              int     `nbagame:"LIVE_PERIOD"`
	LivePCTime              string  `nbagame:"LIVE_PC_TIME"`
	NationalTVBroadcaster   *string `nbagame:"NATL_TV_BROADCASTER_ABBREVIATION"`
	LivePeriodTimeBroadcast string  `nbagame:"LIVE_PERIOD_TIME_BCAST"`
	WHStatus                int     `nbagame:"WH_STATUS"`
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
	PointsQ1         int    `nbagame:"PTS_QTR1"`
	PointsQ2         int    `nbagame:"PTS_QTR2"`
	PointsQ3         int    `nbagame:"PTS_QTR3"`
	PointsQ4         int    `nbagame:"PTS_QTR4"`
	PointsOT1        int    `nbagame:"PTS_OT1"`
	PointsOT2        int    `nbagame:"PTS_OT2"`
	PointsOT3        int    `nbagame:"PTS_OT3"`
	PointsOT4        int    `nbagame:"PTS_OT4"`
	PointsOT5        int    `nbagame:"PTS_OT5"`
	PointsOT6        int    `nbagame:"PTS_OT6"`
	PointsOT7        int    `nbagame:"PTS_OT7"`
	PointsOT8        int    `nbagame:"PTS_OT8"`
	PointsOT9        int    `nbagame:"PTS_OT9"`
	PointsOT10       int    `nbagame:"PTS_OT10"`
	Points           int    `nbagame:"PTS"`
}

// LastMeetingRow represents the schema returned for 'LastMeeting' result
// sets, returned from the 'boxscoresummary' resource.
type LastMeetingRow struct {
	GameID                       string `nbagame:"GAME_ID"`
	LastGameID                   string `nbagame:"LAST_GAME_ID"`
	LastGameDateEST              string `nbagame:"LAST_GAME_DATE_EST"`
	LastGameHomeTeamID           int    `nbagame:"LAST_GAME_HOME_TEAM_ID"`
	LastGameHomeTeamCity         string `nbagame:"LAST_GAME_HOME_TEAM_CITY"`
	LastGameHomeTeamName         string `nbagame:"LAST_GAME_HOME_TEAM_NAME"`
	LastGameHomeTeamAbbreviation string `nbagame:"LAST_GAME_HOME_TEAM_ABBREVIATION"`
	LastGameHomeTeamPoints       int    `nbagame:"LAST_GAME_HOME_TEAM_POINTS"`
	LastGameVisitorTeamID        int    `nbagame:"LAST_GAME_VISITOR_TEAM_ID"`
	LastGameVisitorTeamCity      string `nbagame:"LAST_GAME_VISITOR_TEAM_CITY1"` // wat
	LastGameVisitorTeamName      string `nbagame:"LAST_GAME_VISITOR_TEAM_NAME"`
	LastGameVisitorTeamPoints    int    `nbagame:"LAST_GAME_VISITOR_TEAM_POINTS"`
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
