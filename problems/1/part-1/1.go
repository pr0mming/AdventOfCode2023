package problems_1_1

import (
	"fmt"
	"strconv"

	common_functions "aoc.2023/lib/common/functions"
)

func SolveChallenge(problemId string) string {
	// Process the input
	inputFilePath := fmt.Sprintf("problems/%s/input.txt", problemId)
	scanner := common_functions.CreateInputScanner(inputFilePath)
	defer scanner.File.Close()

	answer := 0

	for scanner.Scan() {
		line := scanner.Text()
		var numberForLine string

		// We iterate the input until find the first number
		for i := 0; i < len(line); i++ {
			charTmp := line[i]

			// For each char we compare it if corresponds to a valid ASCII char (0-9)
			if charTmp >= 48 && charTmp <= 57 {
				numberForLine += string(charTmp)
				break
			}
		}

		// We iterate the input (backwards) until find the last number
		for i := len(line) - 1; i >= 0; i-- {
			charTmp := line[i]

			if charTmp >= 48 && charTmp <= 57 {
				numberForLine += string(charTmp)
				break
			}
		}

		// We convert the expression to number to sum it to the answer
		intTmp, err := strconv.Atoi(numberForLine)
		if err != nil {
			panic(err)
		}

		answer += intTmp
	}

	return strconv.Itoa(answer)
}
