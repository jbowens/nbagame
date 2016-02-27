package data

import "strings"

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
	}
)

// ShotType is an enum of attributes of a shot attempt. Ex: was it a layup?
// a dunk? bank shot? put back? etc.
type ShotType int

func (st ShotType) String() string {
	return shotTypeToString[st]
}

func (st ShotType) MarshalText() ([]byte, error) {
	s := st.String()
	return []byte(strings.Replace(s, " ", "_", -1)), nil
}

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
)

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
