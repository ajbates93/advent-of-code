package main

import (
	"advent-of-code/utils"
	"fmt"
	"os"
	"strconv"
	"strings"
)

var REQUESTED_COLOURS = map[string]int{"red": 12, "green": 13, "blue": 14}

// Helpers -----------

func isPossible(gameTotals, colours map[string]int) bool {
	for colour, amount := range colours {
		if gameTotals[colour] < amount {
			return true
		} else {
			break
		}
	}
	return false
}

// Calculations --------

func parseGame(game string) map[string]int {
	totals := make(map[string]int)
	formatted := strings.Split(game, ": ")
	for _, round := range strings.Split(formatted[1], "; ") {
		for _, pack := range strings.Split(round, ", ") {
			amount, colour := parsePack(pack)
			totals[colour] += amount
		}
	}
	return totals
}

func parsePack(pack string) (int, string) {
	split := strings.Split(pack, " ")
	amount, _ := strconv.Atoi(split[0])
	colour := split[1]
	return amount, colour
}

func main() {
	fileName := os.Args[1]
	lines, err := utils.ReadFileByLine(fileName)
	// var staticLines = []string{
	// 	"Game 1: 3 blue, 4 red; 1 red, 2 green, 6 blue; 2 green",
	// 	"Game 2: 1 blue, 2 green; 3 green, 4 blue, 1 red; 1 green, 1 blue",
	// 	"Game 3: 8 green, 6 blue, 20 red; 5 blue, 4 red, 13 green; 5 green, 1 red",
	// 	"Game 4: 1 green, 3 red, 6 blue; 3 green, 6 red; 3 green, 15 blue, 14 red",
	// 	"Game 5: 6 red, 1 blue, 3 green; 2 blue, 1 red, 2 green",
	// }

	var possibleGames []int
	sum := 0
	for i, game := range lines {
		totals := parseGame(game)
		if isPossible(totals, REQUESTED_COLOURS) {
			possibleGames = append(possibleGames, i+1)
			sum += i + 1
		}
	}

	fmt.Printf("Possible games: %d\n", possibleGames)
	fmt.Printf("Sum of game IDs: %d\n", sum)

	if err != nil {
		fmt.Println("Error reading file")
	}
}
