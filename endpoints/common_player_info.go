package endpoints

import (
	"strconv"
	"strings"
	"time"

	"github.com/jbowens/nbagame/data"
)

const (
	dateFormat = "2006-01-02T15:04:05"
)

// CommonPlayerInfoParams defines parameters for a CommonPlayerInfo request.
type CommonPlayerInfoParams struct {
	LeagueID string `json:"LeagueID"`
	PlayerID int    `json:"PlayerID"`
}

// CommonPlayerInfoResponse is the type for all result sets returned by the
// 'commonplayerinfo' resource.
type CommonPlayerInfoResponse struct {
	CommonPlayerInfo []*CommonPlayerInfoRow `nbagame:"CommonPlayerInfo"`
}

// CommonPlayerInfoRow represents the schema returned for 'CommonPlayerInfo'
// result sets, from the 'commonplayerinfo' resource.
//
// Example URL:
// http://stats.nba.com/stats/commonplayerinfo?LeagueID=00&PlayerID=201566&SeasonType=Regular+Season
type CommonPlayerInfoRow struct {
	PlayerID         int    `nbagame:"PERSON_ID"`
	FirstName        string `nbagame:"FIRST_NAME"`
	LastName         string `nbagame:"LAST_NAME"`
	Birthdate        string `nbagame:"BIRTHDATE"`
	School           string `nbagame:"SCHOOL"`
	Country          string `nbagame:"COUNTRY"`
	LastAffiliation  string `nbagame:"LAST_AFFILIATION"`
	Height           string `nbagame:"HEIGHT"`
	Weight           string `nbagame:"WEIGHT"`
	SeasonExperience int    `nbagame:"SEASON_EXP"`
	Jersey           string `nbagame:"JERSEY"`
	Position         string `nbagame:"POSITION"`
	RosterStatus     string `nbagame:"ROSTERSTATUS"`
	TeamID           int    `nbagame:"TEAM_ID"`
	TeamName         string `nbagame:"TEAM_NAME"`
	TeamAbbreviation string `nbagame:"TEAM_ABBREVIATION"`
	TeamCity         string `nbagame:"TEAM_CITY"`
	FromYear         int    `nbagame:"FROM_YEAR"`
	ToYear           int    `nbagame:"TO_YEAR"`
	DLeagueFlag      string `nbagame:"DLEAGUE_FLAG"`
}

// ToPlayerDetails converts a row to a PlayerDetails struct.
func (r *CommonPlayerInfoRow) ToPlayerDetails() (*data.PlayerDetails, error) {
	playerDetails := &data.PlayerDetails{
		PlayerID:         r.PlayerID,
		FirstName:        r.FirstName,
		LastName:         r.LastName,
		School:           r.School,
		Country:          r.Country,
		SeasonExperience: r.SeasonExperience,
		Jersey:           strings.TrimSpace(r.Jersey),
		Position:         strings.TrimSpace(r.Position),
		TeamID:           r.TeamID,
		TeamName:         r.TeamName,
		TeamAbbreviation: r.TeamAbbreviation,
		TeamCity:         r.TeamCity,
		CareerStartYear:  strconv.Itoa(r.FromYear),
		CareerEndYear:    strconv.Itoa(r.ToYear),
	}

	// Convert birthdate string into a time.Time
	bday, err := time.Parse(dateFormat, r.Birthdate)
	if err != nil {
		return nil, err
	}
	playerDetails.Birthdate = &bday

	// Convert height from string into inches integer.
	heightParts := strings.Split(r.Height, "-")
	if len(heightParts) == 2 {
		feet, _ := strconv.Atoi(heightParts[0])
		inches, _ := strconv.Atoi(heightParts[1])
		playerDetails.Height = feet*12 + inches
	}

	// Convert weight into an integer.
	playerDetails.Weight, _ = strconv.Atoi(r.Weight)

	// Convert Roster Status string to enum
	if r.RosterStatus == "Active" {
		playerDetails.RosterStatus = data.Active
	}

	// Convert DLeague flag to boolean
	if r.DLeagueFlag != "N" {
		playerDetails.DLeague = true
	}

	return playerDetails, nil
}
