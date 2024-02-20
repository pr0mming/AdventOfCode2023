package problems_15_1

import (
	"fmt"
	"strconv"
	"strings"

	common_functions "aoc.2023/lib/common/functions"
)

func SolveChallenge(problemId string) string {
	// Process the input
	inputFilePath := fmt.Sprintf("problems/%s/input.txt", problemId)
	scanner := common_functions.CreateInputScanner(inputFilePath)
	defer scanner.File.Close()

	var (
		answer int = 0
		steps  []string
	)

	// Because it's just one line
	if scanner.Scan() {
		line := scanner.Text()

		// Get inputs
		steps = strings.Split(line, ",")
	}

	answer = getHASHString(steps)

	return strconv.Itoa(answer)
}

func getHASHString(steps []string) int {
	sum := 0

	// Get sum for each input
	for _, step := range steps {
		sum += getHASHChar(step)
	}

	return sum
}

func getHASHChar(input string) int {
	output := 0

	// The same steps of the problem:
	for _, char := range input {
		output += int(char) // Increase the current value by the ASCII code you just determined.

		output *= 17  // Set the current value to itself multiplied by 17.
		output %= 256 // Set the current value to the remainder of dividing itself by 256.
	}

	return output
}
