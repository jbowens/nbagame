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
// TODO(jackson): Figure out how to map to the database model.
type Event struct {
	GameID GameID    `json:"game_id"`
	Type   EventType `json:"type"`
	Period int       `json:"period"`
	Shot   *Shot     `json:"shot,omitempty"`

	// DEPRECATED
	Score           *Score               `json:"score,omitempty"`
	WallClock       string               `json:"wall_clock"`
	Descriptions    []string             `json:"descriptions"`
	InvolvedPlayers []*PlayerDescription `json:"involved_players,omitempty"`
	Types           []EventType          `json:"types"`
	PeriodTime      string               `json:"period_time"`
}

func (e *Event) Is(typ EventType) bool {
	for _, t := range e.Types {
		if t == typ {
			return true
		}
	}
	return false
}

type Score struct {
	Home    int `json:"home"`
	Visitor int `json:"visitor"`
}
