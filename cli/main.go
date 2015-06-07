package main

import (
	"fmt"
	"time"

	"github.com/jbowens/nbagame"
)

func printTodaysGames() {
	games, err := nbagame.API.Games.ByDate(time.Now())
	if err != nil {
		panic("Retrieving games error: " + err.Error())
	}

	for _, game := range games {
		fmt.Printf("Game %v between %v and %v, %v\n", game.ID, game.HomeTeamID, game.VisitorTeamID, game.Status)
	}
}

func printGameDetails(gameID string) {
	details, err := nbagame.API.Games.Details(gameID)
	if err != nil {
		panic("Game details error: " + err.Error())
	}

	fmt.Printf("%+v\n", details)
}

func printBoxScore(gameID string) {
	boxScore, err := nbagame.API.Games.BoxScore(gameID)
	if err != nil {
		panic("Box Score error: " + err.Error())
	}

	for _, teamStats := range boxScore.TeamStats {
		fmt.Printf("%s %s - %v\n", teamStats.TeamCity, teamStats.TeamName, teamStats.Points)
	}
}

func printPlayByPlay(gameID string) {
	events, err := nbagame.API.Games.PlayByPlay(gameID)
	if err != nil {
		panic("Play by play error: " + err.Error())
	}

	for _, event := range events {
		fmt.Printf("%+v\n", event)
	}
}

func printAllPlayers() {
	allPlayers, err := nbagame.API.Players.All()
	if err != nil {
		panic("All players error: " + err.Error())
	}

	for _, player := range allPlayers {
		fmt.Printf("%v %s %s\n", player.ID, player.FirstName, player.LastName)
	}
}

func printAllTeams() {
	allTeams, err := nbagame.API.Teams.All()
	if err != nil {
		panic("All teams error: " + err.Error())
	}

	for _, team := range allTeams {
		fmt.Printf("%v, %v, %v, %v\n", team.WinPercentage, team.ID, team.Name, team.City)
	}
}

func printPlayerDetails(playerID int) {
	playerDetails, err := nbagame.API.Players.Details(playerID)
	if err != nil {
		panic("Player Details error: " + err.Error())
	}

	fmt.Printf("%+v\n", playerDetails)
}

func main() {
	// printPlayerDetails(202322)
	// printBoxScore("0021401147")
	// printGameDetails("0021401147")
	// printPlayByPlay("0021401147")
	printTodaysGames()
}
