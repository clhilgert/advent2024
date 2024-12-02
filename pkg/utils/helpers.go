package utils

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"runtime"
	"strconv"
	"strings"
)

func GetInputFile() *os.File {
	var callerFile string
	var ok bool

	// depth -- 1: main -> GetInputFile() or 2: main -> SliceFromInput() -> GetInputFile()
	for i := 1; i <= 2; i++ {
		_, callerFile, _, ok = runtime.Caller(i)
		if !ok {
			log.Fatal("Unable to determine caller information")
		}
		if strings.Contains(callerFile, "main.go") {
			break
		}
	}

	callerDir := filepath.Dir(callerFile)
	inputPath := filepath.Join(callerDir, "input.txt")

	file, err := os.Open(inputPath)
	if err != nil {
		log.Fatal(err)
	}

	return file
}

func SliceFromInput() []string {
	file := GetInputFile()
	var lines []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		row := scanner.Text()
		lines = append(lines, row)
	}
	return lines
}

func AbsInt(n int) int {
	if n < 0 {
		return -n
	}
	return n
}

func ParseLineToInts(line string) ([]int, error) {
	strNums := strings.Fields(line)
	nums := make([]int, len(strNums))

	for i, str := range strNums {
		num, err := strconv.Atoi(str)
		if err != nil {
			return nil, fmt.Errorf("error converting '%s' to int: %w", str, err)
		}
		nums[i] = num
	}

	return nums, nil
}
