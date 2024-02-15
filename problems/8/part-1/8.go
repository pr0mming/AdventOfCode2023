package problems_8_1

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"

	common_functions "aoc.2023/lib/common/functions"
	common_types "aoc.2023/lib/common/types"
)

const STARTING_POINT_REF = "AAA"
const ENDING_POINT_REF = "ZZZ"

func SolveChallenge(problemId string) string {
	// Process the input
	inputFilePath := fmt.Sprintf("problems/%s/input.txt", problemId)
	scanner := common_functions.CreateInputScanner(inputFilePath)
	defer scanner.File.Close()

	var (
		instructions string
		networkMap   = make(map[string][2]string)
	)

	// Keep the instructions
	for scanner.Scan() {
		line := scanner.Text()

		if len(line) == 0 {
			break
		}

		instructions = line
	}

	// Use logic to map the string inputs in the map
	networkMap = getNetworkInput(*scanner)

	var (
		answer          int    = 0
		currentKeyPath  string = STARTING_POINT_REF // AAA
		currentIndexIns int    = 0                  // Used to control if we are at the last instruction
	)

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

		answer++

		// If the current point is ZZZ
		if pathKey == ENDING_POINT_REF {
			break
		}

		// Update the path to analyze in the next iteration
		currentKeyPath = pathKey
		currentIndexIns++
	}

	return strconv.Itoa(answer)
}

func getNetworkInput(scanner common_types.FileInputScanner) map[string][2]string {
	var (
		networkMap   = make(map[string][2]string)
		splitPattern = regexp.MustCompile(`(\w+)\s*=\s*\(([^)]+)\)`)
	)

	for scanner.Scan() {
		line := scanner.Text()

		matches := splitPattern.FindStringSubmatch(line)

		point := matches[1]
		paths := strings.Split(matches[2], ",")

		pathValue := [2]string{strings.TrimSpace(paths[0]), strings.TrimSpace(paths[1])}

		// We use a map to keep the values as ["AAA"] = [POINT1, POINT 2] to ease accessibility
		networkMap[point] = pathValue
	}

	return networkMap
}
