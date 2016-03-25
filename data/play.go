package data

import (
	"database/sql/driver"
	"strings"
)

var (
	eventTypeToString = map[EventType]string{
		EventTypeOther:        "other",
		EventTypeMadeShot:     "made shot",
		EventTypeMissedShot:   "missed shot",
		EventTypeFreeThrow:    "free throw",
		EventTypeRebound:      "rebound",
		EventTypeTurnover:     "turnover",
		EventTypeFoul:         "foul",
		EventTypeViolation:    "violation",
		EventTypeSubstitution: "substitution",
		EventTypeTimeout:      "timeout",
		EventTypeJumpBall:     "jump ball",
		EventTypeEjection:     "ejection",
		EventTypePeriodStart:  "period start",
		EventTypePeriodEnd:    "period end",
	}
)

// EventType is an enum for the type of event that occurred in the game.
type EventType int

const (
	EventTypeOther        EventType = 0
	EventTypeMadeShot               = 1
	EventTypeMissedShot             = 2
	EventTypeFreeThrow              = 3
	EventTypeRebound                = 4
	EventTypeTurnover               = 5
	EventTypeFoul                   = 6
	EventTypeViolation              = 7
	EventTypeSubstitution           = 8
	EventTypeTimeout                = 9
	EventTypeJumpBall               = 10
	EventTypeEjection               = 11
	EventTypePeriodStart            = 12
	EventTypePeriodEnd              = 13
)

func (et EventType) String() string {
	if s, ok := eventTypeToString[et]; ok {
		return s
	}
	return "other"
}

func (et EventType) MarshalText() ([]byte, error) {
	return []byte(strings.Replace(et.String(), " ", "_", -1)), nil
}

func (et EventType) Value() (driver.Value, error) {
	b, err := et.MarshalText()
	if err != nil {
		return nil, err
	}
	return string(b), nil
}

// Event describes an event that occurs within a game.
type Event struct {
	GameID             GameID             `json:"game_id"`
	Number             int                `json:"number"`
	Type               EventType          `json:"type"`
	Period             int                `json:"period"`
	Score              *Score             `json:"score,omitempty"`
	PeriodTimeSeconds  int                `json:"period_time_secs"`
	WallClockString    string             `json:"wall_clock,omitempty"`
	Player1            *PlayerDescription `json:"player1,omitempty"`
	Player2            *PlayerDescription `json:"player2,omitempty"`
	Player3            *PlayerDescription `json:"player3,omitempty"`
	HomeDescription    *string            `json:"home_description,omitempty"`
	NeutralDescription *string            `json:"neutral_description,omitempty"`
	VisitorDescription *string            `json:"visitor_description,omitempty"`
	Shot               *Shot              `json:"shot,omitempty"`
}

type Score struct {
	Home    int `json:"home"`
	Visitor int `json:"visitor"`
}
