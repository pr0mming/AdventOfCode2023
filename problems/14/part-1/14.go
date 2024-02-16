package problems_14_1

import (
	"fmt"
	"slices"
	"strconv"

	common_functions "aoc.2023/lib/common/functions"
	common_types "aoc.2023/lib/common/types"
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

	totalLoads := 0
	rocksStack, blockPointsMap := processPlatformInput(*scanner)

	for !rocksStack.IsEmpty() {
		rockPos, hasRock := rocksStack.Pop()

		if hasRock {
			var (
				i int = rockPos[0]
				j int = rockPos[1]
			)

			if _, ok := blockPointsMap[j]; ok {

				nearRowPos := getNearBlockRow(blockPointsMap[j], i, j)
				totalLoads += (i - nearRowPos)

			} else {
				totalLoads += i
			}

			addPointToMap(&blockPointsMap, j, i)
		}
	}

	answer := totalLoads

	return strconv.Itoa(answer)
}

func processPlatformInput(scanner common_types.FileInputScanner) (common_types.Stack[[2]int], map[int][]int) {
	var (
		i              int = 0
		rocksStack     common_types.Stack[[2]int]
		blockPointsMap = make(map[int][]int)
	)

	if scanner.Scan() {
		line := scanner.Text()

		for j := 0; j < len(line); j++ {
			switch line[j] {

			case ROUNDED_ROCK_FLAG,
				CUBE_SHAPED_ROCK_FLAG:
				addPointToMap(&blockPointsMap, j, i)
			}
		}

		i++
	}

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

	return rocksStack, blockPointsMap
}

func getNearBlockRow(rowsPos []int, i, j int) int {
	nearBlockPos := 0

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
