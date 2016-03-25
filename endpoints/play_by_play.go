package endpoints

import (
	"strconv"
	"strings"

	"github.com/jbowens/nbagame/data"
)

type PersonType int

const (
	HomePlayer    PersonType = 4
	VisitorPlayer PersonType = 5
)

// PlayByPlayParams defines parameters for a PlayByPlay request.
// http://stats.nba.com/stats/playbyplayv2?EndPeriod=10&EndRange=55800&GameID=0021401229&RangeType=2&Season=2014-15&SeasonType=Regular+Season&StartPeriod=1&StartRange=0
type PlayByPlayParams struct {
	Season      string `json:"Season"`
	SeasonType  string `json:"SeasonType"`
	GameID      string `json:"GameID"`
	StartPeriod int    `json:"StartPeriod"`
	EndPeriod   int    `json:"EndPeriod"`
	RangeType   int    `json:"RangeType"`
	StartRange  int    `json:"StartRange"`
	EndRange    int    `json:"EndRange"`
}

// PlayByPlayResponse is the type for all result sets returned by the
// 'playbyplay' resource.
type PlayByPlayResponse struct {
	PlayByPlay []*PlayByPlayRow `nbagame:"PlayByPlay"`
}

func (resp *PlayByPlayResponse) ToData() []*data.Event {
	var events []*data.Event
	for _, row := range resp.PlayByPlay {
		events = append(events, row.ToData())
	}
	return events
}

// PlayByPlayRow represents the schema returned for 'PlayByPlay' result
// sets, returned from the 'playbyplay' resource.
type PlayByPlayRow struct {
	GameID                 string  `nbagame:"GAME_ID"`
	EventNumber            int     `nbagame:"EVENTNUM"`
	EventMessageType       int     `nbagame:"EVENTMSGTYPE"`
	EventMessageActionType int     `nbagame:"EVENTMSGACTIONTYPE"`
	Period                 int     `nbagame:"PERIOD"`
	WallClockTimeString    string  `nbagame:"WCTIMESTRING"`
	PeriodClockTimeString  string  `nbagame:"PCTIMESTRING"`
	HomeDescription        *string `nbagame:"HOMEDESCRIPTION"`
	NeutralDescription     *string `nbagame:"NEUTRALDESCRIPTION"`
	VisitorDescription     *string `nbagame:"VISITORDESCRIPTION"`
	ScoreString            *string `nbagame:"SCORE"`       // ex. "94 - 97"
	ScoreMargin            *string `nbagame:"SCOREMARGIN"` // ex. 5, -3, or "TIE" o_O
	// First person involved in play
	Person1Type             int    `nbagame:"PERSON1TYPE"`
	Player1ID               int    `nbagame:"PLAYER1_ID"`
	Player1Name             string `nbagame:"PLAYER1_NAME"`
	Player1TeamID           int    `nbagame:"PLAYER1_TEAM_ID"`
	Player1TeamCity         string `nbagame:"PLAYER1_TEAM_CITY"`
	Player1TeamNickname     string `nbagame:"PLAYER1_TEAM_NICKNAME"`
	Player1TeamAbbreviation string `nbagame:"PLAYER1_TEAM_ABBREVIATION"`
	// Second person involved in play
	Person2Type             int    `nbagame:"PERSON2TYPE"`
	Player2ID               int    `nbagame:"PLAYER2_ID"`
	Player2Name             string `nbagame:"PLAYER2_NAME"`
	Player2TeamID           int    `nbagame:"PLAYER2_TEAM_ID"`
	Player2TeamCity         string `nbagame:"PLAYER2_TEAM_CITY"`
	Player2TeamNickname     string `nbagame:"PLAYER2_TEAM_NICKNAME"`
	Player2TeamAbbreviation string `nbagame:"PLAYER2_TEAM_ABBREVIATION"`
	// Third person involved in play
	Person3Type             int    `nbagame:"PERSON3TYPE"`
	Player3ID               int    `nbagame:"PLAYER3_ID"`
	Player3Name             string `nbagame:"PLAYER3_NAME"`
	Player3TeamID           int    `nbagame:"PLAYER3_TEAM_ID"`
	Player3TeamCity         string `nbagame:"PLAYER3_TEAM_CITY"`
	Player3TeamNickname     string `nbagame:"PLAYER3_TEAM_NICKNAME"`
	Player3TeamAbbreviation string `nbagame:"PLAYER3_TEAM_ABBREVIATION"`
}

func (r *PlayByPlayRow) ToData() *data.Event {
	var player1, player2, player3 *data.PlayerDescription
	if r.Player1ID != 0 && r.Player1Name != "" && r.Player1TeamID != 0 {
		player1 = &data.PlayerDescription{
			ID:     r.Player1ID,
			Name:   r.Player1Name,
			TeamID: r.Player1TeamID,
		}
	}

	if r.Player2ID != 0 && r.Player2Name != "" && r.Player2TeamID != 0 {
		player2 = &data.PlayerDescription{
			ID:     r.Player2ID,
			Name:   r.Player2Name,
			TeamID: r.Player2TeamID,
		}
	}

	if r.Player3ID != 0 && r.Player3Name != "" && r.Player3TeamID != 0 {
		player3 = &data.PlayerDescription{
			ID:     r.Player3ID,
			Name:   r.Player3Name,
			TeamID: r.Player3TeamID,
		}
	}

	event := &data.Event{
		GameID:             data.GameID(r.GameID),
		Period:             r.Period,
		Score:              r.Score(),
		PeriodTimeSeconds:  MinuteSecondStringToSeconds(r.PeriodClockTimeString),
		WallClockString:    r.WallClockTimeString,
		Player1:            player1,
		Player2:            player2,
		Player3:            player3,
		HomeDescription:    r.HomeDescription,
		NeutralDescription: r.NeutralDescription,
		VisitorDescription: r.VisitorDescription,
	}
	return event
}

func (r *PlayByPlayRow) DescriptionContains(substr string) bool {
	if r.HomeDescription != nil && strings.Contains(*r.HomeDescription, substr) {
		return true
	}
	if r.NeutralDescription != nil && strings.Contains(*r.NeutralDescription, substr) {
		return true
	}
	if r.VisitorDescription != nil && strings.Contains(*r.VisitorDescription, substr) {
		return true
	}
	return false
}

func (r *PlayByPlayRow) Score() *data.Score {
	if r.ScoreString == nil {
		return nil
	}

	pieces := strings.Split(*r.ScoreString, "-")
	home, err := strconv.Atoi(strings.TrimSpace(pieces[0]))
	if err != nil {
		return nil
	}

	visitor, err := strconv.Atoi(strings.TrimSpace(pieces[1]))
	if err != nil {
		return nil
	}

	return &data.Score{
		Home:    home,
		Visitor: visitor,
	}
}

func all(args ...*string) (res []string) {
	for _, a := range args {
		if a != nil {
			res = append(res, *a)
		}
	}
	return res
}
