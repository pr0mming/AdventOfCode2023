package problems_13_2

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

	var (
		answer int = 0
		input  []string
	)

	for {
		continueFile := scanner.Scan()
		line := scanner.Text()

		// If we have a blank space, then we process the current input
		// Or if we're in last input
		if len(line) == 0 || !continueFile {

			answer += processInputPattern(input)
			input = nil

		} else {
			// Collect the rows of the current input
			input = append(input, line)
		}

		if !continueFile {
			break
		}
	}

	return strconv.Itoa(answer)
}

func processInputPattern(input []string) int {
	// First reflected columns
	var reflectionsResult = findVerticalReflections(input)

	// If is > -1 then we have reflected columns!
	if reflectionsResult > -1 {
		return reflectionsResult
	}

	// First reflected rows
	reflectionsResult = findHorizontalReflections(input)

	if reflectionsResult > -1 {
		return reflectionsResult * 100
	}

	// There isn't reflected columns and rows (it shouldn't happen)
	return 0
}

func findVerticalReflections(input []string) int {
	var newInput = make([]string, len(input[0]))

	// This is to convert columns to rows
	// A B  =>  A C
	// C D      B D
	// To avoid compare char by char

	for _, row := range input {
		for i := 0; i < len(row); i++ {
			newInput[i] += string(row[i])
		}
	}

	// Reuse the same logic for horizontal reflections
	return findHorizontalReflections(newInput)
}

func findHorizontalReflections(input []string) int {
	for i := 1; i < len(input); i++ {

		reflectionResult := checkPerfectReflection(input, i)

		// Try the next row ...
		// But if we have a reflected row then return the value
		if reflectionResult {
			return i
		}

	}

	return -1
}

func checkPerfectReflection(input []string, reflectedIndex int) bool {
	smudgeChecked := false // We can change a # to a . (and viceversa) just once!

	// Iterate from the reflectedIndex (+1 and -1) value, to verify if it'is a perfect reflection (until edges)
	for j, k := reflectedIndex, reflectedIndex-1; j < len(input) && k >= 0; j, k = j+1, k-1 {
		if input[j] != input[k] {
			// If we haven't fixed AND the row has a smudge
			// then we change the var value (because it's possible just once!)
			if !smudgeChecked && hasSmudge(input[j], input[k]) {
				smudgeChecked = true

				continue
			}

			// Id both rows doesn't have reflection and doesn't have smudge then there is not reflection
			return false
		}
	}

	// Now we have to check if we change a char in the row
	return smudgeChecked
}

func hasSmudge(a, b string) bool {
	diffChar := 0

	// Compare char by char of the row, if we have just one different (at the same index order) it has a smudge!
	for i := 0; i < len(a); i++ {
		if a[i] != b[i] {
			diffChar++
		}
	}

	if diffChar == 1 {
		return true
	}

	return false
}
