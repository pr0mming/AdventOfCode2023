package problems_3_2

import (
	"fmt"
	"regexp"
	"strconv"

	common_functions "aoc.2023/lib/common/functions"
)

// Use global variables to avoid reinitialized inside the loop
var NUMBERS_PATTERN = regexp.MustCompile(`([0-9]+)`)
var SYMBOLS_PATTERN = regexp.MustCompile(`[^0-9.]`)
var MATRIX_ROWS_LIMIT = 0

func SolveChallenge(problemId string) string {
	// Process the input
	inputFilePath := fmt.Sprintf("problems/%s/input.txt", problemId)
	scanner := common_functions.CreateInputScanner(inputFilePath)
	defer scanner.File.Close()

	var answer uint64
	var matrix []string
	var numbersMatrix [][][]int

	for scanner.Scan() {
		line := scanner.Text()

		matrix = append(matrix, line)
		numbersMatrix = append(numbersMatrix, NUMBERS_PATTERN.FindAllStringIndex(line, -1))
	}

	// We use these limits to know if we can move on the iteration across the matrix
	MATRIX_ROWS_LIMIT = len(matrix) - 1

	for i, v := range matrix {
		resultTmp := computeLine(v, i, matrix, numbersMatrix)
		answer += resultTmp
	}

	return strconv.FormatUint(answer, 10)
}

func computeLine(input string, lineIndex int, matrix []string, numbersMatrix [][][]int) uint64 {
	var totalSum uint64 = 0
	// Extract all the allowed symbols (index start and end) to iterate
	var matches [][]int = SYMBOLS_PATTERN.FindAllStringIndex(input, -1)

	for _, matchIndex := range matches {
		var lineProd uint64 = 1

		i := matchIndex[0]
		j := matchIndex[1]

		var resultsTmp []uint64

		// Evaluate only up and down ↑↓ (we use -1 and +1 to change the row to compare)
		if lineIndex > 0 {
			resultsTmp = append(resultsTmp, checkAdjNumber(i, j, matrix[lineIndex-1], numbersMatrix[lineIndex-1])...)
		}

		if lineIndex < MATRIX_ROWS_LIMIT && len(resultsTmp) < 2 {
			resultsTmp = append(resultsTmp, checkAdjNumber(i, j, matrix[lineIndex+1], numbersMatrix[lineIndex+1])...)
		}

		// Evaluate only right and left ← →
		if i > 0 && len(resultsTmp) < 2 {
			resultsTmp = append(resultsTmp, checkAdjNumber(i, j, matrix[lineIndex], numbersMatrix[lineIndex])...)
		}

		// If we have two adj numbers is enough (according to the challenge) to make the total
		if len(resultsTmp) == 2 {
			for _, v := range resultsTmp {
				lineProd *= v
			}

			totalSum += lineProd
		}
	}

	return totalSum
}

func checkAdjNumber(i int, j int, line string, numbers [][]int) []uint64 {
	var adjNumbers []uint64

	for _, pos := range numbers {
		// Check if the number coords is between a symbol coords
		if pos[0] <= i && pos[1] >= i || pos[0] <= j && pos[1] >= j {
			adjNumbers = append(adjNumbers, common_functions.ParseUint(line[pos[0]:pos[1]]))
		}
	}

	return adjNumbers
}
