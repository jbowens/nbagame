package data

import (
	"database/sql/driver"
	"strings"
)

// TimeoutType describes the type of ejection.
type EjectionType int

const (
	EjectionTypeNone              EjectionType = 0
	EjectionTypeSecondTechnical                = 1
	EjectionTypeSecondFlagrantOne              = 2
	EjectionTypeFlagrantTwo                    = 3
	EjectionTypeOther                          = 4
)

func (et EjectionType) String() string {
	if s, ok := ejectionToString[et]; ok {
		return s
	}
	return "unknown"
}

func (et EjectionType) MarshalText() ([]byte, error) {
	return []byte(strings.Replace(et.String(), " ", "_", -1)), nil
}

func (et EjectionType) Value() (driver.Value, error) {
	b, err := et.MarshalText()
	if err != nil {
		return nil, err
	}
	return string(b), nil
}

var (
	ejectionToString = map[EjectionType]string{
		EjectionTypeNone:              "no",
		EjectionTypeSecondTechnical:   "second technical",
		EjectionTypeSecondFlagrantOne: "second flagrant one",
		EjectionTypeFlagrantTwo:       "flagrant two",
		EjectionTypeOther:             "other",
	}
)
