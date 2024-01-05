package main

import (
	"bufio"
	"fmt"
	"os"
	"path/filepath"
	"regexp"
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
	var answer = 0

	for scanner.Scan() {
		line := scanner.Text()

		answer += getCubesProd(line)
	}

	return answer
}

func getCubesProd(input string) int {
	cubesDict := make(map[string]int)

	// Define a regular expression bagPattern to match numbers and colors
	bagPattern := regexp.MustCompile(`(\d+)\s+(\w+)`)

	// Split the input into individual groups
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
