package data

import (
	"database/sql/driver"
	"strings"
)

const (
	ShotTypeNone ShotType = iota
	ShotTypeJump
	ShotTypeHook
	ShotTypeLayup
	ShotTypeDunk
	ShotTypeRunning
	ShotTypeDriving
	ShotTypeAlleyOop
	ShotTypeReverse
	ShotTypeTurnaround
	ShotTypeFadeaway
	ShotTypeBank
	ShotTypeFingerRoll
	ShotTypePutBack
	ShotTypeFloating
	ShotTypePullUp
	ShotTypeStepBack
	ShotTypeTipIn
	ShotTypeCutting
	ShotTypeFollowUp
)

// ShotType is an enum of attributes of a shot attempt. Ex: was it a layup?
// a dunk? bank shot? put back? etc.
type ShotType int

var (
	shotTypeToString = map[ShotType]string{
		ShotTypeNone:       "no",
		ShotTypeJump:       "jump",
		ShotTypeHook:       "hook",
		ShotTypeLayup:      "layup",
		ShotTypeDunk:       "dunk",
		ShotTypeRunning:    "running",
		ShotTypeDriving:    "driving",
		ShotTypeAlleyOop:   "alleyoop",
		ShotTypeReverse:    "reverse",
		ShotTypeTurnaround: "turnaround",
		ShotTypeFadeaway:   "fadeaway",
		ShotTypeBank:       "bank",
		ShotTypeFingerRoll: "finger roll",
		ShotTypePutBack:    "put back",
		ShotTypeFloating:   "floating",
		ShotTypePullUp:     "pull up",
		ShotTypeStepBack:   "step back",
		ShotTypeTipIn:      "tip in",
		ShotTypeCutting:    "cutting",
		ShotTypeFollowUp:   "follow up",
	}
)

func (st ShotType) String() string {
	if s, ok := shotTypeToString[st]; ok {
		return s
	}
	return "unknown"
}

func (st ShotType) MarshalText() ([]byte, error) {
	s := st.String()
	return []byte(strings.Replace(s, " ", "_", -1)), nil
}

func (st ShotType) Value() (driver.Value, error) {
	b, err := st.MarshalText()
	if err != nil {
		return nil, err
	}
	return string(b), nil
}

// Shot describes a shot attempt.
type Shot struct {
	Made            bool            `json:"made"`
	PointsScored    int             `json:"points_scored"`
	PointsAttempted int             `json:"points_attempted"`
	Description     ShotDescription `json:"description"`
}

// ShotDescription describes a shot as a slice of ShotTypes.
type ShotDescription []ShotType

// Is checks if the shot description contains the provided types.
func (sd ShotDescription) Is(typs ...ShotType) bool {
	set := map[ShotType]bool{}
	for _, typ := range sd {
		set[typ] = true
	}

	var ok bool = true
	for _, t := range typs {
		ok = ok && set[t]
	}
	return ok
}
