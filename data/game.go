package data

// GameID holds a unique identifier for an NBA game. The identifier is unique
// across all seasons and teams.
type GameID string

// GameStatus indicates the status of a game.
type GameStatus int

const (
	// Final indicates a Game's score is Final and the game has finished.
	Final GameStatus = 3
	// TODO: Identify and populate the rest of the GameStatus fields.
)

// Games hold information about a NBA game.
type Game struct {
	ID                GameID
	HomeTeamID        int
	VisitorTeamID     int
	Season            Season
	Status            GameStatus
	LastMeetingGameID int
}
