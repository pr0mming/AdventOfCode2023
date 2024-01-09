package common_functions

import (
	"bufio"
	"os"
	"path/filepath"

	common_types "aoc.2023/lib/common/types"
)

func CreateInputScanner(filePath string) *common_types.FileInputScanner {
	absPath, _ := filepath.Abs(filePath)
	file, err := os.Open(absPath)
	if err != nil {
		panic(err)
	}

	return &common_types.FileInputScanner{
		Scanner: bufio.NewScanner(file),
		File:    file,
	}
}
