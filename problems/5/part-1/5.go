package problems_5_1

import (
	"fmt"
	"slices"
	"strconv"
	"strings"

	common_functions "aoc.2023/lib/common/functions"
)

func SolveChallenge(problemId string) string {
	// Process the input
	inputFilePath := fmt.Sprintf("problems/%s/input.txt", problemId)
	scanner := common_functions.CreateInputScanner(inputFilePath)
	defer scanner.File.Close()

	var seeds []int
	var gardenMaps [][]int

	for scanner.Scan() {
		line := strings.TrimSpace(scanner.Text())

		// This validation is to skip breaklines
		if len(line) > 0 {
			// We save the seeds
			if strings.HasPrefix(line, "seeds:") {
				// Extra work to get the numbers
				separatorIndex := strings.Index(line, ":")
				seedsStr := strings.Fields(line[separatorIndex+1:])

				seeds = common_functions.GetIntegersArr(seedsStr, false)

				continue
			}

			// This is to save the map of resources of the garden
			if !strings.HasSuffix(line, "map:") {
				gardenMaps = append(gardenMaps, common_functions.GetIntegersArr(strings.Fields(line), false))
			} else {
				if gardenMaps != nil {
					// Make the conversion
					computeGardenMap(&seeds, gardenMaps)
					gardenMaps = nil
				}
			}
		}
	}

	computeGardenMap(&seeds, gardenMaps)

	return strconv.Itoa(slices.Min(seeds))
}

func computeGardenMap(seeds *[]int, gardenMaps [][]int) {
	entriesFounded := 0

	for i, seed := range *seeds {
		for _, gardenMap := range gardenMaps {
			// Extact the limit for this map
			limitTmp := (gardenMap[1] + gardenMap[2]) - 1

			// If the seed is in the range
			if gardenMap[1] <= seed && limitTmp >= seed {
				// Calculate the number of the range
				diffTmp := seed - gardenMap[1]
				newValue := gardenMap[0] + diffTmp

				(*seeds)[i] = newValue

				entriesFounded += 1

				// Avoid extra iterations
				if entriesFounded == len(*seeds) {
					return
				}
			}
		}
	}
}
