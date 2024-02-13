package problems_3_1

import (
	"fmt"
	"regexp"
	"strconv"

	common_functions "aoc.2023/lib/common/functions"
)

// Use global variables to avoid reinitialized inside the loop
var (
	NUMBERS_PATTERN    = regexp.MustCompile(`([0-9]+)`)
	NO_SYMBOLS_PATTERN = regexp.MustCompile(`([0-9\.])+`)
	MATRIX_ROWS_LIMIT  = 0
	MATRIX_COLS_LIMIT  = 0
)

func SolveChallenge(problemId string) string {
	// Process the input
	inputFilePath := fmt.Sprintf("problems/%s/input.txt", problemId)
	scanner := common_functions.CreateInputScanner(inputFilePath)
	defer scanner.File.Close()

	var (
		answer = 0
		matrix []string
	)

	for scanner.Scan() {
		line := scanner.Text()

		matrix = append(matrix, line)
	}

	// We use these limits to know if we can move on the iteration across the matrix
	MATRIX_ROWS_LIMIT = len(matrix) - 1
	MATRIX_COLS_LIMIT = len(matrix[0]) - 1

	for i, v := range matrix {
		answer += computeLine(v, i, matrix)
	}

	return strconv.Itoa(answer)
}

func computeLine(input string, lineIndex int, matrix []string) int {
	// Detect all the numbers in the current line
	var (
		matches [][]int = NUMBERS_PATTERN.FindAllStringIndex(input, -1)
		lineSum int     = 0
	)

	for _, matchIndex := range matches {
		for i := matchIndex[0]; i < matchIndex[1]; i++ {
			// Evaluate only rows ↑↓
			if lineIndex > 0 && !NO_SYMBOLS_PATTERN.MatchString(string(matrix[lineIndex-1][i])) ||
				lineIndex < MATRIX_ROWS_LIMIT && !NO_SYMBOLS_PATTERN.MatchString(string(matrix[lineIndex+1][i])) ||
				// Evaluate only columns ← →
				i > 0 && !NO_SYMBOLS_PATTERN.MatchString(string(matrix[lineIndex][i-1])) ||
				i < MATRIX_COLS_LIMIT && !NO_SYMBOLS_PATTERN.MatchString(string(matrix[lineIndex][i+1])) ||
				// Evaluate only diagonals
				lineIndex > 0 && i > 0 && !NO_SYMBOLS_PATTERN.MatchString(string(matrix[lineIndex-1][i-1])) ||
				lineIndex > 0 && i < MATRIX_COLS_LIMIT && !NO_SYMBOLS_PATTERN.MatchString(string(matrix[lineIndex-1][i+1])) ||
				lineIndex < MATRIX_ROWS_LIMIT && i > 0 && !NO_SYMBOLS_PATTERN.MatchString(string(matrix[lineIndex+1][i-1])) ||
				lineIndex < MATRIX_ROWS_LIMIT && i < MATRIX_COLS_LIMIT && !NO_SYMBOLS_PATTERN.MatchString(string(matrix[lineIndex+1][i+1])) {

				numberStrTmp := matrix[lineIndex][matchIndex[0]:matchIndex[1]]
				lineSum += common_functions.Atoi(numberStrTmp)

				break
			}
		}
	}

	return lineSum
}
