package data

import (
	"database/sql/driver"
	"strings"
)

// FoulType describes the type of foul.
// See http://www.nba.com/analysis/rules_12.html
type FoulType int

const (
	FoulTypePersonal                FoulType = 1
	FoulTypeShooting                         = 2
	FoulTypeLooseBall                        = 3
	FoulTypeOffensive                        = 4
	FoulTypeInbound                          = 5
	FoulTypeAwayFromPlay                     = 6
	FoulTypePunching                         = 8
	FoulTypeClearPath                        = 9
	FoulTypeTechnical                        = 11
	FoulTypeUnsportsmanlike                  = 12
	FoulTypeHangingTechnical                 = 13
	FoulTypeFlagrantOne                      = 14
	FoulTypeFlagrantTwo                      = 15
	FoulTypeTeamTechnical                    = 17
	FoulTypeDelay                            = 18
	FoulTypeTaunting                         = 19
	FoulTypeExcessTimeout                    = 25
	FoulTypeOffensiveCharge                  = 26
	FoulTypePersonalBlock                    = 27
	FoulTypePersonalTake                     = 28
	FoulTypeShootingBlock                    = 29
	FoulTypeTooManyPlayersTechnical          = 30
)

func (ft FoulType) String() string {
	if s, ok := foulToString[ft]; ok {
		return s
	}
	return "unknown"
}

func (ft FoulType) MarshalText() ([]byte, error) {
	return []byte(strings.Replace(ft.String(), " ", "_", -1)), nil
}

func (ft FoulType) Value() (driver.Value, error) {
	b, err := ft.MarshalText()
	if err != nil {
		return nil, err
	}
	return string(b), nil
}

var (
	foulToString = map[FoulType]string{
		FoulTypePersonal:                "personal",
		FoulTypeShooting:                "shooting",
		FoulTypeLooseBall:               "loose ball",
		FoulTypeOffensive:               "offensive",
		FoulTypeInbound:                 "inbound",
		FoulTypeAwayFromPlay:            "away from play",
		FoulTypePunching:                "punching",
		FoulTypeClearPath:               "clear path",
		FoulTypeTechnical:               "technical",
		FoulTypeUnsportsmanlike:         "unsportsmanlike",
		FoulTypeHangingTechnical:        "hanging technical",
		FoulTypeFlagrantOne:             "flagrant one",
		FoulTypeFlagrantTwo:             "flagrant two",
		FoulTypeTeamTechnical:           "team technical",
		FoulTypeDelay:                   "delay",
		FoulTypeTaunting:                "taunting",
		FoulTypeExcessTimeout:           "excess timeout",
		FoulTypeOffensiveCharge:         "offensive charge",
		FoulTypePersonalBlock:           "personal block",
		FoulTypePersonalTake:            "personal take",
		FoulTypeShootingBlock:           "shooting block",
		FoulTypeTooManyPlayersTechnical: "too many players technical",
	}
)

type Foul struct {
	Type     FoulType
	Offender *PlayerDescription
}
