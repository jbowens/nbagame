package data

import "database/sql/driver"

// HomeOrAway indiciates whether a game was home or away with respect to a team or
// player.
type HomeOrAway bool

const (
	Home HomeOrAway = true
	Away HomeOrAway = false
)

func (h HomeOrAway) String() string {
	if h == Home {
		return "Home"
	} else {
		return "Away"
	}
}

func (h HomeOrAway) Value() (driver.Value, error) {
	return bool(h), nil
}

// Shot describes an individual shot within a game.
type Shot struct {
	ID                      int        `json:"id" db:"id"`
	GameID                  GameID     `json:"game_id" db:"game_id"`
	PlayerID                int        `json:"player_id" db:"player_id"`
	Number                  int        `json:"number" db:"shot_number"`
	Made                    bool       `json:"made" db:"made"`
	Points                  int        `json:"points" db:"points"`
	HomeOrAway              HomeOrAway `json:"home_or_away" db:"home"`
	Period                  int        `json:"period" db:"period"`
	GameClock               string     `json:"game_clock" db:"game_clock"`
	ShotClock               float64    `json:"shot_clock" db:"shot_clock"`
	Dribbles                int        `json:"dribbles" db:"dribbles"`
	TouchTimeSeconds        float64    `json:"touch_time_seconds" db:"touch_time_seconds"`
	Distance                float64    `json:"distance" db:"distance"`
	PointsType              int        `json:"points_type" db:"points_type"`
	ClosestDefender         int        `json:"closest_defender_player_id" db:"closest_defender_player_id"`
	ClosestDefenderDistance float64    `json:"closest_defender_distance" db:"closest_defender_distance"`
}
