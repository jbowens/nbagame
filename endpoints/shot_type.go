package endpoints

import "github.com/jbowens/nbagame/data"

var (
	// shotActionTypeToShotTypes maps the event message action type
	// on a play event to the shot types describing the shot. This is a
	// monster of a mapping, but it's what we need to do.
	shotActionTypeToShotTypes = map[int][]data.ShotType{
		0:   {data.ShotTypeNone},
		1:   {data.ShotTypeJump},
		2:   {data.ShotTypeRunning, data.ShotTypeJump},
		3:   {data.ShotTypeHook},
		4:   {data.ShotTypeTipIn},
		5:   {data.ShotTypeLayup},
		7:   {data.ShotTypeDunk},
		8:   {data.ShotTypeDunk},
		40:  {data.ShotTypeLayup},
		41:  {data.ShotTypeRunning, data.ShotTypeLayup},
		42:  {data.ShotTypeDriving, data.ShotTypeLayup},
		43:  {data.ShotTypeAlleyOop, data.ShotTypeLayup},
		44:  {data.ShotTypeReverse, data.ShotTypeLayup},
		45:  {data.ShotTypeJump},
		46:  {data.ShotTypeRunning, data.ShotTypeJump},
		47:  {data.ShotTypeTurnaround, data.ShotTypeJump},
		48:  {data.ShotTypeDunk},
		49:  {data.ShotTypeDriving, data.ShotTypeDunk},
		50:  {data.ShotTypeRunning, data.ShotTypeDunk},
		51:  {data.ShotTypeReverse, data.ShotTypeDunk},
		52:  {data.ShotTypeAlleyOop, data.ShotTypeDunk},
		53:  {data.ShotTypeTipIn},
		54:  {data.ShotTypeRunning, data.ShotTypeTipIn},
		55:  {data.ShotTypeHook},
		56:  {data.ShotTypeRunning, data.ShotTypeHook},
		57:  {data.ShotTypeDriving, data.ShotTypeHook},
		58:  {data.ShotTypeTurnaround, data.ShotTypeHook},
		63:  {data.ShotTypeFadeaway, data.ShotTypeJump},
		65:  {data.ShotTypeJump, data.ShotTypeHook},
		66:  {data.ShotTypeJump, data.ShotTypeBank},
		67:  {data.ShotTypeHook, data.ShotTypeBank},
		71:  {data.ShotTypeFingerRoll, data.ShotTypeLayup},
		72:  {data.ShotTypePutBack, data.ShotTypeLayup},
		73:  {data.ShotTypeDriving, data.ShotTypeReverse, data.ShotTypeLayup},
		74:  {data.ShotTypeRunning, data.ShotTypeReverse, data.ShotTypeLayup},
		75:  {data.ShotTypeDriving, data.ShotTypeFingerRoll, data.ShotTypeLayup},
		76:  {data.ShotTypeRunning, data.ShotTypeFingerRoll, data.ShotTypeLayup},
		77:  {data.ShotTypeDriving, data.ShotTypeJump},
		78:  {data.ShotTypeFloating, data.ShotTypeJump},
		79:  {data.ShotTypePullUp, data.ShotTypeJump},
		80:  {data.ShotTypeStepBack, data.ShotTypeJump},
		81:  {data.ShotTypePullUp, data.ShotTypeBank},
		82:  {data.ShotTypeDriving, data.ShotTypeBank},
		83:  {data.ShotTypeFadeaway, data.ShotTypeBank},
		84:  {data.ShotTypeRunning, data.ShotTypeBank},
		85:  {data.ShotTypeTurnaround, data.ShotTypeBank},
		86:  {data.ShotTypeTurnaround, data.ShotTypeFadeaway},
		87:  {data.ShotTypePutBack, data.ShotTypeDunk},
		88:  {data.ShotTypeDriving, data.ShotTypeDunk},
		89:  {data.ShotTypeReverse, data.ShotTypeDunk},
		90:  {data.ShotTypeRunning, data.ShotTypeDunk},
		91:  {data.ShotTypePutBack, data.ShotTypeReverse, data.ShotTypeDunk},
		92:  {data.ShotTypePutBack, data.ShotTypeDunk},
		93:  {data.ShotTypeDriving, data.ShotTypeBank, data.ShotTypeHook},
		94:  {data.ShotTypeJump, data.ShotTypeBank, data.ShotTypeHook},
		95:  {data.ShotTypeRunning, data.ShotTypeBank, data.ShotTypeHook},
		96:  {data.ShotTypeTurnaround, data.ShotTypeBank, data.ShotTypeHook},
		97:  {data.ShotTypeTipIn, data.ShotTypeLayup},
		98:  {data.ShotTypeCutting, data.ShotTypeLayup},
		99:  {data.ShotTypeCutting, data.ShotTypeFingerRoll, data.ShotTypeLayup},
		100: {data.ShotTypeRunning, data.ShotTypeAlleyOop, data.ShotTypeLayup},
		101: {data.ShotTypeDriving, data.ShotTypeFloating, data.ShotTypeJump},
		102: {data.ShotTypeDriving, data.ShotTypeFloating, data.ShotTypeBank, data.ShotTypeJump},
		103: {data.ShotTypeRunning, data.ShotTypePullUp, data.ShotTypeJump},
		104: {data.ShotTypeStepBack, data.ShotTypeBank, data.ShotTypeJump},
		105: {data.ShotTypeTurnaround, data.ShotTypeFadeaway, data.ShotTypeBank, data.ShotTypeJump},
		106: {data.ShotTypeRunning, data.ShotTypeAlleyOop, data.ShotTypeDunk},
		107: {data.ShotTypeTipIn, data.ShotTypeDunk},
		108: {data.ShotTypeCutting, data.ShotTypeDunk},
		109: {data.ShotTypeDriving, data.ShotTypeReverse, data.ShotTypeDunk},
		110: {data.ShotTypeRunning, data.ShotTypeReverse, data.ShotTypeDunk},
	}
)
