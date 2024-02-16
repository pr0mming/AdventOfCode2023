package problems_14_1

import (
	"fmt"
	"slices"
	"strconv"

	common_functions "aoc.2023/lib/common/functions"
	common_types "aoc.2023/lib/common/types"
	common_problem_types "aoc.2023/lib/common/types/problems/14"
)

var (
	ROUNDED_ROCK_FLAG     byte = []byte("O")[0]
	CUBE_SHAPED_ROCK_FLAG byte = []byte("#")[0]
)

func SolveChallenge(problemId string) string {
	// Process the input
	inputFilePath := fmt.Sprintf("problems/%s/input.txt", problemId)
	scanner := common_functions.CreateInputScanner(inputFilePath)
	defer scanner.File.Close()

	problemParameters := processPlatformInput(*scanner)

	var (
		totalRows      int = problemParameters.TotalRows      // Rows of the input to calculate the total load per each moved rock
		totalLoads     int = problemParameters.TotalLoads     // Calculated score using the the first row
		rocksStack         = problemParameters.RocksStack     // Stack of the rock's positions
		blockPointsMap     = problemParameters.BlockPointsMap // Map of the blocked points to check how far can I move a rock
	)

	for !rocksStack.IsEmpty() {
		rockPos, hasRock := rocksStack.Pop()

		if hasRock {
			var (
				i int = rockPos[0] // Row
				j int = rockPos[1] // Column
			)

			// If the column of the rock is this map then we need to calculate how far can the rock get
			if _, ok := blockPointsMap[j]; ok {

				nearRowPos := getNewRockRow(blockPointsMap[j], i, j)
				totalLoads += (totalRows - nearRowPos) // Calculate new score

				// Add the new position of the rock to the map
				addPointToMap(&blockPointsMap, j, nearRowPos)

			} else {
				// If the column of the rock isn't in the map, then we can put the rock in the row 0
				totalLoads += totalRows

				addPointToMap(&blockPointsMap, j, 0)
			}
		}
	}

	answer := totalLoads

	return strconv.Itoa(answer)
}

func processPlatformInput(scanner common_types.FileInputScanner) common_problem_types.PlatformParameters {
	var (
		i              int = 0
		totalLoads     int = 0
		rocksStack     common_types.Stack[[2]int]
		blockPointsMap = make(map[int][]int)
	)

	// The rocks in the first row can't move, so we take in account the ammount of rocks in this row
	// And we add those positions as blocked points
	if scanner.Scan() {
		line := scanner.Text()

		for j := 0; j < len(line); j++ {
			switch line[j] {

			case ROUNDED_ROCK_FLAG:
				addPointToMap(&blockPointsMap, j, i)
				totalLoads++
			case
				CUBE_SHAPED_ROCK_FLAG:
				addPointToMap(&blockPointsMap, j, i)
			}
		}

		i++
	}

	// Process from the row 1
	for scanner.Scan() {
		line := scanner.Text()

		for j := 0; j < len(line); j++ {
			switch line[j] {

			case ROUNDED_ROCK_FLAG:
				rocksStack.Push([2]int{i, j})

			case CUBE_SHAPED_ROCK_FLAG:
				addPointToMap(&blockPointsMap, j, i)
			}
		}

		i++
	}

	return common_problem_types.PlatformParameters{
		TotalRows:      i,
		TotalLoads:     totalLoads * i,
		RocksStack:     rocksStack,
		BlockPointsMap: blockPointsMap,
	}
}

func getNewRockRow(rowsPos []int, i, j int) int {
	nearBlockPos := 0 // It's the highest row position can get a rock

	// Iterate blocked rows, we try to put the rock in the position of blocked point + 1
	for _, v := range rowsPos {
		if v < i {
			nearBlockPos = (v + 1)
		} else {
			break
		}
	}

	return nearBlockPos
}

func addPointToMap(mapInput *map[int][]int, key, value int) {
	if _, ok := (*mapInput)[key]; ok {

		valTmp := (*mapInput)[key]
		valTmp = append(valTmp, value)

		slices.Sort(valTmp)

		(*mapInput)[key] = valTmp

	} else {
		(*mapInput)[key] = []int{value}
	}
}
