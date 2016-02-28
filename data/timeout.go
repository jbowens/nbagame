package data

import (
	"database/sql/driver"
	"strings"
)

// TimeoutType describes the type of a timeout.
type TimeoutType int

const (
	TimeoutTypeNone     TimeoutType = 0
	TimeoutTypeRegular              = 1
	TimeoutTypeShort                = 2
	TimeoutTypeOfficial             = 4
)

func (tt TimeoutType) String() string {
	if s, ok := timeoutToString[tt]; ok {
		return s
	}
	return "unknown"
}

func (tt TimeoutType) MarshalText() ([]byte, error) {
	return []byte(strings.Replace(tt.String(), " ", "_", -1)), nil
}

func (tt TimeoutType) Value() (driver.Value, error) {
	b, err := tt.MarshalText()
	if err != nil {
		return nil, err
	}
	return string(b), nil
}

var (
	timeoutToString = map[TimeoutType]string{
		TimeoutTypeNone:     "no",
		TimeoutTypeRegular:  "regular",
		TimeoutTypeShort:    "short",
		TimeoutTypeOfficial: "official",
	}
)
