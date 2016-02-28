package data

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
