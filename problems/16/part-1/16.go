package problems_16_1

import (
	"fmt"
	"strconv"

	common_functions "aoc.2023/lib/common/functions"
)

const (
	BEAM_UP_DIRECTION       = 0
	BEAM_RIGHT_DIRECTION    = 1
	BEAM_DOWN_DIRECTION     = 2
	BEAM_LEFT_DIRECTION     = 3
	BEAM_SPLITTED_DIRECTION = 4

	ERROR_BAD_TILE_MSG = "this is invalid tile: %s"
)

var (
	EMPTY_SCAPE_TILE         = []byte(".")[0]
	LEAN_RIGHT_MIRROR_TILE   = []byte("/")[0]
	LEAN_LEFT_MIRROR_TILE    = []byte("\\")[0]
	HORIZONTAL_SPLITTER_TILE = []byte("-")[0]
	VERTICAL_SPLITTER_TILE   = []byte("|")[0]
)

var (
	// It's like a hashmap, to save map["0,0"] = map[0], map[1], etc..
	// map[0], map[1], etc... are the possible directions (above)
	ENERGIZED_TILES_MAP = make(map[string]map[int]struct{})
	PUZZLE_MATRIX       []string // Matrix to save the input
)

func SolveChallenge(problemId string) string {
	// Process the input
	inputFilePath := fmt.Sprintf("problems/%s/input.txt", problemId)
	scanner := common_functions.CreateInputScanner(inputFilePath)
	defer scanner.File.Close()

	answer := 0

	for scanner.Scan() {
		line := scanner.Text()

		PUZZLE_MATRIX = append(PUZZLE_MATRIX, line)
	}

	// The beam starts from [0, 0] position in the RIGHT direction
	completeBeamTrip(0, 0, BEAM_RIGHT_DIRECTION)

	answer = len(ENERGIZED_TILES_MAP)

	return strconv.Itoa(answer)
}

func completeBeamTrip(row, column, direction int) {
	// The loop is active while the [X, Y] positions are valid for the matrix
	for row >= 0 && row < len((PUZZLE_MATRIX)) && column >= 0 && column < len((PUZZLE_MATRIX)[0]) {

		// Is possible that the point [X, Y] to analyze is already analyzed...
		// If the point exists in the map isn't enough...
		// We need to check the directions analyzed in the past, if we have one in it so this will trigger an infinite recursion!
		// So we break it!
		mapKey := fmt.Sprintf("%d,%d", row, column)
		if _, ok := ENERGIZED_TILES_MAP[mapKey]; ok {
			directions := ENERGIZED_TILES_MAP[mapKey]

			if _, ok := directions[direction]; ok {
				return
			}

		} else {
			ENERGIZED_TILES_MAP[mapKey] = make(map[int]struct{})
		}

		// Otherwise we analyze the point
		// The false case should be the BEAM_SPLITTED_DIRECTION case because is a final movement
		if !computeBeamDirection(&direction, &row, &column) {
			return
		}
	}
}

func computeBeamDirection(direction, row, column *int) bool {
	tile := PUZZLE_MATRIX[(*row)][(*column)]

	// Save energyzed tile
	// Also save direction to check when is necessary break the recursion loop
	mapKey := fmt.Sprintf("%d,%d", (*row), (*column))
	directions := ENERGIZED_TILES_MAP[mapKey]

	directions[(*direction)] = struct{}{}
	ENERGIZED_TILES_MAP[mapKey] = directions

	switch *direction {
	case BEAM_UP_DIRECTION:

		computeBeamUpDirection(tile, direction, row, column)

	case BEAM_RIGHT_DIRECTION:

		computeBeamRightDirection(tile, direction, row, column)

	case BEAM_DOWN_DIRECTION:

		computeBeamDownDirection(tile, direction, row, column)

	case BEAM_LEFT_DIRECTION:

		computeBeamLeftDirection(tile, direction, row, column)

	default:
		return false //when is BEAM_SPLITTED_DIRECTION
	}

	return true
}

// 4 similar methods for 4 possible directions

func computeBeamUpDirection(tile byte, direction, row, column *int) {
	// Calculate new direction of the beam
	newDirection, newRow, newColumn := getNextBeamDirectionFromUp(tile, (*row), (*column))

	// If we have to take two paths we use recursion call with both paths
	if newDirection == BEAM_SPLITTED_DIRECTION {
		completeBeamTrip((*row), (*column)-1, BEAM_LEFT_DIRECTION)
		completeBeamTrip((*row), (*column)+1, BEAM_RIGHT_DIRECTION)
	}

	// Save to the next iteration
	(*direction) = newDirection
	(*row) = newRow
	(*column) = newColumn
}

