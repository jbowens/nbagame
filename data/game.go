package data

import (
	"database/sql/driver"
	"errors"
	"strings"
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

func (id GameID) Value() (driver.Value, error) {
	return string(id), nil
}

func (id *GameID) Scan(src interface{}) error {
	switch t := src.(type) {
	case string:
		*id = GameID(t)
	case []byte:
		*id = GameID(string(t))
	default:
		return errors.New("Incompatible type for GameID")
	}
	return nil
}

// IsPlayoff returns whether or not the game is a playoff game.
func (id GameID) IsPlayoff() bool {
	// Playoff game IDs start with '004', regular season IDs with '002'.
	// ¯\_(ツ)_/¯
	return strings.HasPrefix(string(id), "004")
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

func (s GameStatus) MarshalText() ([]byte, error) {
	return []byte(s.String()), nil
}

func (s GameStatus) Value() (driver.Value, error) {
	return int64(s), nil
}

// HomeOrAway indicates whether a game was home or away with respect to a team or
// player.
type HomeOrAway bool

const (
	Home HomeOrAway = true
	Away HomeOrAway = false
)

func (h HomeOrAway) String() string {
	if h == Home {
		return "Home"
	} else {
		return "Away"
	}
}

func (h HomeOrAway) MarshalText() ([]byte, error) {
	return []byte(h.String()), nil
}

func (h HomeOrAway) Value() (driver.Value, error) {
	return bool(h), nil
}

// Game holds basic information about a NBA game.
type Game struct {
	ID                GameID     `json:"id,omitempty" db:"id"`
	Playoffs          bool       `json:"playoff,omitempty" db:"playoffs"`
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
	HomePoints    *PointSummary `json:"home_points,omitempty" db:"-"`
	VisitorPoints *PointSummary `json:"visitor_points,omitempty" db:"-"`
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
