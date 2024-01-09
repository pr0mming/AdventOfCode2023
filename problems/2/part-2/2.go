package problems

import (
	"regexp"
	"strconv"
	"strings"

	common_functions "aoc.2023/lib/common/functions"
)

func solveChallenge() int {
	// Process the input
	scanner := common_functions.CreateInputScanner("../input.txt")
	defer scanner.File.Close()

	var answer = 0

	for scanner.Scan() {
		line := scanner.Text()

		answer += getCubesProd(line)
	}

	return answer
}

func getCubesProd(input string) int {
	cubesDict := make(map[string]int)

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
				panic("Error converting string to int:")
			}

			cubeColorTmp := match[2]
			cubeAmmount, ok := cubesDict[cubeColorTmp]

			// If the item doesn't exist we put it in the map, otherwise we compare it if is the new highest
			if !ok || (ok && numOfCubes > cubeAmmount) {
				cubesDict[cubeColorTmp] = numOfCubes
			}
		}
	}

	var cubeProd = 1
	for _, value := range cubesDict {
		cubeProd *= value
	}

	return cubeProd
}