func computeBeamRightDirection(tile byte, direction, row, column *int) {
	newDirection, newRow, newColumn := getNextBeamDirectionFromRight(tile, (*row), (*column))

	if newDirection == BEAM_SPLITTED_DIRECTION {
		completeBeamTrip((*row)-1, (*column), BEAM_UP_DIRECTION)
		completeBeamTrip((*row)+1, (*column), BEAM_DOWN_DIRECTION)
	}

	(*direction) = newDirection
	(*row) = newRow
	(*column) = newColumn
}

func computeBeamDownDirection(tile byte, direction, row, column *int) {
	newDirection, newRow, newColumn := getNextBeamDirectionFromDown(tile, (*row), (*column))

	if newDirection == BEAM_SPLITTED_DIRECTION {
		completeBeamTrip((*row), (*column)-1, BEAM_LEFT_DIRECTION)
		completeBeamTrip((*row), (*column)+1, BEAM_RIGHT_DIRECTION)
	}

	(*direction) = newDirection
	(*row) = newRow
	(*column) = newColumn
}

func computeBeamLeftDirection(tile byte, direction, row, column *int) {
	newDirection, newRow, newColumn := getNextBeamDirectionFromLeft(tile, (*row), (*column))

	if newDirection == BEAM_SPLITTED_DIRECTION {
		completeBeamTrip((*row)-1, (*column), BEAM_UP_DIRECTION)
		completeBeamTrip((*row)+1, (*column), BEAM_DOWN_DIRECTION)
	}

	(*direction) = newDirection
	(*row) = newRow
	(*column) = newColumn
}

// 4 similar methods for 4 possible directions

func getNextBeamDirectionFromUp(tile byte, row, column int) (int, int, int) {
	// Depending on the tile char we need to return the new coords
	switch tile {
	case EMPTY_SCAPE_TILE, VERTICAL_SPLITTER_TILE:
		row--
		return BEAM_UP_DIRECTION, row, column

	case HORIZONTAL_SPLITTER_TILE:
		return BEAM_SPLITTED_DIRECTION, row, column

	case LEAN_LEFT_MIRROR_TILE:
		column--
		return BEAM_LEFT_DIRECTION, row, column

	case LEAN_RIGHT_MIRROR_TILE:
		column++
		return BEAM_RIGHT_DIRECTION, row, column
	}

	panic(fmt.Sprintf(ERROR_BAD_TILE_MSG, string(tile)))
}

func getNextBeamDirectionFromRight(tile byte, row, column int) (int, int, int) {
	switch tile {
	case EMPTY_SCAPE_TILE, HORIZONTAL_SPLITTER_TILE:
		column++
		return BEAM_RIGHT_DIRECTION, row, column

	case VERTICAL_SPLITTER_TILE:
		return BEAM_SPLITTED_DIRECTION, row, column

	case LEAN_LEFT_MIRROR_TILE:
		row++
		return BEAM_DOWN_DIRECTION, row, column

	case LEAN_RIGHT_MIRROR_TILE:
		row--
		return BEAM_UP_DIRECTION, row, column
	}

	panic(fmt.Sprintf(ERROR_BAD_TILE_MSG, string(tile)))
}

func getNextBeamDirectionFromDown(tile byte, row, column int) (int, int, int) {
	switch tile {
	case EMPTY_SCAPE_TILE, VERTICAL_SPLITTER_TILE:
		row++
		return BEAM_DOWN_DIRECTION, row, column

	case HORIZONTAL_SPLITTER_TILE:
		return BEAM_SPLITTED_DIRECTION, row, column

	case LEAN_LEFT_MIRROR_TILE:
		column++
		return BEAM_RIGHT_DIRECTION, row, column

	case LEAN_RIGHT_MIRROR_TILE:
		column--
		return BEAM_LEFT_DIRECTION, row, column
	}

	panic(fmt.Sprintf(ERROR_BAD_TILE_MSG, string(tile)))
}

func getNextBeamDirectionFromLeft(tile byte, row, column int) (int, int, int) {
	switch tile {
	case EMPTY_SCAPE_TILE, HORIZONTAL_SPLITTER_TILE:
		column--
		return BEAM_LEFT_DIRECTION, row, column

	case VERTICAL_SPLITTER_TILE:
		return BEAM_SPLITTED_DIRECTION, row, column

	case LEAN_LEFT_MIRROR_TILE:
		row--
		return BEAM_UP_DIRECTION, row, column

	case LEAN_RIGHT_MIRROR_TILE:
		row++
		return BEAM_DOWN_DIRECTION, row, column
	}

	panic(fmt.Sprintf(ERROR_BAD_TILE_MSG, string(tile)))
}
