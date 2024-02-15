package problems_5_1

import (
	"cmp"
	"fmt"
	"slices"
	"sort"
	"strconv"
	"strings"

	common_functions "aoc.2023/lib/common/functions"
)

func SolveChallenge(problemId string) string {
	// Process the input
	inputFilePath := fmt.Sprintf("problems/%s/input.txt", problemId)
	scanner := common_functions.CreateInputScanner(inputFilePath)
	defer scanner.File.Close()

	var (
		seeds        []int
		gardenMapTmp [][]int
		gardenMaps   [][][]int
	)

	for scanner.Scan() {
		line := strings.TrimSpace(scanner.Text())

		// This validation is to skip breaklines
		if len(line) > 0 {
			// We save the seeds
			if strings.HasPrefix(line, "seeds:") {
				// Extra work to get the numbers
				separatorIndex := strings.Index(line, ":")
				seedsStr := strings.Fields(line[separatorIndex+1:])

				// Convert to integers, (isn't necessary sort)
				seeds = common_functions.GetIntegersArr(seedsStr, false)

				continue
			}

			// This is to save the map of resources of the garden
			if !strings.HasSuffix(line, "map:") {
				// Make conversion
				valuesTmp := strings.Fields(line)
				gardenMapTmp = append(gardenMapTmp, common_functions.GetIntegersArr(valuesTmp, false))
			} else {
				if gardenMapTmp != nil {
					// Sort by index 1 (input)
					sortGardenByOutputKey(gardenMapTmp)
					gardenMaps = append(gardenMaps, gardenMapTmp)

					gardenMapTmp = nil
				}
			}
		}
	}

	// Make the process with the last one
	sortGardenByOutputKey(gardenMapTmp)
	gardenMaps = append(gardenMaps, gardenMapTmp)

	for _, gardenMap := range gardenMaps {
		computeGardenMap(&seeds, gardenMap)
	}

	// Return the minimum transformation
	return strconv.Itoa(slices.Min(seeds))
}

func sortGardenByOutputKey(gardenMap [][]int) {
	// Sort the items of the garden by the index 1, to take in account in Binary Search
	slices.SortFunc(gardenMap, func(a, b []int) int {
		return cmp.Compare(a[1], b[1])
	})
}

func computeGardenMap(seeds *[]int, targetGardenMap [][]int) {
	for i, input := range *seeds {
		// Check if the seed is in the range using Binary Search
		iFounded := sort.Search(len(targetGardenMap), func(i int) bool {
			mapTmp := targetGardenMap[i]

			limitOutputTmp := (mapTmp[1] + mapTmp[2]) - 1

			return mapTmp[1] >= input || limitOutputTmp >= input
		})

		// If the element is in there ...
		if iFounded < len(targetGardenMap) {
			mapTmp := targetGardenMap[iFounded]

			limitOutputTmp := (mapTmp[1] + mapTmp[2]) - 1

			// Update the same value seed with the equivalence
			if mapTmp[1] <= input && limitOutputTmp >= input {
				(*seeds)[i] = mapTmp[0] + (input - mapTmp[1])
			}
		}

		// Otherwise we keep the same value
	}
}
