package db

import (
	"database/sql"

	"bitbucket.org/liamstask/goose/lib/goose"
	"github.com/jbowens/nbagame/data"
	"github.com/square/squalor"
	_ "github.com/ziutek/mymysql/native"
)

// DB encapsulates a connection to an NBAGame database.
type DB struct {
	DB        *squalor.DB
	Games     *squalor.Model
	Officials *squalor.Model
	Players   *squalor.Model
	Teams     *squalor.Model
}

// New creates a new connection to an NBAGame database. It takes
// an environment that should be defined in the goose dbconf.yml file.
func New(env string, confDirectory string) (*DB, error) {
	dbconf, err := goose.NewDBConf(confDirectory, env, "")
	if err != nil {
		return nil, err
	}

	mysqlDB, err := sql.Open(dbconf.Driver.Name, dbconf.Driver.OpenStr)
	if err != nil {
		return nil, err
	}

	db := &DB{
		DB: squalor.NewDB(mysqlDB),
	}
	if err := db.init(); err != nil {
		return nil, err
	}
	return db, err
}

func (db *DB) init() (err error) {
	db.Games, err = db.DB.BindModel("games", &data.GameDetails{})
	if err != nil {
		return err
	}
	db.Officials, err = db.DB.BindModel("officials", &data.Official{})
	if err != nil {
		return err
	}
	db.Players, err = db.DB.BindModel("players", &data.PlayerDetails{})
	if err != nil {
		return err
	}
	db.Teams, err = db.DB.BindModel("teams", &data.Team{})
	if err != nil {
		return err
	}

	return nil
}
