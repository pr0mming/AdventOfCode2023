package main

import (
	"bufio"
	"fmt"
	"os"
	"path/filepath"
	"regexp"
	"strconv"
	"strings"
)

var NUMBERS_MAP = map[string]string{
	"one":   "1",
	"two":   "2",
	"three": "3",
	"four":  "4",
	"five":  "5",
	"six":   "6",
	"seven": "7",
	"eight": "8",
	"nine":  "9",
}

func main() {
	filePath := "../input.txt"

	scanner, err := createScanner(filePath)
	if err != nil {
		fmt.Println("Error opening the file:", err)
		return
	}
	defer scanner.file.Close()

	answer := solveChallenge(scanner)
	fmt.Println(answer)
}

type fileScanner struct {
	*bufio.Scanner
	file *os.File
}

func createScanner(filePath string) (*fileScanner, error) {
	absPath, _ := filepath.Abs(filePath)
	file, err := os.Open(absPath)
	if err != nil {
		return nil, err
	}

	return &fileScanner{
		Scanner: bufio.NewScanner(file),
		file:    file,
	}, nil
}

func solveChallenge(scanner *fileScanner) int {
	answer := 0

	numbersPattern := `(\d|one|two|three|four|five|six|seven|eight|nine)`

	// Compile the regular expression
	re := regexp.MustCompile(numbersPattern)

	for scanner.Scan() {
		line := scanner.Text()

		var matches []string

		for {
			numberMatchIndex := re.FindStringIndex(line)

			if numberMatchIndex == nil {
				break
			}

			matches = append(matches, line[numberMatchIndex[0]:numberMatchIndex[1]])

			line = line[numberMatchIndex[0]+1:]
		}

		var numberForLine [2]string

		numberForLine[0] = matches[0]

		if len(matches) > 1 {
			numberForLine[1] = matches[len(matches)-1]
		} else {
			numberForLine[1] = numberForLine[0]
		}

		answer += getNumberInt(numberForLine)
	}

	return answer
}

func getNumberInt(numberArr [2]string) int {
	for i, v := range numberArr {
		if len(v) > 1 {
			numberArr[i] = NUMBERS_MAP[v]
		}
	}

	numberInt, err := strconv.Atoi(strings.Join(numberArr[:], ""))
	if err != nil {
		panic(err)
	}

	//fmt.Println(numberArr, numberInt)

	return numberInt
}
