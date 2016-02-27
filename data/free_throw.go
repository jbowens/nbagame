package data

import "strings"

// FreeThrowType defines the type of free throw.
type FreeThrowType int

const (
	FreeThrowOneOfOne          = 10
	FreeThrowOneOfTwo          = 11
	FreeThrowTwoOfTwo          = 12
	FreeThrowOneOfThree        = 13
	FreeThrowTwoOfThree        = 14
	FreeThrowThreeOfThree      = 15
	FreeThrowTechnical         = 16
	FreeThrowFlagrantOneOfTwo  = 18
	FreeThrowFlagrantTwoOfTwo  = 19
	FreeThrowFlagrantOneOfOne  = 20
	FreeThrowClearPathOneOfTwo = 25
	FreeThrowClearPathTwoOfTwo = 26
)

func (ft FreeThrowType) String() string {
	return freeThrowToString[ft]
}

func (ft FreeThrowType) MarshalText() ([]byte, error) {
	return []byte(strings.Replace(ft.String(), " ", "_", -1)), nil
}

var (
	freeThrowToString = map[FreeThrowType]string{
		FreeThrowOneOfOne:          "one of one",
		FreeThrowOneOfTwo:          "one of two",
		FreeThrowTwoOfTwo:          "two of two",
		FreeThrowOneOfThree:        "one of three",
		FreeThrowTwoOfThree:        "two of three",
		FreeThrowThreeOfThree:      "three of three",
		FreeThrowTechnical:         "technical",
		FreeThrowFlagrantOneOfTwo:  "flagrant one of two",
		FreeThrowFlagrantTwoOfTwo:  "flagrant two of two",
		FreeThrowFlagrantOneOfOne:  "flagrant one of one",
		FreeThrowClearPathOneOfTwo: "clear path one of two",
		FreeThrowClearPathTwoOfTwo: "clear path two of two",
	}
)
