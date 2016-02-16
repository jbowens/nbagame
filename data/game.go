package data

import (
	"database/sql/driver"
	"errors"
	"time"
)

var gameStatusStrings = map[GameStatus]string{
	Unknown:   "UNKNOWN",
	Scheduled: "SCHEDULED",
	Live:      "LIVE",
	Final:     "FINAL",
}

// GameID holds a unique identifier for an NBA game. The identifier is unique
// across all seasons and teams.
type GameID string

func (id GameID) String() string {
	return string(id)
}

func (g GameID) Value() (driver.Value, error) {
	return string(g), nil
}

func (g *GameID) Scan(src interface{}) error {
	switch t := src.(type) {
	case string:
		*g = GameID(t)
	case []byte:
		*g = GameID(string(t))
	default:
		return errors.New("Incompatible type for GameID")
	}
	return nil
}

// Date is a wrapper around a time.Time, but only displays
// the date portion when serialized as JSON.
type Date time.Time

func (d Date) String() string {
	return time.Time(d).Format("01/02/2006")
}

func (d Date) MarshalText() ([]byte, error) {
	return []byte(d.String()), nil
}

func (d Date) Value() (driver.Value, error) {
	return time.Time(d), nil
}

func (d *Date) Scan(src interface{}) error {
	switch t := src.(type) {
	case time.Time:
		*d = Date(t)
	default:
		return errors.New("Incompatible type for Date")
	}
	return nil
}

// GameStatus indicates the status of a game.
type GameStatus int

const (
	// Unknown is used for unrecognized game status IDs.
	Unknown GameStatus = 0
	// Scheduled indicates that a game is scheduled but has not yet begun.
	Scheduled GameStatus = 1
	// Live indicates that a game is in progress.
	Live GameStatus = 2
	// Final indicates a Game's score is Final and the game has finished.
	Final GameStatus = 3
)

func (s GameStatus) String() string {
	return gameStatusStrings[s]
}

func (s GameStatus) Value() (driver.Value, error) {
	return int64(s), nil
}

// Game holds basic information about a NBA game.
type Game struct {
	ID                GameID     `json:"id,omitempty" db:"id"`
	HomeTeamID        int        `json:"home_team_id" db:"home_team_id"`
	VisitorTeamID     int        `json:"visitor_team_id" db:"visitor_team_id"`
	Season            Season     `json:"season" db:"season"`
	Status            GameStatus `json:"status" db:"status"`
	LastMeetingGameID GameID     `json:"last_meeting_game_id" db:"last_meeting_game_id"`
}

// GameDetails provides detailed information and summary of an NBA game.
type GameDetails struct {
	Game
	Date          Date          `json:"date" db:"time"`
	LengthMinutes int           `json:"length_minutes" db:"length_minutes"`
	Attendance    int           `json:"attendance" db:"attendance"`
	Officials     []*Official   `json:"officials" db:"-"`
	HomePoints    *PointSummary `json:"home_points" db:"-"`
	VisitorPoints *PointSummary `json:"visitor_points" db:"-"`
	LeadChanges   int           `json:"lead_changes" db:"lead_changes"`
	TimesTied     int           `json:"times_tied" db:"times_tied"`
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
