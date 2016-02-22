package data

// TurnoverType defines the different types of turnovers.
type TurnoverType int

// TODO(jackson): Fill in the gaps in the turnover table.
const (
	TurnoverUnknown              TurnoverType = 0
	TurnoverBadPass                           = 1
	TurnoverLostBall                          = 2
	TurnoverOutOfBounds                       = 3
	TurnoverTraveling                         = 4
	TurnoverFoul                              = 5
	TurnoverDoubleDribble                     = 6
	TurnoverDiscontinueDribble                = 7
	TurnoverThreeSecondViolation              = 8
	TurnoverFiveSecondViolation               = 9
	TurnoverEightSecondViolation              = 10
	TurnoverShotClock                         = 11
	TurnoverInbound                           = 12
	TurnoverBackcourt                         = 13
	TurnoverOffensiveGoaltending              = 15
	TurnoverLaneViolation                     = 17
	TurnoverJumpBallViolation                 = 18
	TurnoverKickedBall                        = 19
	TurnoverIllegalAssist                     = 20
	TurnoverPalming                           = 21
	TurnoverPunchedBall                       = 33
	TurnoverSwingingElbows                    = 34
	TurnoverBasketFromBelow                   = 35
	TurnoverIllegalScreen                     = 36
	TurnoverOffensiveFoul                     = 37
	TurnoverFiveSecondInbound                 = 38
	TurnoverStepOutOfBounds                   = 39
	TurnoverOutOfBoundsLostBall               = 40
	TurnoverSteal                             = 41
	TurnoverPlayerOutOfBounds                 = 43
	TurnoverOutOfBoundsBadPass                = 45
)
