package problems_14_2

import (
	"fmt"
	"slices"
	"sort"
	"strconv"

	common_functions "aoc.2023/lib/common/functions"
	common_types "aoc.2023/lib/common/types"
)

var (
	ROUNDED_ROCK_FLAG     byte = []byte("O")[0]
	CUBE_SHAPED_ROCK_FLAG byte = []byte("#")[0]
)

var (
	blockRowsMap    = make(map[int][]int)
	blockColumnsMap = make(map[int][]int)
)

var (
	MAX_ROW_INDEX    int
	MAX_COLUMN_INDEX int
)

const (
	ROLL_NORTH_DIRECTION = 0
	ROLL_WEST_DIRECTION  = 1
	ROLL_SOUTH_DIRECTION = 2
	ROLL_EAST_DIRECTION  = 3
)

func SolveChallenge(problemId string) string {
	// Process the input
	inputFilePath := fmt.Sprintf("problems/%s/input.txt", problemId)
	scanner := common_functions.CreateInputScanner(inputFilePath)
	defer scanner.File.Close()

	var (
		answer            int = 0
		totalLoadsHistory [][2]int
		rocksArr          = processPlatformInput(*scanner)
	)

	for {
		totalLoads := getSpinCycle(&rocksArr)
		historyIndex := slices.IndexFunc(totalLoadsHistory, func(item [2]int) bool {
			return item[0] == totalLoads[0] && item[1] == totalLoads[1]
		})

		if historyIndex == -1 {
			totalLoadsHistory = append(totalLoadsHistory, totalLoads)
		} else {
			break
		}
	}

	return strconv.Itoa(answer)
}

func processPlatformInput(scanner common_types.FileInputScanner) [][2]int {
	var (
		i        int = 0
		rocksArr [][2]int
	)

	// Process from the row 1
	for scanner.Scan() {
		line := scanner.Text()
		MAX_COLUMN_INDEX = len(line)

		for j := 0; j < len(line); j++ {
			switch line[j] {

			case ROUNDED_ROCK_FLAG:
				rocksArr = append(rocksArr, [2]int{i, j})
				addBlockedPos(i, j)

			case CUBE_SHAPED_ROCK_FLAG:
				addBlockedPos(i, j)
			}
		}

		i++
	}

	MAX_ROW_INDEX = i

	return rocksArr
}

func getSpinCycle(rocksArr *[][2]int) [2]int {
	var (
		totalRowsLoad int = 0
		totalColsLoad int = 0
	)

	for direction := 0; direction < 4; direction++ {

		switch direction {
		case ROLL_NORTH_DIRECTION, ROLL_WEST_DIRECTION:

			getNortAndhWestSpin(rocksArr, direction)

		case ROLL_SOUTH_DIRECTION:

			getSouthSpin(rocksArr, direction)

		case ROLL_EAST_DIRECTION:

			totalRowsLoad, totalColsLoad = getEastSpin(rocksArr, direction)

		}
	}

	return [2]int{totalRowsLoad, totalColsLoad}
}

func getNortAndhWestSpin(rocksArr *[][2]int, direction int) {
	sort.Slice(*rocksArr, func(i, j int) bool {
		if (*rocksArr)[i][0] != (*rocksArr)[j][0] {
			return (*rocksArr)[i][0] < (*rocksArr)[j][0]
		}

		return (*rocksArr)[i][1] < (*rocksArr)[j][1]
	})

	for i := 0; i < len(*rocksArr); i++ {
		var (
			row int = (*rocksArr)[i][0] // Row
			col int = (*rocksArr)[i][1] // Column
		)

		(*rocksArr)[i] = getNewCycleRockPosition(row, col, direction)
	}
}

func getSouthSpin(rocksArr *[][2]int, direction int) {
	sort.Slice(*rocksArr, func(i, j int) bool {
		if (*rocksArr)[i][0] != (*rocksArr)[j][0] {
			return (*rocksArr)[i][0] > (*rocksArr)[j][0]
		}

		return (*rocksArr)[i][1] > (*rocksArr)[j][1]
	})

	for i := 0; i < len(*rocksArr); i++ {
		var (
			row int = (*rocksArr)[i][0] // Row
			col int = (*rocksArr)[i][1] // Column
		)

		(*rocksArr)[i] = getNewCycleRockPosition(row, col, direction)
	}
}

