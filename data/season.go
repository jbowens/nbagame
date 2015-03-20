package data

import (
	"fmt"
	"strconv"
	"strings"
	"time"
)

// Season represents a season identifier. These may be returned by some API endpoints
// and used as parameters for others.
type Season string

var (
	// CurrentSeason holds the season identifier for the current season. The season
	// switches starting on July 1st.
	CurrentSeason Season
)

func init() {
	now := time.Now()
	year := now.Year()

	var seasonStr string
	if now.Month() >= time.July {
		seasonStr = fmt.Sprintf("%d-%s", year, strconv.Itoa(year + 1)[2:])
	} else {
		seasonStr = fmt.Sprintf("%d-%s", year-1, strconv.Itoa(year)[2:])
	}
	CurrentSeason = Season(seasonStr)
}

// AddYears returns the season identifier the given number of years away. Years may
// be negative to go backwards.
func (s Season) AddYears(years int) Season {
	pieces := strings.SplitN(string(s), "-", 2)
	first, _ := strconv.Atoi(pieces[0])
	second, _ := strconv.Atoi(pieces[1])

	secondStr := strconv.Itoa(second + years)
	return Season(fmt.Sprintf("%d-%s", first+years, secondStr[len(secondStr)-2:]))
}

// Previous returns the season before the given season.
func (s Season) Previous() Season {
	return s.AddYears(-1)
}

// Next returns the season after the given season.
func (s Season) Next() Season {
	return s.AddYears(1)
}
