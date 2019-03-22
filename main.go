package main

import (
	"fmt"
	match "match-making/match"
)

type GameType = match.GameType

func main() {
	fmt.Println("Starting Game simulation")
	g := GameType{Name: "1v1", PlayerPot: 2}
	fmt.Println(g)

}
