package problems

import (
	"regexp"
	"strconv"

	common_functions "aoc.2023/lib/common/functions"
)

func solveChallenge() uint64 {
	// Process the input
	scanner := common_functions.CreateInputScanner("../input.txt")
	defer scanner.File.Close()

	var answer uint64
	var matrix []string
	var numbersMatrix [][][]int

	numbersPattern := regexp.MustCompile(`([0-9])+`)

	for scanner.Scan() {
		line := scanner.Text()

		matrix = append(matrix, line)
		numbersMatrix = append(numbersMatrix, numbersPattern.FindAllStringIndex(line, -1))
	}

	for i, v := range matrix {
		resultTmp := computeLine(v, i, matrix, numbersMatrix)
		answer += resultTmp
	}

	return answer
}

func computeLine(input string, lineIndex int, matrix []string, numbersMatrix [][][]int) uint64 {
	symbolsPattern := regexp.MustCompile(`[^0-9.]`)

	var totalSum uint64 = 0
	var matches [][]int = symbolsPattern.FindAllStringIndex(input, -1)
	var matrixRowsLimit = len(matrix) - 1

	for _, matchIndex := range matches {
		var lineProd uint64 = 1

		i := matchIndex[0]
		j := matchIndex[1]

		var resultsTmp []uint64

		// Evaluate only up and down
		if lineIndex > 0 {
			resultsTmp = append(resultsTmp, checkNearNumber(i, j, matrix[lineIndex-1], numbersMatrix[lineIndex-1])...)
		}

		if lineIndex < matrixRowsLimit && len(resultsTmp) < 2 {
			resultsTmp = append(resultsTmp, checkNearNumber(i, j, matrix[lineIndex+1], numbersMatrix[lineIndex+1])...)
		}

		// Evaluate only right and left
		if i > 0 && len(resultsTmp) < 2 {
			resultsTmp = append(resultsTmp, checkNearNumber(i, j, matrix[lineIndex], numbersMatrix[lineIndex])...)
		}

		if len(resultsTmp) == 2 {
			for _, v := range resultsTmp {
				lineProd *= v
			}

			totalSum += lineProd
		}
	}

	return totalSum
}

func checkNearNumber(i int, j int, line string, numbers [][]int) []uint64 {
	var adjNumbers []uint64

	for _, pos := range numbers {
		if pos[0] <= i && pos[1] >= i || pos[0] <= j && pos[1] >= j {
			adjNumbers = append(adjNumbers, getIntegerByString(line[pos[0]:pos[1]]))
		}
	}

	return adjNumbers
}

func getIntegerByString(numberStrTmp string) uint64 {
	numberIntTmp, err := strconv.ParseUint(numberStrTmp, 10, 64)
	if err != nil {
		panic(err)
	}

	return numberIntTmp
}
