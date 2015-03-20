package endpoints

import (
	"testing"
)

type testRequestParams struct {
	Season      string `json:"Season"`
	PlayerCount int    `json:"PlayerCount"`
}

func TestMakeParams(t *testing.T) {
	testParams := testRequestParams{
		Season:      "2013-14",
		PlayerCount: 20,
	}

	params, err := DefaultRequester.makeParams(testParams)
	if err != nil {
		t.Fatal(err)
	}

	if params.Get("Season") != "2013-14" {
		t.Errorf("Season doesn't match: %v", params)
	}
	if params.Get("PlayerCount") != "20" {
		t.Errorf("PlayerCount doesn't match: %v", params)
	}
}

func TestRequest(t *testing.T) {
	params := CommonAllPlayersParams{
		Season:   "2014-15",
		LeagueID: "00",
	}

	var resp CommonAllPlayersResponse
	if err := DefaultRequester.Request("commonallplayers", params, &resp); err != nil {
		t.Fatal(err)
	}
	if len(resp.CommonAllPlayers) == 0 {
		t.Error("Empty response for commonallplayers request.")
	}
}
