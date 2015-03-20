package endpoints

// CommonAllPlayersParams defines all the parameters for a CommonAllPlayers request.
type CommonAllPlayersParams struct {
	LeagueID            string `json:"LeagueID"`
	Season              string `json:"Season"`
	IsOnlyCurrentSeason int    `json:"IsOnlyCurrentSeason"`
}
