package endpoints

import "testing"

func TestBoxScoreSummary(t *testing.T) {
	params := BoxScoreSummaryParams{
		GameID: "0021401185",
	}

	var resp BoxScoreSummaryResponse
	if err := DefaultRequester.Request("boxscoresummaryv2", params, &resp); err != nil {
		t.Fatal(err)
	}
	if len(resp.GameSummary) == 0 {
		t.Error("Empty response for boxscoresummary request.")
	}
}
