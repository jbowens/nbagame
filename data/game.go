package data

import "time"

// GameID holds a unique identifier for an NBA game. The identifier is unique
// across all seasons and teams.
type GameID string

func (id GameID) String() string {
	return string(id)
}

// GameStatus indicates the status of a game.
type GameStatus int

// TODO: Identify and populate the rest of the GameStatus values.
const (
	// Unknown is used for unrecognized game status IDs.
	Unknown GameStatus = 0
	// Live indicates that a game is in progress.
	Live GameStatus = 2
	// Final indicates a Game's score is Final and the game has finished.
	Final GameStatus = 3
)

// Game holds basic information about a NBA game.
type Game struct {
	ID                GameID     `json:"id"`
	HomeTeamID        int        `json:"home_team_id"`
	VisitorTeamID     int        `json:"visitor_team_id"`
	Season            Season     `json:"season"`
	Status            GameStatus `json:"status"`
	LastMeetingGameID GameID     `json:"last_meeting_game_id"`
}

// GameDetails provides detailed information and summary of an NBA game.
type GameDetails struct {
	Game
	Date          time.Time     `json:"date"`
	LengthMinutes int           `json:"length_minutes"`
	Attendance    int           `json:"attendance"`
	Officials     []*Official   `json:"officials"`
	HomePoints    *PointSummary `json:"home_points"`
	VisitorPoints *PointSummary `json:"visitor_points"`
	LeadChanges   int           `json:"lead_changes"`
	TimesTied     int           `json:"times_tied"`
}

// PointSummary provides aggregate team point statistics.
type PointSummary struct {
	InPaint       int   `json:"in_paint"`
	SecondChance  int   `json:"second_chance"`
	FromBench     int   `json:"from_bench"`
	FirstQuarter  int   `json:"first_quarter"`
	SecondQuarter int   `json:"second_quarter"`
	ThirdQuarter  int   `json:"third_quarter"`
	FourthQuarter int   `json:"fourth_quarter"`
	Overtime      []int `json:"overtime"`
	Total         int   `json:"total"`
}
