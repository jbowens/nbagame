package data

import "time"

// GameID holds a unique identifier for an NBA game. The identifier is unique
// across all seasons and teams.
type GameID string

func (id GameID) String() string {
	return string(id)
}

// GameStatus indicates the status of a game.
type GameStatus int

// TODO: Identify and populate the rest of the GameStatus values.
const (
	// Unknown is used for unrecognized game status IDs.
	Unknown GameStatus = 0
	// Live indicates that a game is in progress.
	Live GameStatus = 2
	// Final indicates a Game's score is Final and the game has finished.
	Final GameStatus = 3
)

// Game holds basic information about a NBA game.
type Game struct {
	ID                GameID
	HomeTeamID        int
	VisitorTeamID     int
	Season            Season
	Status            GameStatus
	LastMeetingGameID GameID
}

// GameDetails provides detailed information and summary of an NBA game.
type GameDetails struct {
	Game
	Date          time.Time
	LengthMinutes int
	Attendance    int
	Officials     []*Official
	HomePoints    *PointSummary
	VisitorPoints *PointSummary
	LeadChanges   int
	TimesTied     int
}

// PointSummary provides aggregate team point statistics.
type PointSummary struct {
	InPaint       int
	SecondChance  int
	FromBench     int
	FirstQuarter  int
	SecondQuarter int
	ThirdQuarter  int
	FourthQuarter int
	Overtime      []int
	Total         int
}
