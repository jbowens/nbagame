package results

import (
	"reflect"
	"testing"
)

type testPlayer struct {
	PlayerID int    `nbagame:"PLAYER_ID"`
	Name     string `nbagame:"NAME"`
}

func TestSimpleDecoderExample(t *testing.T) {
	rs := &ResultSet{
		Name:    "common_all_players",
		Headers: []string{"PLAYER_ID", "NAME"},
		RowSet: [][]interface{}{
			[]interface{}{20, "Delonte West"},
			[]interface{}{56, "Chris Bosh"},
			[]interface{}{295, "Chef Curry"},
		},
	}

	var players []*testPlayer
	if err := rs.Decode(&players); err != nil {
		t.Fatal(err)
	}
	if len(players) != 3 {
		t.Errorf("Expected 3 players, but got: %+v", players)
	}

	if !reflect.DeepEqual(players, []*testPlayer{
		{20, "Delonte West"},
		{56, "Chris Bosh"},
		{295, "Chef Curry"},
	}) {
		t.Errorf("Populated players didn't match expected: %+v", players)
	}
}
