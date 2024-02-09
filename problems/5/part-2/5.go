package problems_5_2

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
	seeds, gardenMaps := computeAlmanacInput(problemId)

	lastGardenMap := gardenMaps[len(gardenMaps)-1]
	var lastRangeTmp = lastGardenMap[0]

	for _, gardenMap := range lastGardenMap {
		itemLimit := lastRangeTmp[1] + lastRangeTmp[2]
		var transformations [][]int = [][]int{{lastRangeTmp[1], itemLimit}}

		for i := len(gardenMaps) - 2; i >= 0; i-- {
			var transformationsTmp [][]int

			for j := 0; j < len(gardenMaps[i]) && gardenMaps[i][j][0] <= itemLimit; j++ {
				//rangeLimit := (gardenMaps[i][j][0] + gardenMaps[i][j][2])

				for _, transformation := range transformations {

					if transformation[0] < gardenMaps[i][j][0] {
						diff := (gardenMaps[i][j][0] - transformation[0]) + 1
						transformation[0] += diff
						transformationsTmp = append(transformationsTmp, []int{gardenMaps[i][j][1], gardenMaps[i][j][1] + diff})
					}

					if transformation[0] >= gardenMaps[i][j][0] {
						rangeLimit := (gardenMaps[i][j][0] + gardenMaps[i][j][2])

						if rangeLimit >= transformation[1] {
							diff := (transformation[1] - transformation[0])
							transformation[0] += diff
							transformationsTmp = append(transformationsTmp, []int{gardenMaps[i][j][1], gardenMaps[i][j][1] + diff})
						} else {
							diff := gardenMaps[i][j][2]
							transformation[0] += diff
							transformationsTmp = append(transformationsTmp, []int{gardenMaps[i][j][1], gardenMaps[i][j][1] + diff})
						}
					}

				}
			}

			transformations = transformationsTmp
		}

		// Add final logic here!
		offset := 0

		for _, transformation := range transformations {
			for i := 1; i < len(seeds); i += 2 {
				if transformation[0] >= seeds[i-1] && seeds[i-1] <= transformation[i] {
					index := transformation[0] - seeds[i-1]
					fmt.Println(gardenMap[index])
				}
			}

			offset += (transformation[1] - transformation[0])
		}

	}

	for _, gardenMap := range gardenMaps {
		computeGardenMap(&seeds, gardenMap)
	}

	// Return the minimum transformation
	return strconv.Itoa(slices.Min(seeds))
}

func computeAlmanacInput(problemId string) ([]int, [][][]int) {
	// Process the input
	inputFilePath := fmt.Sprintf("problems/%s/input-2.txt", problemId)
	scanner := common_functions.CreateInputScanner(inputFilePath)
	defer scanner.File.Close()

	var seeds []int
	var gardenMapTmp [][]int
	var gardenMaps [][][]int

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

				// for i := 1; i < len(seeds); i += 2 {
				// 	seeds[i] = seeds[i-1] + seeds[i]
				// }

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
					sortGardenByOutputKey(gardenMapTmp, 0)
					gardenMaps = append(gardenMaps, gardenMapTmp)

					gardenMapTmp = nil
				}
			}
		}
	}

	// Make the process with the last one
	sortGardenByOutputKey(gardenMapTmp, 1)

	if gardenMapTmp[0][1] > 0 {
		gardenMapTmp = append(gardenMapTmp, []int{0, 0, gardenMapTmp[0][1]})
	}

	for i := 1; i < len(gardenMapTmp); i++ {
		diff := (gardenMapTmp[i-1][1] + gardenMapTmp[i-1][2])

		if diff < gardenMapTmp[i][1] {
			gardenMapTmp = append(gardenMapTmp, []int{diff, diff, gardenMapTmp[i][1] - diff})
		}
	}

	sortGardenByOutputKey(gardenMapTmp, 0)
	gardenMaps = append(gardenMaps, gardenMapTmp)

	return seeds, gardenMaps
}

func sortGardenByOutputKey(gardenMap [][]int, indexOrder int) {
	// Sort the items of the garden by the index 1, to take in account in Binary Search
	slices.SortFunc(gardenMap, func(a, b []int) int {
		return cmp.Compare(a[indexOrder], b[indexOrder])
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
