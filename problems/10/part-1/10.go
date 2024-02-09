package problems_10_1

import (
	"fmt"
	"strconv"
	"strings"

	common_functions "aoc.2023/lib/common/functions"
	common_types "aoc.2023/lib/common/types"
)

const ANIMAL_FLAG = "S"

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

	var answer = 0

	pipeNetwork, animalPos := processPipeNetworkInput(*scanner)

	var animalSteps = 0
	var animalPositions = [4]int{
		DIRECTION_RIGHT_ENUM,
		DIRECTION_BOTTOM_ENUM,
		DIRECTION_LEFT_ENUM,
		DIRECTION_UP_ENUM,
	}

	for _, animalDir := range animalPositions {

		var (
			pathResult      = animalPos
			directionResult = animalDir
			animalStepsTmp  = 0
		)

		for {
			pathResult, directionResult = computePath(pathResult, pipeNetwork, directionResult)

			if directionResult == DIRECTION_UNSET_ENUM {
				break
			}

			animalStepsTmp++
		}

		if animalStepsTmp > animalSteps {
			animalSteps = animalStepsTmp
		}
	}

	answer = animalSteps - 1

	return strconv.Itoa(answer)
}

func processPipeNetworkInput(scanner common_types.FileInputScanner) ([]string, [2]int) {
	var pipeNetwork []string
	var animalPos [2]int

	for i := 0; scanner.Scan(); i++ {
		line := scanner.Text()

		pipeNetwork = append(pipeNetwork, line)

		animalColPos := strings.Index(line, ANIMAL_FLAG)
		if animalColPos > -1 {
			animalPos = [2]int{i, animalColPos}
		}
	}

	return pipeNetwork, animalPos
}

func computePath(path [2]int, pipeNetwork []string, pipeDirection int) ([2]int, int) {

	switch pipeDirection {
	case DIRECTION_RIGHT_ENUM:
		path[1]++

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
