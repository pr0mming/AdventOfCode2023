package main

import (
	"bufio"
	"fmt"
	"os"
	"path/filepath"
	"strconv"
)

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

func solveChallenge(scanner *fileScanner) int {
	answer := 0

	for scanner.Scan() {
		line := scanner.Text()
		var numberForLine string

		for i := 0; i < len(line); i++ {
			charTmp := line[i]

			if charTmp >= 48 && charTmp <= 57 {
				numberForLine += string(charTmp)
				break
			}
		}

		for i := len(line) - 1; i >= 0; i-- {
			charTmp := line[i]

			if charTmp >= 48 && charTmp <= 57 {
				numberForLine += string(charTmp)
				break
			}
		}

		intTmp, err := strconv.Atoi(numberForLine)
		if err != nil {
			panic(err)
		}

		answer += intTmp
	}

	return answer
}
