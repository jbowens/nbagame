package endpoints

// ShotChartDetailParams defines parameters for a shotchartdetail request.
// http://stats.nba.com/stats/shotchartdetail?CFID=&CFPARAMS=&ContextFilter=&ContextMeasure=FGA&DateFrom=&DateTo=&EndPeriod=10&EndRange=28800&GameSegment=&LastNGames=0&LeagueID=00&Location=&Month=0&OpponentTeamID=0&Outcome=&Period=0&PlayerID=2406&Position=&RangeType=2&RookieYear=&Season=2014-15&SeasonSegment=&SeasonType=Regular+Season&StartPeriod=1&StartRange=0&TeamID=1610612765&VsConference=&VsDivision=
type ShotChartDetailParams struct {
	CFID           string `json:"CFID"`
	CFPARAMS       string `json:"CFPARAMS"`
	ContextFilter  string `json:"ContextFilter"`
	ContextMeasure string `json:"ContextMeasure"`
	DateFrom       string `json:"DateFrom"`
	DateTo         string `json:"DateTo"`
	EndPeriod      int    `json:"EndPeriod"`
	EndRange       int    `json:"EndRange"`
	GameID         string `json:"GameID"`
	GameSegment    string `json:"GameSegment"`
	LastNGames     int    `json:"LastNGames"`
	LeagueID       string `json:"LeagueID"`
	Location       string `json:"Location"`
	Month          int    `json:"Month"`
	OpponentTeamID int    `json:"OpponentTeamID"`
	Outcome        string `json:"Outcome"`
	Period         int    `json:"Period"`
	PlayerID       int    `json:"PlayerID"`
	Position       string `json:"Position"`
	RangeType      int    `json:"RangeType"`
	RookieYear     string `json:"RookieYear"`
	Season         string `json:"Season"`
	SeasonSegment  string `json:"SeasonSegment"`
	SeasonType     string `json:"SeasonType"`
	StartPeriod    int    `json:"StartPeriod"`
	StartRange     int    `json:"StartRange"`
	TeamID         int    `json:"TeamID"`
	VsConference   string `json:"VsConference"`
	VsDivision     string `json:"VsDivision"`
}

// ShotChartDetailResponse represents the response returned by the shotchartdetail
// endpoint.
type ShotChartDetailResponse struct {
	ShotDetails []*ShotDetailRow `nbagame:"Shot_Chart_Detail"`
}

// ShotDetailRow represents the schema returned for 'Shot_Chart_Detail' result sets,
// returned from the 'shotchartdetail' resource.
type ShotDetailRow struct {
	GridType         string `nbagame:"GRID_TYPE"`
	GameID           string `nbagame:"GAME_ID"`
	GameEventID      int    `nbagame:"GAME_EVENT_ID"`
	PlayerID         int    `nbagame:"PLAYER_ID"`
	PlayerName       string `nbagame:"PLAYER_NAME"`
	TeamID           int    `nbagame:"TEAM_ID"`
	TeamName         string `nbagame:"TEAM_NAME"`
	Period           int    `nbagame:"PERIOD"`
	MinutesRemaining int    `nbagame:"MINUTES_REMAINING"`
	SecondsRemaining int    `nbagame:"SECONDS_REMAINING"`
	EventType        string `nbagame:"EVENT_TYPE"`
	ActionType       string `nbagame:"ACTION_TYPE"`
	ShotZoneBasic    string `nbagame:"SHOT_ZONE_BASIC"`
	ShotZoneArea     string `nbagame:"SHOT_ZONE_AREA"`
	ShotZoneRange    string `nbagame:"SHOT_ZONE_RANGE"`
	ShotDistance     int    `nbagame:"SHOT_DISTANCE"`
	LocationX        int    `nbagame:"LOC_X"`
	LocationY        int    `nbagame:"LOC_Y"`
	ShotAttempted    int    `nbagame:"SHOT_ATTEMPTED_FLAG"`
	ShotMade         int    `nbagame:"SHOT_MADE_FLAG"`
}