func getEastSpin(rocksArr *[][2]int, direction int) (int, int) {
	sort.Slice(*rocksArr, func(i, j int) bool {
		if (*rocksArr)[i][0] != (*rocksArr)[j][0] {
			return (*rocksArr)[i][0] > (*rocksArr)[j][0]
		}

		return (*rocksArr)[i][1] > (*rocksArr)[j][1]
	})

	var (
		totalRowsLoad int = 0
		totalColsLoad int = 0
	)

	for i := 0; i < len(*rocksArr); i++ {
		var (
			row int = (*rocksArr)[i][0] // Row
			col int = (*rocksArr)[i][1] // Column
		)

		(*rocksArr)[i] = getNewCycleRockPosition(row, col, direction)

		totalRowsLoad += (MAX_ROW_INDEX) - (*rocksArr)[i][0]
		totalColsLoad += (MAX_COLUMN_INDEX) - (*rocksArr)[i][1]
	}

	return totalRowsLoad, totalColsLoad
}

func getNewCycleRockPosition(i, j, direction int) [2]int {
	switch direction {
	case ROLL_NORTH_DIRECTION:
		i = getNewNorthPos(i, j)

	case ROLL_WEST_DIRECTION:
		j = getNewWestPos(i, j)

	case ROLL_SOUTH_DIRECTION:
		i = getNewSouthPos(i, j)

	case ROLL_EAST_DIRECTION:
		j = getNewEastPos(i, j)
	}

	return [2]int{i, j}
}

func getNewNorthPos(row, col int) int {
	if _, ok := blockColumnsMap[col]; ok {
		nearBlockPos := 0 // It's the highest row position can get a rock

		// Iterate blocked rows, we try to put the rock in the position of blocked point + 1
		rowsPos := blockColumnsMap[col]

		for _, v := range rowsPos {
			if v < row {
				nearBlockPos = (v + 1)
			} else {
				break
			}
		}

		if row != nearBlockPos {
			deleteBlockedPos(row, col)
			addBlockedPos(nearBlockPos, col)
		}

		return nearBlockPos
	}

	return 0
}

func getNewWestPos(row, col int) int {
	if _, ok := blockRowsMap[row]; ok {
		nearBlockPos := 0 // It's the highest row position can get a rock

		// Iterate blocked rows, we try to put the rock in the position of blocked point + 1
		colsPos := blockRowsMap[row]

		for _, v := range colsPos {
			if v < col {
				nearBlockPos = (v + 1)
			} else {
				break
			}
		}

		if col != nearBlockPos {
			deleteBlockedPos(row, col)
			addBlockedPos(row, nearBlockPos)
		}

		return nearBlockPos
	}

	return 0
}

func getNewSouthPos(row, col int) int {
	if _, ok := blockColumnsMap[col]; ok {

		rowsPos := blockColumnsMap[col]
		nearBlockPos := MAX_ROW_INDEX - 1 // It's the highest row position can get a rock

		// Iterate blocked rows, we try to put the rock in the position of blocked point + 1
		for i := len(rowsPos) - 1; i >= 0; i-- {
			v := rowsPos[i]

			if v > row {
				nearBlockPos = (v - 1)
			} else {
				break
			}
		}

		if row != nearBlockPos {
			deleteBlockedPos(row, col)
			addBlockedPos(nearBlockPos, col)
		}

		return nearBlockPos
	}

	return MAX_ROW_INDEX - 1
}

func getNewEastPos(row, col int) int {
	if _, ok := blockRowsMap[row]; ok {

		colsPos := blockRowsMap[row]
		nearBlockPos := MAX_COLUMN_INDEX - 1 // It's the highest row position can get a rock

		// Iterate blocked rows, we try to put the rock in the position of blocked point + 1
		for i := len(colsPos) - 1; i >= 0; i-- {
			v := colsPos[i]

			if v > col {
				nearBlockPos = (v - 1)
			} else {
				break
			}
		}

		if col != nearBlockPos {
			deleteBlockedPos(row, col)
			addBlockedPos(row, nearBlockPos)
		}

		return nearBlockPos
	}

	return MAX_COLUMN_INDEX - 1
}

func addBlockedPos(row, column int) {
	addItemToMap(&blockRowsMap, row, column)
	addItemToMap(&blockColumnsMap, column, row)
}

func deleteBlockedPos(row, column int) {
	rows := blockColumnsMap[column]
	index := slices.Index(rows, row)

	blockColumnsMap[column] = append(rows[:index], rows[index+1:]...)

	columns := blockRowsMap[row]
	index = slices.Index(columns, column)

	blockRowsMap[row] = append(columns[:index], columns[index+1:]...)
}

func addItemToMap(m *map[int][]int, key, value int) {
	if _, ok := (*m)[key]; ok {

		valTmp := (*m)[key]
		valTmp = append(valTmp, value)

		slices.Sort(valTmp)

		(*m)[key] = valTmp

	} else {
		(*m)[key] = []int{value}
	}
}
