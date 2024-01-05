package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"path/filepath"
	"sort"
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

		firstIndexSeparator := strings.Index(line, ":")
		secondIndexSeparator := strings.Index(line, "|")

		winningNumbers := strings.Fields(line[firstIndexSeparator+1 : secondIndexSeparator])
		myNumbers := strings.Fields(line[secondIndexSeparator+1:])

		answer += computeScore(winningNumbers, myNumbers)
	}

	return answer
}

func computeScore(winningNumbers []string, myNumbers []string) int {
	integersForWN := getIntegersSlice(winningNumbers)
	integersForMN := getIntegersSlice(myNumbers)
	n := 0

	for _, v := range integersForMN {
		i := sort.Search(len(integersForWN), func(i int) bool {
			return integersForWN[i] >= v
		})

		if i < len(integersForWN) && integersForWN[i] == v {
			n += 1
		}
	}

	if n > 0 {
		return int(math.Pow(2, float64(n-1)))
	}

	return 0
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

	sort.Ints(integers)

	return integers
}
