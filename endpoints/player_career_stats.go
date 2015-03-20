package endpoints

// PlayerCareerStats defines all the parameters for a PlayerCareerStats request.
type PlayerCareerStats struct {
	PerMode  string `json:"PerMode"`
	PlayerID int    `json:"PlayerID"`
	LeagueID string `json:"LeagueID"`
}
