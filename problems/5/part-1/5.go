package problems

import (
	"slices"
	"strconv"
	"strings"

	common_functions "aoc.2023/lib/common/functions"
)

func solveChallenge() int {
	// Process the input
	scanner := common_functions.CreateInputScanner("../input.txt")
	defer scanner.File.Close()

	var seeds []int
	var gardenMaps [][]int

	for scanner.Scan() {
		line := strings.TrimSpace(scanner.Text())

		if len(line) > 0 {
			if strings.HasPrefix(line, "seeds:") {
				separatorIndex := strings.Index(line, ":")
				seedsStr := strings.Fields(line[separatorIndex+1:])

				seeds = getIntegersSlice(seedsStr)

				continue
			}

			if !strings.HasSuffix(line, "map:") {
				gardenMaps = append(gardenMaps, getIntegersSlice(strings.Fields(line)))
			} else {
				if gardenMaps != nil {
					computeGardenMap(&seeds, gardenMaps)
					gardenMaps = nil
				}
			}
		}
	}

	computeGardenMap(&seeds, gardenMaps)

	return slices.Min(seeds)
}

func computeGardenMap(seeds *[]int, gardenMaps [][]int) {
	entriesFounded := 0

	for i, seed := range *seeds {
		for _, gardenMap := range gardenMaps {
			limitTmp := (gardenMap[1] + gardenMap[2]) - 1

			if gardenMap[1] <= seed && limitTmp >= seed {
				diffTmp := seed - gardenMap[1]
				newValue := gardenMap[0] + diffTmp

				(*seeds)[i] = newValue

				entriesFounded += 1

				if entriesFounded == len(*seeds) {
					return
				}
			}
		}
	}
}

func getIntegersSlice(arr []string) []int {
	integers := make([]int, len(arr))

	for i, v := range arr {
		n, err := strconv.Atoi(v)

		if err != nil {
			panic(err)
		}

		integers[i] = n
	}

	return integers
}
