package endpoints

import (
	"fmt"
	"testing"
)

func TestShotChartDetail(t *testing.T) {
	params := ShotChartDetailParams{
		ContextMeasure: "FGA",
		EndPeriod:      10,
		EndRange:       28800,
		GameID:         "0021401203",
		LeagueID:       "00",
		PlayerID:       2406,
		Season:         "2014-15",
		SeasonType:     "Regular Season",
		StartPeriod:    1,
		TeamID:         1610612765,
	}

	var resp ShotChartDetailResponse
	if err := DefaultRequester.Request("shotchartdetail", params, &resp); err != nil {
		t.Fatal(err)
	}
	if len(resp.ShotDetails) == 0 {
		t.Error("Empty response for boxscoresummary request.")
	}

	for _, shot := range resp.ShotDetails {
		fmt.Printf("%v\t%v\t%v\n", shot.Period, shot.MinutesRemaining*60+shot.SecondsRemaining, shot.ShotMade)
	}
}
