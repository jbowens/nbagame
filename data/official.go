package data

// Official represents an NBA official.
type Official struct {
	ID           int    `json:"id" db:"id"`
	FirstName    string `json:"first_name" db:"first_name"`
	LastName     string `json:"last_name" db:"last_name"`
	JerseyNumber string `json:"jersey_number" db:"jersey_number"`
}
