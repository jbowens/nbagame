package data

var (
	eventTypeToString = map[EventType]string{
		Other:        "other",
		Foul:         "foul",
		FreeThrow:    "free_throw",
		JumpBall:     "jump_ball",
		Rebound:      "rebound",
		ShotAttempt:  "shot_attempt",
		Steal:        "steal",
		Substitution: "substitution",
		Timeout:      "timeout",
		Turnover:     "turnover",
		Violation:    "violation",
	}
	shotAttemptAttributeToString = map[ShotAttemptAttribute]string{
		Blocked:      "blocked",
		Dunk:         "dunk",
		Fadeaway:     "fadeaway",
		FingerRoll:   "finger_roll",
		Floater:      "floater",
		Hook:         "hook",
		JumpShot:     "jump_shot",
		Layup:        "layup",
		Missed:       "missed",
		PutBack:      "put_back",
		PullUp:       "pull_up",
		Reverse:      "reverse",
		StepBack:     "step_back",
		ThreePointer: "three_pointer",
		TipIn:        "tip_in",
		Turnaround:   "turnaround",
		WhileDriving: "while_driving",
		WhileRunning: "while_running",
	}
)

// EventType is an enum for the type of event that occurred in the game.
type EventType int

const (
	Other EventType = iota
	Foul
	FreeThrow
	JumpBall
	Rebound
	ShotAttempt
	Steal
	Substitution
	Timeout
	Turnover
	Violation
)

func (et EventType) String() string {
	return eventTypeToString[et]
}

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

func (saa ShotAttemptAttribute) String() string {
	return shotAttemptAttributeToString[saa]
}

// Event describes an event that occurs within a game.
// TODO(jackson): Figure out how to map to the database model.
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
