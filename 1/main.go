package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strings"
)

var numberWords = []string{"one", "two", "three", "four", "five", "six", "seven", "eight", "nine", "ten"}

func stripNonNumericFromLine(line *string) {
	// Allow spelled out numbers

	fmt.Println("ORIGINAL LINE: ", *line)
	regex := regexp.MustCompile(`(?i)\b([a-z]+)\b`)
	matcher := regex.FindAllString(*line, -1)

	for _, word := range matcher {
		if !contains(numberWords, strings.ToLower(word)) {
			*line = strings.Replace(*line, word, "", -1)
		}
	}
	fmt.Println("Line after replace: ", *line)

	*line = regexp.MustCompile(`[^\d]`).ReplaceAllString(*line, "")

	fmt.Println("Final line: ", *line)
}

func contains(list []string, word string) bool {
	for _, v := range list {
		if v == word {
			return true
		}
	}
	return false
}

func readFromFile(fileName string) {
	file, err := os.Open(fileName)
	if err != nil {
		panic(err)
	}

	defer file.Close()

	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)

	for scanner.Scan() {
		line := scanner.Text()
		stripNonNumericFromLine(&line)
		fmt.Println(line)
	}

	if err := scanner.Err(); err != nil {
		panic(err)
	}
}

func main() {
	readFromFile("codes1.txt")
}
