package problems

import (
	"regexp"
	"strconv"

	common_functions "aoc.2023/lib/common/functions"
)

func solveChallenge() int {
	// Process the input
	scanner := common_functions.CreateInputScanner("../input.txt")
	defer scanner.File.Close()

	var answer = 0
	var matrix []string

	for scanner.Scan() {
		line := scanner.Text()

		matrix = append(matrix, line)
	}

	for i, v := range matrix {
		answer += computeLine(v, i, matrix)
	}

	return answer
}

func computeLine(input string, lineIndex int, matrix []string) int {
	numbersPattern := regexp.MustCompile(`([0-9]+)`)
	noSymbolsPattern := regexp.MustCompile(`([0-9\.])+`)

	var matches [][]int = numbersPattern.FindAllStringIndex(input, -1)
	var lineSum int = 0

	var matrixRowsLimit = len(matrix) - 1
	var matrixColsLimit = len(matrix[0]) - 1

	for _, matchIndex := range matches {

		for i := matchIndex[0]; i < matchIndex[1]; i++ {
			// Evaluate only rows
			if lineIndex > 0 && !noSymbolsPattern.MatchString(string(matrix[lineIndex-1][i])) ||
				lineIndex < matrixRowsLimit && !noSymbolsPattern.MatchString(string(matrix[lineIndex+1][i])) ||
				// Evaluate only columns
				i > 0 && !noSymbolsPattern.MatchString(string(matrix[lineIndex][i-1])) ||
				i < matrixColsLimit && !noSymbolsPattern.MatchString(string(matrix[lineIndex][i+1])) ||
				// Evaluate only diagonals
				lineIndex > 0 && i > 0 && !noSymbolsPattern.MatchString(string(matrix[lineIndex-1][i-1])) ||
				lineIndex > 0 && i < matrixColsLimit && !noSymbolsPattern.MatchString(string(matrix[lineIndex-1][i+1])) ||
				lineIndex < matrixRowsLimit && i > 0 && !noSymbolsPattern.MatchString(string(matrix[lineIndex+1][i-1])) ||
				lineIndex < matrixRowsLimit && i < matrixColsLimit && !noSymbolsPattern.MatchString(string(matrix[lineIndex+1][i+1])) {

				numberStrTmp := matrix[lineIndex][matchIndex[0]:matchIndex[1]]
				lineSum += getIntegerByString(numberStrTmp)

				break
			}

		}

	}

	return lineSum
}

func getIntegerByString(numberStrTmp string) int {
	numberIntTmp, err := strconv.Atoi(numberStrTmp)
	if err != nil {
		panic(err)
	}

	return numberIntTmp
}
