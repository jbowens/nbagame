package endpoints

import (
	"github.com/jbowens/nbagame/data"
)

// PlayerShotLogParams defines parameters for a playerdashptshotlog request.
// http://stats.nba.com/stats/playerdashptshotlog?DateFrom=&DateTo=&GameSegment=&LastNGames=0&LeagueID=00&Location=&Month=0&OpponentTeamID=0&Outcome=&Period=0&PlayerID=2544&Season=2014-15&SeasonSegment=&SeasonType=Playoffs&TeamID=0&VsConference=&VsDivision=
type PlayerShotLogParams struct {
	DateFrom       string `json:"DateFrom"`
	DateTo         string `json:"DateTo"`
	GameSegment    string `json:"GameSegment"`
	LastNGames     int    `json:"LastNGames"`
	LeagueID       string `json:"LeagueID"`
	Location       string `json:"Location"`
	Month          int    `json:"Month"`
	OpponentTeamID int    `json:"OpponentTeamID"`
	Outcome        string `json:"Outcome"`
	Period         int    `json:"Period"`
	PlayerID       int    `json:"PlayerID"`
	Season         string `json:"Season"`
	SeasonSegment  string `json:"SeasonSegment"`
	SeasonType     string `json:"SeasonType"`
	TeamID         int    `json:"TeamID"`
	VsConference   string `json:"VsConference"`
	VsDivision     string `json:"VsDivision"`
}

// PlayerShotLogResponse is the type for all result sets returned by the
// 'playerdashptshotlog' resource.
type PlayerShotLogResponse struct {
	ShotLog []*PlayerShotRow `nbagame:"PtShotLog"`
}

// ToData() converts the PlayerShotLogResponse to a slice of data.Shots.
func (resp *PlayerShotLogResponse) ToData() []*data.Shot {
	var shots []*data.Shot
	for _, row := range resp.ShotLog {
		shots = append(shots, row.ToData())
	}
	return shots
}

// PlayerShotRow represents the schema returned for 'PtShotLog' result sets,
// returned from the 'playerdashptshotlog' resource.
type PlayerShotRow struct {
	GameID                  string  `nbagame:"GAME_ID"`
	Matchup                 string  `nbagame:"MATCHUP"`
	Location                string  `nbagame:"LOCATION"`
	WinOrLoss               string  `nbagame:"W"`
	FinalMargin             int     `nbagame:"FINAL_MARGIN"`
	ShotNumber              int     `nbagame:"SHOT_NUMBER"`
	Period                  int     `nbagame:"PERIOD"`
	GameClock               string  `nbagame:"GAME_CLOCK"`
	ShotClock               float64 `nbagame:"SHOT_CLOCK"`
	Dribbles                int     `nbagame:"DRIBBLES"`
	TouchTime               float64 `nbagame:"TOUCH_TIME"`
	Distance                float64 `nbagame:"SHOT_DIST"`
	PointsType              int     `nbagame:"PTS_TYPE"`
	Result                  string  `nbagame:"SHOT_RESULT"`
	ClosestDefender         string  `nbagame:"CLOSEST_DEFENDER"`
	ClosestDefenderPlayerID int     `nbagame:"CLOSEST_DEFENDER_PLAYER_ID"`
	ClosestDefenderDistance float64 `nbagame:"CLOSE_DEF_DIST"`
	FieldGoalsMade          int     `nbagame:"FGM"`
	Points                  int     `nbagame:"PTS"`
}

// ToData converts a PlayerShotRow into a data Shot struct.
func (r *PlayerShotRow) ToData() *data.Shot {
	var homeOrAway data.HomeOrAway
	if r.Location == "H" {
		homeOrAway = data.Home
	}

	var made bool
	if r.Result == "made" {
		made = true
	}

	return &data.Shot{
		GameID:                  data.GameID(r.GameID),
		Number:                  r.ShotNumber,
		Made:                    made,
		Points:                  r.Points,
		HomeOrAway:              homeOrAway,
		Period:                  r.Period,
		GameClock:               r.GameClock,
		ShotClock:               r.ShotClock,
		Dribbles:                r.Dribbles,
		TouchTimeSeconds:        r.TouchTime,
		Distance:                r.Distance,
		PointsType:              r.PointsType,
		ClosestDefender:         r.ClosestDefenderPlayerID,
		ClosestDefenderDistance: r.ClosestDefenderDistance,
	}
}
