package main

import (
	"bufio"
	"fmt"
	"os"
	"path/filepath"
	"regexp"
	"strconv"
)

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

func solveChallenge(scanner *fileScanner) int {
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
