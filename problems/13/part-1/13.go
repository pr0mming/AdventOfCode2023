package problems_13_1

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
		input  []string
	)

	for {
		continueFile := scanner.Scan()
		line := scanner.Text()

		if len(strings.TrimSpace(line)) == 0 || !continueFile {

			answer += processInputPattern(input)
			input = nil

		} else {
			input = append(input, line)
		}

		if !continueFile {
			break
		}
	}

	return strconv.Itoa(answer)
}

func processInputPattern(input []string) int {
	verticalReflections := findVerticalReflections(input)

	if verticalReflections == 0 {
		return findHorizontalReflections(input)
	}

	return verticalReflections * 100
}

func findVerticalReflections(input []string) int {
	var newInput = make([]string, len(input[0]))

	for _, row := range input {
		for i := 0; i < len(row); i++ {
			newInput[i] = string(row[i]) + newInput[i]
		}
	}

	return findHorizontalReflections(newInput)
}

func findHorizontalReflections(input []string) int {
	var (
		matches         = make(map[string][]int)
		reflections int = 0
	)

	for i, row := range input {

		if _, ok := matches[row]; ok {
			indexes := matches[row]

			indexes = append(indexes, i)
			matches[row] = indexes
		} else {
			matches[row] = []int{i}
		}

	}

	for _, v := range matches {
		if len(v) == 2 {
			if v[1]-v[0] == 1 {
				break
			}
		}

		reflections++
	}

	if reflections > 0 {
		return (reflections / 2) + 1
	}

	return reflections
}
