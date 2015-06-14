package endpoints

import (
	"strconv"
	"strings"

	"github.com/jbowens/nbagame/data"
)

// ConvertGameStatus takes an integer status ID returned by an endpoint and converts
// it to a GameStatus.
func ConvertGameStatus(statusID int) data.GameStatus {
	for _, status := range knownGameStatus {
		if int(status) == statusID {
			return status
		}
	}
	return data.Unknown
}

// HourMinuteStringToMinutes converts a string containing hours and minutes, ie "2:06"
// to the number of minutes.
func HourMinuteStringToMinutes(hourString string) int {
	pieces := strings.Split(hourString, ":")
	if len(pieces) < 2 {
		return 0
	}

	hours, err := strconv.Atoi(pieces[0])
	if err != nil {
		return 0
	}

	minutes, err := strconv.Atoi(pieces[1])
	if err != nil {
		return 0
	}

	return 60*hours + minutes
}

func nonZero(nums ...int) []int {
	return nums
}
