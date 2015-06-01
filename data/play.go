package data

// EventType is an enum for the type of event that occurred in the game.
type EventType int

const (
	Other        EventType = 0
	Foul         EventType = 1
	FreeThrow    EventType = 2
	JumpBall     EventType = 3
	Rebound      EventType = 4
	ShotAttempt  EventType = 5
	Steal        EventType = 6
	Substitution EventType = 7
	Timeout      EventType = 8
	Turnover     EventType = 9
	Violation    EventType = 10
)

// ShotAttemptAttribute is an enum for attributes providing more details about a shot
// attempt event, for ex. was it a layup? a dunk? alleyoop? was it blocked?
type ShotAttemptAttribute int

const (
	AlleyOop     ShotAttemptAttribute = 1
	Blocked      ShotAttemptAttribute = 2
	Dunk         ShotAttemptAttribute = 3
	Fadeaway     ShotAttemptAttribute = 4
	FingerRoll   ShotAttemptAttribute = 5
	Floater      ShotAttemptAttribute = 6
	Hook         ShotAttemptAttribute = 7
	JumpShot     ShotAttemptAttribute = 8
	Layup        ShotAttemptAttribute = 9
	Missed       ShotAttemptAttribute = 10
	PutBack      ShotAttemptAttribute = 11
	PullUp       ShotAttemptAttribute = 12
	Reverse      ShotAttemptAttribute = 13
	StepBack     ShotAttemptAttribute = 14
	ThreePointer ShotAttemptAttribute = 15
	TipIn        ShotAttemptAttribute = 16
	Turnaround   ShotAttemptAttribute = 17
	WhileDriving ShotAttemptAttribute = 18
	WhileRunning ShotAttemptAttribute = 19
)

// Event describes an event that occurs within a game.
type Event struct {
	GameID          GameID
	Types           []EventType
	Period          int
	Score           *Score
	WallClock       string
	PeriodTime      string
	Descriptions    []string
	InvolvedPlayers []*PlayerDescription
	ShotAttributes  []ShotAttemptAttribute
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
	Home    int
	Visitor int
}
