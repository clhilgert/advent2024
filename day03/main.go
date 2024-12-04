package main

import (
	"fmt"
	"io"
	"log"
	"regexp"
	"strconv"

	"github.com/clhilgert/advent2024/pkg/utils"
)

func main() {
	file := utils.GetInputFile()
	data, err := io.ReadAll(file)
	if err != nil {
		log.Fatalf("Error reading file %v\n", err)
	}

	input := string(data)

	mulRe := regexp.MustCompile(`mul\((\d{1,3}),(\d{1,3})\)`)
	doRe := regexp.MustCompile(`do\(\)`)
	dontRe := regexp.MustCompile(`don't\(\)`)

	// Part 1
	part1Sum := 0
	part1Matches := mulRe.FindAllStringSubmatch(input, -1)
	for _, match := range part1Matches {
		x, _ := strconv.Atoi(match[1])
		y, _ := strconv.Atoi(match[2])
		part1Sum += x * y
	}

	// Part 2
	mulEnabled := true
	part2Sum := 0

	for len(input) > 0 {
		mulMatch := mulRe.FindStringSubmatchIndex(input)
		doMatch := doRe.FindStringIndex(input)
		dontMatch := dontRe.FindStringIndex(input)

		nextMatch := -1
		nextType := ""
		if mulMatch != nil {
			nextMatch = mulMatch[0]
			nextType = "mul"
		}
		if doMatch != nil && (nextMatch == -1 || doMatch[0] < nextMatch) {
			nextMatch = doMatch[0]
			nextType = "do"
		}
		if dontMatch != nil && (nextMatch == -1 || dontMatch[0] < nextMatch) {
			nextMatch = dontMatch[0]
			nextType = "don't"
		}

		if nextMatch == -1 {
			break
		}

		switch nextType {
		case "mul":
			if mulEnabled {
				x, _ := strconv.Atoi(input[mulMatch[2]:mulMatch[3]])
				y, _ := strconv.Atoi(input[mulMatch[4]:mulMatch[5]])
				part2Sum += x * y
			}
			input = input[mulMatch[1]:]
		case "do":
			mulEnabled = true
			input = input[doMatch[1]:]
		case "don't":
			mulEnabled = false
			input = input[dontMatch[1]:]
		}
	}

	fmt.Printf("Part 1: %d\n", part1Sum)
	fmt.Printf("Part 2: %d\n", part2Sum)
}
