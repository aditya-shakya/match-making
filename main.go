package main

import (
	"bufio"
	"fmt"
	"log"
	match "match-making/match"
	"os"
	"strconv"
	"strings"
)

type GameType = match.GameType
type GameRequest = match.GameRequest

func main() {
	fmt.Println("Starting Game simulation")

	coordinator := match.InitCoordinator()
	fmt.Println(coordinator)
	fmt.Println("Coordinator started")

	g := GameType{Name: "1v1", PlayerPot: 2}
	file, err := os.Open("1v1.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		s := scanner.Text()
		splitted := strings.Split(s, " ")
		i0, err0 := strconv.Atoi(splitted[0])
		i1, err1 := strconv.Atoi(splitted[1])
		if err0 == nil && err1 == nil {
			gr := match.NewRequest(i0, g, i1, true)
			coordinator.Add(gr)
		}
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	fmt.Println(coordinator)

}
