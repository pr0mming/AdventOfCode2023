package common_types

import (
	"bufio"
	"os"
)

type FileInputScanner struct {
	*bufio.Scanner
	File *os.File
}
