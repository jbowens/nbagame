package data

// Official represents an NBA official.
type Official struct {
	ID           int    `json:"id"`
	FirstName    string `json:"first_name"`
	LastName     string `json:"last_name"`
	JerseyNumber string `json:"jersey_number"`
}
