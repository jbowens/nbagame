package nbagame

import (
	"testing"

	"github.com/jbowens/nbagame/data"
)

const (
	atlantaHawksTeamID = 1610612737
	twentyFourteen     = data.Season("2014-15")
)

func TestGetHistoricalPlayers(t *testing.T) {
	allPlayers, err := DefaultClient.HistoricalPlayers()
	if err != nil {
		t.Fatal(err)
	}

	if len(allPlayers) < 100 {
		t.Errorf("Did not get all players, got %v players", len(allPlayers))
	}
}

func TestGetGamesPlayedByTeam(t *testing.T) {
	gameIDs, err := DefaultClient.GamesPlayedBy(twentyFourteen, atlantaHawksTeamID)
	if err != nil {
		t.Fatal(err)
	}
	for _, id := range gameIDs {
		t.Log(id)
	}
	if len(gameIDs) < 50 {
		t.Errorf("Expected 50 or more games, but got %v", len(gameIDs))
	}
}
