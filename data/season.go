package data

import (
	"database/sql/driver"
	"errors"
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

func (s Season) Value() (driver.Value, error) {
	return string(s), nil
}

func (s *Season) UnmarshalText(text []byte) error {
	var startYear, endYear int
	_, err := fmt.Sscanf(string(text), "%4d-%2d", &startYear, &endYear)
	if err != nil {
		return err
	}
	if endYear != ((startYear + 1) % 100) {
		return errors.New("invalid season, start and end must be consecutive")
	}
	*s = Season(fmt.Sprintf("%d-%d", startYear, endYear))
	return nil
}

// FallYear returns the beginning, fall year of this season. For ex,
// for "2014-15" it will return 2014.
func (s Season) FallYear() int {
	pieces := strings.SplitN(string(s), "-", 2)
	first, _ := strconv.Atoi(pieces[0])
	return first
}

// FallYear returns the end, spring year of this season. For ex,
// for "2014-15" it will return 2015.
func (s Season) SpringYear() int {
	pieces := strings.SplitN(string(s), "-", 2)
	first, _ := strconv.Atoi(pieces[0])
	return first + 1
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

// String returns a string representation of the season.
func (s Season) String() string {
	return string(s)
}
