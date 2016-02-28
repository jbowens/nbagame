package data

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

var (
	timeoutToString = map[TimeoutType]string{
		TimeoutTypeNone:     "no",
		TimeoutTypeRegular:  "regular",
		TimeoutTypeShort:    "short",
		TimeoutTypeOfficial: "official",
	}
)
