package problems_8_2

import (
	"fmt"
	"regexp"
	"slices"
	"strconv"
	"strings"

	common_functions "aoc.2023/lib/common/functions"
)

var STARTING_POINT_SUFFIX = "A"
var ENDING_POINT_SUFFIX = "Z"

func SolveChallenge(problemId string) string {
	// Process the input
	inputFilePath := fmt.Sprintf("problems/%s/input.txt", problemId)
	scanner := common_functions.CreateInputScanner(inputFilePath)
	defer scanner.File.Close()

	var instructions string
	var networkMap = make(map[string][2]string)

	// Keep the instructions
	for scanner.Scan() {
		line := scanner.Text()

		if len(line) == 0 {
			break
		}

		instructions = line
	}

	// Use logic to map the string inputs in the map
	splitPattern := regexp.MustCompile(`(\w+)\s*=\s*\(([^)]+)\)`)
	var startingNodes []string

	for scanner.Scan() {
		line := scanner.Text()

		matches := splitPattern.FindStringSubmatch(line)

		point := matches[1]
		paths := strings.Split(matches[2], ",")

		pathValue := [2]string{strings.TrimSpace(paths[0]), strings.TrimSpace(paths[1])}

		// We use a map to keep the values as ["AAA"] = [POINT1, POINT 2] to ease accessibility
		networkMap[point] = pathValue

		// Now we need n nodes to analyze, this time those that start with "A"
		if strings.HasSuffix(point, "A") {
			startingNodes = append(startingNodes, point)
		}
	}

	// Keep the number of steps (cost of reaching the end) for each point
	// We can you a bit of numbers to guess the answer instead of a full work iteration per each entry
	var steps = make([]int, len(startingNodes))

	for i, node := range startingNodes {
		nodeSteps := computeStepByNode(node, instructions, networkMap)

		steps[i] = nodeSteps
	}

	// We need to sort and get the greatest value
	// So that we can try to use this value against the other to check if mod % 2 == 0 (for all entries)
	// All of this process while we add up the greatest value
	slices.Sort(steps)

	maxStep := steps[len(steps)-1]
	answer := 0

	for i := maxStep; ; i += maxStep {
		isMatch := true

		for j := 0; j < len(steps)-1; j++ {
			if i%steps[j] > 0 {
				isMatch = false
				break
			}
		}

		if isMatch {
			answer = i
			break
		}
	}

	return strconv.Itoa(answer)
}

func computeStepByNode(startNode string, instructions string, networkMap map[string][2]string) int {
	steps := 0
	currentKeyPath := startNode // AAA
	currentIndexIns := 0        // Used to control if we are at the last instruction

	for {
		// Avoid overflow
		if currentIndexIns >= len(instructions) {
			currentIndexIns = 0
		}

		// Extract the instruction to analyze (R or L)
		instructionTmp := string(instructions[currentIndexIns])
		pathKey := ""

		switch instructionTmp {
		case "L":
			pathKey = networkMap[currentKeyPath][0]
		case "R":
			pathKey = networkMap[currentKeyPath][1]
		default:
			panic("Invalid instruction")
		}

		steps++

		// If the current point is ZZZ
		if strings.HasSuffix(pathKey, ENDING_POINT_SUFFIX) {
			break
		}

		// Update the path to analyze in the next iteration
		currentKeyPath = pathKey
		currentIndexIns++
	}

	return steps
}
