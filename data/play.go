package data

// EventType is an enum for the type of event that occurred in the game.
type EventType int

const (
	Other EventType = iota
	Foul
	FreeThrow
	Rebound
	ShotAttempt
	Steal
	Substitution
	Timeout
	Turnover
	Violation
)

// ShotAttemptAttribute is an enum for attributes providing more details about a shot
// attempt event, for ex. was it a layup? a dunk? alleyoop? was it blocked?
type ShotAttemptAttribute int

const (
	AlleyOop ShotAttemptAttribute = iota
	Blocked
	Dunk
	Fadeaway
	FingerRoll
	Floater
	Hook
	JumpShot
	Layup
	Missed
	PutBack
	PullUp
	Reverse
	StepBack
	ThreePointer
	TipIn
	Turnaround
	WhileDriving
	WhileRunning
)

// Event describes an event that occurs within a game.
type Event struct {
	GameID          GameID                 `json:"game_id"`
	Types           []EventType            `json:"types"`
	Period          int                    `json:"period"`
	Score           *Score                 `json:"score,omitempty"`
	WallClock       string                 `json:"wall_clock"`
	PeriodTime      string                 `json:"period_time"`
	Descriptions    []string               `json:"descriptions"`
	InvolvedPlayers []*PlayerDescription   `json:"involved_players,omitempty"`
	ShotAttributes  []ShotAttemptAttribute `json:"shot_attributes,omitempty"`
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
