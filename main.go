package main

import (
	"bufio"
	"fmt"
	"log"
	match "match-making/match"
	"os"
)

type GameType = match.GameType

func main() {
	fmt.Println("Starting Game simulation")
	file, err := os.Open("1v1.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		fmt.Println(scanner.Text())
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	g := GameType{Name: "1v1", PlayerPot: 2}

	fmt.Println(g)

}
