package endpoints

import "testing"

func TestCommonPlayerInfo(t *testing.T) {
	params := CommonPlayerInfoParams{
		LeagueID: "00",
		PlayerID: 201566,
	}

	var resp CommonPlayerInfoResponse
	if err := DefaultRequester.Request("commonplayerinfo", params, &resp); err != nil {
		t.Fatal(err)
	}
	if len(resp.CommonPlayerInfo) == 0 {
		t.Error("Empty response for commonplayerinfo request.")
	}

	row := resp.CommonPlayerInfo[0]
	if row.Birthdate == "" {
		t.Errorf("expected Russell Westbrook to have a birthday, but got: %+v", row)
	}

	details, err := row.ToPlayerDetails()
	if err != nil {
		t.Fatal(err)
	}
	if details.Birthdate == nil {
		t.Errorf("expected player details to have birthdate %s but got %+v",
			row.Birthdate, details)
	}
}
