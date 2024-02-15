package problems_10_1

import (
	"fmt"
	"math"
	"strconv"
	"strings"

	common_functions "aoc.2023/lib/common/functions"
	common_types "aoc.2023/lib/common/types"
)

// Map all variables as Enums
var ANIMAL_FLAG = []byte("S")[0]

const (
	DIRECTION_RIGHT_ENUM  = 0
	DIRECTION_BOTTOM_ENUM = 1
	DIRECTION_LEFT_ENUM   = 2
	DIRECTION_UP_ENUM     = 3
	DIRECTION_UNSET_ENUM  = -1
)

var (
	VERTICAL_PIPE_ENUM   = []byte("|")[0]
	HORIZONTAL_PIPE_ENUM = []byte("-")[0]
	NORTH_EAST_PIPE_ENUM = []byte("L")[0]
	NORTH_WEST_PIPE_ENUM = []byte("J")[0]
	SOUTH_WEST_PIPE_ENUM = []byte("7")[0]
	SOUTH_EAST_PIPE_ENUM = []byte("F")[0]
	GROUND_PIPE_ENUM     = []byte(".")[0]
)

func SolveChallenge(problemId string) string {
	// Process the input
	inputFilePath := fmt.Sprintf("problems/%s/input.txt", problemId)
	scanner := common_functions.CreateInputScanner(inputFilePath)
	defer scanner.File.Close()

	var answer float64 = 0

	// Get the map and position of the animal to start trip
	pipeNetwork, animalPos := processPipeNetworkInput(*scanner)

	// At the beginning the animal has only 4 directions to travel...
	var (
		animalSteps     float64 = 0
		animalPositions         = [4]int{
			DIRECTION_RIGHT_ENUM,
			DIRECTION_BOTTOM_ENUM,
			DIRECTION_LEFT_ENUM,
			DIRECTION_UP_ENUM,
		}
	)

	for _, animalDir := range animalPositions {

		var (
			pathResult              = animalPos
			directionResult         = animalDir
			animalStepsTmp  float64 = 0 // Use it to save the greatest value
		)

		// Infinite loop that breaks when we can't travel in the pipe network
		for {
			// Recalculate the new step
			pathResult, directionResult = computePath(pathResult, pipeNetwork, directionResult)

			// If we reached a limit in the map (out of boundaries limit)
			// Or there isn't a valid connected pipe
			// Or there is a ground
			// Or if there is the animal again! (infinite loop)
			if directionResult == DIRECTION_UNSET_ENUM {

				// We make sure if it's an infinite path,
				// then is necessary divide by 2 to get the greatest path
				posTmp := pipeNetwork[pathResult[0]][pathResult[1]]
				if posTmp == ANIMAL_FLAG {
					animalStepsTmp = math.Floor((animalStepsTmp + 1) / 2)
				}

				break
			}

			animalStepsTmp++
		}

		if animalStepsTmp > animalSteps {
			animalSteps = animalStepsTmp
		}
	}

	answer = animalSteps

	return strconv.FormatFloat(answer, 'f', -1, 64)
}

func processPipeNetworkInput(scanner common_types.FileInputScanner) ([]string, [2]int) {
	var (
		pipeNetwork   []string
		animalPos     [2]int
		animalFlagStr = string(ANIMAL_FLAG)
	)

	for i := 0; scanner.Scan(); i++ {
		line := scanner.Text()

		pipeNetwork = append(pipeNetwork, line)

		// Detect animal position
		animalColPos := strings.Index(line, animalFlagStr)
		if animalColPos > -1 {
			animalPos = [2]int{i, animalColPos}
		}
	}

	return pipeNetwork, animalPos
}

func computePath(path [2]int, pipeNetwork []string, pipeDirection int) ([2]int, int) {
	// Direction is to know how to move through the map (columns and rows)
	switch pipeDirection {
	case DIRECTION_RIGHT_ENUM:
		path[1]++

		// Out of boundaries
		if path[1] >= len(pipeNetwork[1]) {
			return path, DIRECTION_UNSET_ENUM
		}

		newDirection := validateRightDirection(pipeNetwork[path[0]][path[1]])
		return path, newDirection
	case DIRECTION_BOTTOM_ENUM:
		path[0]++

		if path[0] >= len(pipeNetwork) {
			return path, DIRECTION_UNSET_ENUM
		}

		newDirection := validateBottomDirection(pipeNetwork[path[0]][path[1]])
		return path, newDirection
	case DIRECTION_LEFT_ENUM:
		path[1]--

		if path[1] < 0 {
			return path, DIRECTION_UNSET_ENUM
		}

		newDirection := validateLeftDirection(pipeNetwork[path[0]][path[1]])
		return path, newDirection
	case DIRECTION_UP_ENUM:
		path[0]--

		if path[0] < 0 {
			return path, DIRECTION_UNSET_ENUM
		}

		newDirection := validateUpDirection(pipeNetwork[path[0]][path[1]])
		return path, newDirection
	default:
		panic("Bad pipe direction")
	}
}

// These 4 methods are to calculate the new direction according to the pipe
func validateRightDirection(pipe byte) int {
	switch pipe {
	case HORIZONTAL_PIPE_ENUM:
		return DIRECTION_RIGHT_ENUM

	case NORTH_WEST_PIPE_ENUM:
		return DIRECTION_UP_ENUM

	case SOUTH_WEST_PIPE_ENUM:
		return DIRECTION_BOTTOM_ENUM

	default:
		return DIRECTION_UNSET_ENUM
	}
}

func validateBottomDirection(pipe byte) int {
	switch pipe {
	case VERTICAL_PIPE_ENUM:
		return DIRECTION_BOTTOM_ENUM

	case NORTH_EAST_PIPE_ENUM:
		return DIRECTION_RIGHT_ENUM

	case NORTH_WEST_PIPE_ENUM:
		return DIRECTION_LEFT_ENUM

	default:
		return DIRECTION_UNSET_ENUM
	}
}

func validateLeftDirection(pipe byte) int {
	switch pipe {
	case HORIZONTAL_PIPE_ENUM:
		return DIRECTION_LEFT_ENUM

	case NORTH_EAST_PIPE_ENUM:
		return DIRECTION_UP_ENUM

	case SOUTH_EAST_PIPE_ENUM:
		return DIRECTION_BOTTOM_ENUM

	default:
		return DIRECTION_UNSET_ENUM
	}
}

func validateUpDirection(pipe byte) int {
	switch pipe {
	case VERTICAL_PIPE_ENUM:
		return DIRECTION_UP_ENUM

	case SOUTH_EAST_PIPE_ENUM:
		return DIRECTION_RIGHT_ENUM

	case SOUTH_WEST_PIPE_ENUM:
		return DIRECTION_LEFT_ENUM

	default:
		return DIRECTION_UNSET_ENUM
	}
}
