package problems_15_2

import (
	"fmt"
	"slices"
	"strconv"
	"strings"

	common_functions "aoc.2023/lib/common/functions"
	types "aoc.2023/lib/common/types/problems/15"
)

const (
	ADD_STEP_ENUM    = 0
	REMOVE_STEP_ENUM = 1
)

func SolveChallenge(problemId string) string {
	// Process the input
	inputFilePath := fmt.Sprintf("problems/%s/input.txt", problemId)
	scanner := common_functions.CreateInputScanner(inputFilePath)
	defer scanner.File.Close()

	var (
		answer int = 0
		steps  []string
	)

	// Because it's just one line
	if scanner.Scan() {
		line := scanner.Text()

		// Get inputs
		steps = strings.Split(line, ",")
	}

	answer = getHashMapOperation(steps)

	return strconv.Itoa(answer)
}

func getHASHChar(input string) int {
	output := 0

	// The same steps of the problem:
	for _, char := range input {
		output += int(char) // Increase the current value by the ASCII code you just determined.

		output *= 17  // Set the current value to itself multiplied by 17.
		output %= 256 // Set the current value to the remainder of dividing itself by 256.
	}

	return output
}

func getHashMapOperation(steps []string) int {
	var (
		output int                   = 0
		boxes  [256][]types.BoxValue // 256 because is the max number of boxes possible
	)

	for _, v := range steps {

		stepParam := getStepParameters(v)

		switch stepParam.Operation {
		case ADD_STEP_ENUM:

			addStep(&boxes, stepParam)

		case REMOVE_STEP_ENUM:

			removeStep(&boxes, stepParam)

		}
	}

	for i, v := range boxes {

		// Filter boxes with values
		if len(v) > 0 {
			slot := 1
			for _, item := range v {
				output += (i + 1) * (slot) * (item.FocalLength)
				slot++
			}
		}

	}

	return output
}

func getStepParameters(input string) types.BoxValue {
	var (
		output types.BoxValue
	)

	const (
		addChar    = "="
		removeChar = "-"
	)

	// For "rn=1" we get "rn", "=" and "1"
	// For "cm-" we get "cm", "-"

	if strings.Contains(input, addChar) {

		index := strings.Index(input, addChar)
		leftValue := input[:index]
		rightValue := input[index+1:]

		output.Step = leftValue
		output.FocalLength = common_functions.Atoi(rightValue)
		output.Operation = ADD_STEP_ENUM

	} else if strings.Contains(input, removeChar) {

		index := strings.Index(input, removeChar)
		leftValue := input[:index]

		output.Step = leftValue
		output.Operation = REMOVE_STEP_ENUM

	}

	return output
}

func addStep(boxes *[256][]types.BoxValue, parameters types.BoxValue) {
	step := parameters.Step
	boxNumber := getHASHChar(step) // Calculate wich box belongs

	values := boxes[boxNumber]
	focalLength := parameters.FocalLength

	indexItem := slices.IndexFunc(values, func(a types.BoxValue) bool {
		return a.Step == step
	})

	// If the step is new then we add at the end of the array
	if indexItem == -1 {

		values = append(values, types.BoxValue{
			Step:        step,
			FocalLength: focalLength,
		})

	} else {
		// Othwerwise we update its value
		values[indexItem].FocalLength = focalLength
	}

	boxes[boxNumber] = values
}

func removeStep(boxes *[256][]types.BoxValue, parameters types.BoxValue) {
	step := parameters.Step
	boxNumber := getHASHChar(step) // Calculate wich box belongs

	if len(boxes[boxNumber]) > 0 {
		values := boxes[boxNumber]
		indexItem := slices.IndexFunc(values, func(a types.BoxValue) bool {
			return a.Step == step
		})

		// If the record exists we delete it by index
		if indexItem > -1 {
			values = append(values[:indexItem], values[indexItem+1:]...)

			boxes[boxNumber] = values
		}
	}
}
