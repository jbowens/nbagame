package main

import (
	"fmt"

	"github.com/jbowens/nbagame"
)

func printAllPlayers() {
	allPlayers, err := nbagame.API.PlayersForCurrentSeason()
	if err != nil {
		panic("All players error: " + err.Error())
	}

	for _, player := range allPlayers {
		fmt.Printf("%v %s %s\n", player.ID, player.FirstName, player.LastName)
	}
}

func printAllTeams() {
	allTeams, err := nbagame.API.AllTeams()
	if err != nil {
		panic("All teams error: " + err.Error())
	}

	for _, team := range allTeams {
		fmt.Printf("%v, %v, %v, %v\n", team.WinPercentage, team.ID, team.Name, team.City)
	}
}

func main() {
	printAllTeams()
}
