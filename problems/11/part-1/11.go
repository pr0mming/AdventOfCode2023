package problems_11_1

import (
	"fmt"
	"strconv"

	common_functions "aoc.2023/lib/common/functions"
	common_types "aoc.2023/lib/common/types"
)

var GALAXY_CHAR = []byte("#")[0]

func SolveChallenge(problemId string) string {
	// Process the input
	inputFilePath := fmt.Sprintf("problems/%s/input.txt", problemId)
	scanner := common_functions.CreateInputScanner(inputFilePath)
	defer scanner.File.Close()

	// galaxiesPos, is the list of positions of galaxies [X, Y]
	// nonExpandedRows, is a list to represent the rows of the map (like a map[int]bool) to check if the row is expanded
	// nonExpandedColumns, the same above but with columns
	galaxiesPos, nonExpandedRows, nonExpandedColumns := processImageInput(*scanner)
	answer := 0

	// Loop to form the couples
	for i := 0; i < len(galaxiesPos); i++ {
		for j := i + 1; j < len(galaxiesPos); j++ {

			pointA := galaxiesPos[i]
			pointB := galaxiesPos[j]

			rowsSteps := getExpandedPaths(nonExpandedRows, pointA[0], pointB[0])
			colsSteps := getExpandedPaths(nonExpandedColumns, pointA[1], pointB[1])

			answer += rowsSteps + colsSteps
		}
	}

	return strconv.Itoa(answer)
}

func getExpandedPaths(paths []bool, a, b int) int {
	var (
		rowsOffset    int = 0
		expandedPaths int = 0
	)

	// Avoid get negative numbers
	if a > b {
		aTmp := a
		a = b
		b = aTmp
	}

	rowsOffset = b - a

	// Check extra steps (expanded paths for rows or columns)
	for i := a + 1; i < b; i++ {
		if !paths[i] {
			expandedPaths++
		}
	}

	return rowsOffset + expandedPaths
}

func processImageInput(scanner common_types.FileInputScanner) ([][2]int, []bool, []bool) {
	var (
		galaxiesPos [][2]int     // Is the list of positions of galaxies [X, Y]
		rows        int      = 0 // Used to save the number of rows of the map
		cols        int      = 0 // Used to save the number of columns of the map
	)

	for scanner.Scan() {
		line := scanner.Text()
		cols = len(line)

		// Loop to save the galaxies
		for j := 0; j < len(line); j++ {
			if line[j] == GALAXY_CHAR {
				galaxiesPos = append(galaxiesPos, [2]int{rows, j})
			}
		}

		rows++
	}

	// Now we make a representation for rows and columns like a map[int]bool
	// This is used to calculate the extra steps between galaxies
	var (
		nonExpandedRows    = make([]bool, rows)
		nonExpandedColumns = make([]bool, cols)
	)

	for _, galaxyPos := range galaxiesPos {
		row := galaxyPos[0]
		col := galaxyPos[1]

		nonExpandedRows[row] = true
		nonExpandedColumns[col] = true
	}

	return galaxiesPos, nonExpandedRows, nonExpandedColumns
}
