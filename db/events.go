package db

import "github.com/jbowens/nbagame/data"

// Event is a db model for the data.Event type.
type Event struct {
	GameID             data.GameID `db:"game_id"`
	Seq                int         `db:"seq"`
	EventType          string      `db:"event_type"`
	Period             int         `db:"period"`
	ScoreHome          int         `db:"score_home"`
	ScoreVisitor       int         `db:"score_visitor"`
	PeriodTime         int         `db:"period_time"`
	WallClock          string      `db:"wall_clock"`
	Player1ID          *int        `db:"player1_id"`
	Player2ID          *int        `db:"player2_id"`
	Player3ID          *int        `db:"player3_id"`
	HomeDescription    *string     `db:"home_description"`
	NeutralDescription *string     `db:"neutral_description"`
	VisitorDescription *string     `db:"visitor_description"`
}

func createEventModel(evt *data.Event) *Event {
	ret := &Event{
		GameID:             evt.GameID,
		Seq:                evt.Number,
		EventType:          evt.Type.String(),
		Period:             evt.Period,
		WallClock:          evt.WallClockString,
		HomeDescription:    evt.HomeDescription,
		NeutralDescription: evt.NeutralDescription,
		VisitorDescription: evt.VisitorDescription,
	}
	if evt.Score != nil {
		ret.ScoreHome = evt.Score.Home
		ret.ScoreVisitor = evt.Score.Visitor
	}
	if evt.Player1 != nil {
		ret.Player1ID = &evt.Player1.ID
	}
	if evt.Player2 != nil {
		ret.Player2ID = &evt.Player2.ID
	}
	if evt.Player3 != nil {
		ret.Player3ID = &evt.Player3.ID
	}
	return ret
}

func (db *DB) RecordGameEvent(gameID data.GameID, evt *data.Event) error {
	model := createEventModel(evt)

	txn, err := db.DB.Begin()
	if err != nil {
		return err
	}
	defer txn.Rollback()

	err = txn.Upsert(model)
	if err != nil {
		return err
	}

	return txn.Commit()
}
