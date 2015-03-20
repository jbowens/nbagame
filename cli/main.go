package main

import (
	"fmt"

	"github.com/jbowens/nbagame"
)

func main() {
	allPlayers, err := nbagame.API.AllPlayers()
	if err != nil {
		panic("All players error: " + err.Error())
	}

	for _, player := range allPlayers {
		fmt.Printf("%v %s %s\n", player.ID, player.FirstName, player.LastName)
	}
}
