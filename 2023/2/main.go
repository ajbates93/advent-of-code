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
			return false
		}
	}
	return true
}

// Calculations --------

func parseGame(game string) map[string]int {
	totals := make(map[string]int)
	for _, round := range strings.Split(game, ";") {
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
