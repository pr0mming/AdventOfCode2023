package main

import (
	"bufio"
	"fmt"
	"os"
	"path/filepath"
	"slices"
	"strconv"
	"strings"
)

type fileScanner struct {
	*bufio.Scanner
	file *os.File
}

func createScanner(filePath string) (*fileScanner, error) {
	absPath, _ := filepath.Abs(filePath)
	file, err := os.Open(absPath)
	if err != nil {
		return nil, err
	}

	return &fileScanner{
		Scanner: bufio.NewScanner(file),
		file:    file,
	}, nil
}

func main() {
	filePath := "../input.txt"

	scanner, err := createScanner(filePath)
	if err != nil {
		fmt.Println("Error opening the file:", err)
		return
	}
	defer scanner.file.Close()

	answer := solveChallenge(scanner)
	fmt.Println(answer)
}

func solveChallenge(scanner *fileScanner) int {
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
