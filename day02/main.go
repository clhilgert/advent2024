package main

import (
	"bufio"
	"fmt"

	"github.com/clhilgert/advent2024/pkg/utils"
)

func main() {
	file := utils.GetInputFile()
	defer file.Close()

	scanner := bufio.NewScanner(file)
	count := 0

	for scanner.Scan() {
		nums, err := utils.ParseLineToInts(scanner.Text())
		if err != nil {
			fmt.Println("Error parsing line:", err)
			return
		}

		// Part 1
		if isSafe(nums) {
			count++
		} else {
			// Part 2
			for i := range nums {
				newNums := append(append([]int{}, nums[:i]...), nums[i+1:]...)
				if isSafe(newNums) {
					count++
					break
				}
			}
		}
	}

	if err := scanner.Err(); err != nil {
		fmt.Println("Error reading input:", err)
	}

	fmt.Println(count)
}

func isSafe(nums []int) bool {
	if len(nums) < 2 {
		return false
	}

	prev := nums[0]
	pos := false
	neg := false
	safe := true

	if prev < nums[1] {
		pos = true
	} else if prev > nums[1] {
		neg = true
	} else {
		safe = false
	}

	for i := 1; i < len(nums); i++ {
		if nums[i] == prev {
			safe = false
		} else if utils.AbsInt(nums[i]-prev) > 3 {
			safe = false
		} else if (prev > nums[i]) && pos {
			safe = false
		} else if (prev < nums[i]) && neg {
			safe = false
		}
		prev = nums[i]
	}

	return safe
}
