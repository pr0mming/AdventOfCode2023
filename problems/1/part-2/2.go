package problems_1_2

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"

	common_functions "aoc.2023/lib/common/functions"
)

// This constant is used for getting the right equivalence
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

func SolveChallenge(problemId string) string {
	// Process the input
	inputFilePath := fmt.Sprintf("problems/%s/input.txt", problemId)
	scanner := common_functions.CreateInputScanner(inputFilePath)
	defer scanner.File.Close()

	// We prepare a Regex approach to detect the written numbers or the [0-9] range
	answer := 0
	numbersPattern := `(\d|one|two|three|four|five|six|seven|eight|nine)`
	re := regexp.MustCompile(numbersPattern)

	for scanner.Scan() {
		line := scanner.Text()

		var matches []string

		// We iterate the input until find all the occurrences
		// it's due I wasn't able to implement the occurrences overlaped using only the regex
		for {
			numberMatchIndex := re.FindStringIndex(line)

			if numberMatchIndex == nil {
				break
			}

			matches = append(matches, line[numberMatchIndex[0]:numberMatchIndex[1]])

			// We cut the current input from the last position of the occurrence to recalculate again...
			line = line[numberMatchIndex[0]+1:]
		}

		var numberForLine [2]string

		// How it's supposed to get always at least 1 entry we set this array with that value
		numberForLine[0] = matches[0]

		// If there is only 1 entry then we set the 2th position of the array with the 1th position
		// Otherwise we choose the last occurrence
		if len(matches) > 1 {
			numberForLine[1] = matches[len(matches)-1]
		} else {
			numberForLine[1] = numberForLine[0]
		}

		answer += getNumberInt(numberForLine)
	}

	return strconv.Itoa(answer)
}

func getNumberInt(numberArr [2]string) int {
	for i, v := range numberArr {
		// For example, "one" is a valid string for this condition, so that we extract the number equivalence form the map
		// Otherwise we keep the number
		if len(v) > 1 {
			numberArr[i] = NUMBERS_MAP[v]
		}
	}

	// We join the array and convert it to integer
	numberInt, err := strconv.Atoi(strings.Join(numberArr[:], ""))
	if err != nil {
		panic(err)
	}

	return numberInt
}
