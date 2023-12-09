package utils

import (
	"bufio"
	"os"
)

func ReadFile(fileName string) (string, error) {
	data, err := os.ReadFile(fileName)
	if err != nil {
		return "", err
	}
	return string(data), nil
}

func ReadFileByLine(fileName string) ([]string, error) {
	file, err := os.Open(fileName)
	defer file.Close()
	if err != nil {
		return nil, err
	}

	scanner := bufio.NewScanner(file)
	var lines []string
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	return lines, scanner.Err()
}

func ReadFileByLineBytes(fileName string) ([][]byte, error) {
	lines, err := ReadFileByLine(fileName)
	if err != nil {
		return nil, err
	}

	byteArrLines := make([][]byte, len(lines))
	for i := 0; i < len(lines); i++ {
		byteArrLines[i] = []byte(lines[i])
	}

	return byteArrLines, nil
}
