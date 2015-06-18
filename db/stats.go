package db

import (
	"database/sql"

	"github.com/jbowens/nbagame/data"
)

// RecordPlayerGameStats records stats about a player's performance in an individual
// game. It performs a lookup first to ensure that the player's stats haven't already
// been synced. If they have, it will update them.
func (db *DB) RecordPlayerGameStats(playerID int, gameID data.GameID, teamID int, stats *data.Stats) error {
	playerGameStats := data.PlayerGameStats{
		PlayerID: playerID,
		GameID:   gameID,
		TeamID:   teamID,
	}

	txn, err := db.DB.Begin()
	if err != nil {
		return err
	}
	defer txn.Rollback()

	// Get the existing stat line, if one exists.
	err = txn.Get(&playerGameStats, playerID, gameID)
	if err != nil && err != sql.ErrNoRows {
		return err
	}

	// If we've already seen this stats row, we just need to update it.
	if err == nil {
		stats.ID = playerGameStats.StatsID
		return txn.Replace(stats)
	}

	// Create a new stats line.
	if err := txn.Insert(stats); err != nil {
		return err
	}

	// Create the association between the player+game and the stats line.
	playerGameStats.StatsID = stats.ID
	if err := txn.Insert(&playerGameStats); err != nil {
		return err
	}

	return txn.Commit()
}

// RecordTeamGameStats records stats about a team's performance in an individual
// game. It performs a lookup first to ensure that the team's stats haven't already
// been synced. If they have, it will update them.
func (db *DB) RecordTeamGameStats(teamID int, gameID data.GameID, stats *data.Stats) error {
	teamGameStats := data.TeamGameStats{
		TeamID: teamID,
		GameID: gameID,
	}

	txn, err := db.DB.Begin()
	if err != nil {
		return err
	}
	defer txn.Rollback()

	// Get the existing stat line, if one exists.
	err = txn.Get(&teamGameStats, teamID, gameID)
	if err != nil && err != sql.ErrNoRows {
		return err
	}

	// If we've already seen this stats row, we just need to update it.
	if err == nil {
		stats.ID = teamGameStats.StatsID
		return txn.Replace(stats)
	}

	// Create a new stats line.
	if err := txn.Insert(stats); err != nil {
		return err
	}

	// Create the association between the team+game and the stats line.
	teamGameStats.StatsID = stats.ID
	if err := txn.Insert(&teamGameStats); err != nil {
		return err
	}

	return txn.Commit()
}
