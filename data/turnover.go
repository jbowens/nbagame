package data

import (
	"database/sql/driver"
	"strings"
)

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
	TurnoverDoublePersonal                    = 23
	TurnoverOppositeBasket                    = 32
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
	TurnoverTooManyPlayers                    = 44
	TurnoverOutOfBoundsBadPass                = 45
)

func (t TurnoverType) String() string {
	if s, ok := turnoverToString[t]; ok {
		return s
	}
	return "unknown turnover"
}

func (t TurnoverType) MarshalText() ([]byte, error) {
	return []byte(strings.Replace(t.String(), " ", "_", -1)), nil
}

func (t TurnoverType) Value() (driver.Value, error) {
	b, err := t.MarshalText()
	if err != nil {
		return nil, err
	}
	return string(b), nil
}

var (
	turnoverToString = map[TurnoverType]string{
		TurnoverUnknown:              "unknown turnover",
		TurnoverBadPass:              "bad pass",
		TurnoverLostBall:             "lost ball",
		TurnoverOutOfBounds:          "out of bounds",
		TurnoverTraveling:            "traveling",
		TurnoverFoul:                 "foul",
		TurnoverDoubleDribble:        "double dribble",
		TurnoverDiscontinueDribble:   "discontinue dribble",
		TurnoverThreeSecondViolation: "three second violation",
		TurnoverFiveSecondViolation:  "five second violation",
		TurnoverEightSecondViolation: "eight second violation",
		TurnoverShotClock:            "shot clock violation",
		TurnoverInbound:              "inbound",
		TurnoverBackcourt:            "backcourt violation",
		TurnoverOffensiveGoaltending: "offensive goaltending",
		TurnoverLaneViolation:        "lane violation",
		TurnoverJumpBallViolation:    "jump ball violation",
		TurnoverKickedBall:           "kicked ball",
		TurnoverIllegalAssist:        "illegal assist",
		TurnoverPalming:              "palming",
		TurnoverDoublePersonal:       "double personal",
		TurnoverOppositeBasket:       "opposite basket",
		TurnoverPunchedBall:          "punched ball",
		TurnoverSwingingElbows:       "swinging elbows",
		TurnoverBasketFromBelow:      "basket from below",
		TurnoverIllegalScreen:        "illegal screen",
		TurnoverOffensiveFoul:        "offensive foul",
		TurnoverFiveSecondInbound:    "five second inbound",
		TurnoverStepOutOfBounds:      "step out of bounds",
		TurnoverOutOfBoundsLostBall:  "out of bounds lost ball",
		TurnoverSteal:                "steal",
		TurnoverPlayerOutOfBounds:    "player out of bounds",
		TurnoverTooManyPlayers:       "too many players",
		TurnoverOutOfBoundsBadPass:   "out of bounds bad pass",
	}
)
