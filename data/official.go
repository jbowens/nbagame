package data

// Official represents an NBA official.
type Official struct {
	ID           int    `json:"id" db:"id"`
	FirstName    string `json:"first_name" db:"first_name"`
	LastName     string `json:"last_name" db:"last_name"`
	JerseyNumber string `json:"jersey_number" db:"jersey_number"`
}

// Officiated represents the fact that the given NBA official officiated
// the given NBA game.
type Officiated struct {
	GameID     GameID `json:"game_id" db:"game_id"`
	OfficialID int    `json:"official_id" db:"official_id"`
}
