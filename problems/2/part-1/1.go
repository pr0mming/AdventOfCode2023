package problems

import (
	"regexp"
	"strconv"
	"strings"

	common_functions "aoc.2023/lib/common/functions"
)

// This constant is used for keep the color and ammount per cubes to validate
var CONSTRAINTS_MAP = map[string]int{
	"red":   12,
	"green": 13,
	"blue":  14,
}

func SolveChallenge() string {
	// Process the input
	scanner := common_functions.CreateInputScanner("../input.txt")
	defer scanner.File.Close()

	answer := 0

	for scanner.Scan() {
		line := scanner.Text()

		answer += verifyMatch(line)
	}

	return strconv.Itoa(answer)
}

func verifyMatch(input string) int {
	// This regex will extract the groups: [ammount] [color] given each input
	bagPattern := regexp.MustCompile(`(\d+)\s+(\w+)`)

	// Get the sets per game
	groups := strings.Split(input, ";")

	// Iterate over each group and extract numbers per color
	for _, group := range groups {
		// Find all matches in the group
		matches := bagPattern.FindAllStringSubmatch(group, -1)

		// Iterate over matches and update the colorCount map
		for _, match := range matches {
			// Convert the matched number from string to int
			numOfCubes, err := strconv.Atoi(match[1])
			if err != nil {
				panic(err)
			}

			colorCubeConstraint := CONSTRAINTS_MAP[match[2]]

			// If the game doesn't follow the constraints so it's an invalid game (dismiss in the total sum)
			if numOfCubes > colorCubeConstraint {
				return 0
			}
		}
	}

	gameIdPattern := regexp.MustCompile(`Game\s*(\d+):`)

	// Find the game ID
	match := gameIdPattern.FindStringSubmatch(input)

	// Convert the matched number from string to int
	gameId, err := strconv.Atoi(match[1])
	if err != nil {
		panic(err)
	}

	return gameId
}
