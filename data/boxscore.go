package data

// BoxScore holds an individual game's box score.
type BoxScore struct {
	TeamStats   []*TeamStats
	PlayerStats []*PlayerStats
}

// Team returns the stats for the team with the given ID. If no team
// with the given ID played in the game, then nil is returned.
func (box *BoxScore) Team(teamID int) *TeamStats {
	for _, teamStats := range box.TeamStats {
		if teamStats.TeamID == teamID {
			return teamStats
		}
	}
	return nil
}

// Player returns the stats for the player with the given ID. If the
// player with the given ID did not play in the game, then nil is returned.
func (box *BoxScore) Player(playerID int) *PlayerStats {
	for _, playerStats := range box.PlayerStats {
		if playerStats.PlayerID == playerID {
			return playerStats
		}
	}
	return nil
}
