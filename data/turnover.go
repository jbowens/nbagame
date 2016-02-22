package data

// TurnoverType defines the different types of turnovers.
type TurnoverType int

// TODO(jackson): Fill in the gaps in the table below.
const (
	UnknownTurnover              TurnoverType = 0
	BadPassTurnover                           = 1
	LostBallTurnover                          = 2
	OutOfBoundsTurnover                       = 3
	TravelingTurnover                         = 4
	FoulTurnover                              = 5
	DoubleDribbleTurnover                     = 6
	DiscontinueDribbleTurnover                = 7
	ThreeSecondViolationTurnover              = 8
	FiveSecondViolationTurnover               = 9
	EightSecondViolationTurnover              = 10
	ShotClockTurnover                         = 11
	InboundTurnover                           = 12
	BackcourtTurnover                         = 13
	OffensiveGoaltendingTurnover              = 15
	LaneViolationTurnover                     = 17
	JumpBallViolationTurnover                 = 18
	KickedBallTurnover                        = 19
	IllegalAssistTurnover                     = 20
	PalmingTurnover                           = 21
	PunchedBallTurnover                       = 33
	SwiningElbowsTurnover                     = 34
	BasketFromBelowTurnover                   = 35
	IllegalScreenTurnover                     = 36
	OffensiveFoulTurnover                     = 37
	FiveSecondInboundTurnover                 = 38
	StepOutOfBoundsTurnover                   = 39
	OutOfBoundsLostBallTurnover               = 40
	StealTurnover                             = 41
	PlayerOutOfBoundsTurnover                 = 43
	OutOfBoundsBadPassTurnover                = 45
)
