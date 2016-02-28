package data

import (
	"database/sql/driver"
	"strings"
)

// ViolationType describes the type of a violation.
type ViolationType int

const (
	ViolationTypeNone                 ViolationType = 0
	ViolationTypeDelayOfGame                        = 1
	ViolationTypeDefensiveGoaltending               = 2
	ViolationTypeLane                               = 3
	ViolationTypeJumpBall                           = 4
	ViolationTypeKickedBall                         = 5
	ViolationTypeDoubleLane                         = 6
)

func (vt ViolationType) String() string {
	if s, ok := violationToString[vt]; ok {
		return s
	}
	return "unknown"
}

func (vt ViolationType) MarshalText() ([]byte, error) {
	return []byte(strings.Replace(vt.String(), " ", "_", -1)), nil
}

func (vt ViolationType) Value() (driver.Value, error) {
	b, err := vt.MarshalText()
	if err != nil {
		return nil, err
	}
	return string(b), nil
}

var (
	violationToString = map[ViolationType]string{
		ViolationTypeNone:                 "none",
		ViolationTypeDelayOfGame:          "delay of game",
		ViolationTypeDefensiveGoaltending: "defensive goaltending",
		ViolationTypeLane:                 "lane",
		ViolationTypeJumpBall:             "jump ball",
		ViolationTypeKickedBall:           "kicked ball",
		ViolationTypeDoubleLane:           "double lane",
	}
)
