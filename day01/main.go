package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"

	"github.com/clhilgert/advent2024/pkg/utils"
)

func main() {
	file := utils.GetInputFile()
	nums1, nums2 := parseAndSortNums(file)
	fmt.Printf("Part One: %d\n", partOne(nums1, nums2))
	fmt.Printf("Part Two: %d\n", partTwo(nums1, nums2))
}

func partOne(nums1, nums2 []int) int {
	var distance int
	for i := 0; i < len(nums1); i++ {
		distance += absInt(nums2[i] - nums1[i])
	}
	return distance
}

func partTwo(nums1, nums2 []int) int {
	var score int
	freq := make(map[int]int)
	for _, num := range nums2 {
		freq[num]++
	}
	for _, num := range nums1 {
		if freq, exists := freq[num]; exists {
			score += num * freq
		}
	}
	return score
}

func parseAndSortNums(file *os.File) ([]int, []int) {
	var nums1, nums2 []int
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		fields := strings.Fields(scanner.Text())
		num1, err1 := strconv.Atoi(fields[0])
		num2, err2 := strconv.Atoi(fields[1])
		if err1 != nil || err2 != nil {
			fmt.Println("error parsing nums")
			os.Exit(1)
		}
		nums1 = append(nums1, num1)
		nums2 = append(nums2, num2)
	}
	sort.Slice(nums1, func(i, j int) bool {
		return nums1[i] < nums1[j]
	})
	sort.Slice(nums2, func(i, j int) bool {
		return nums2[i] < nums2[j]
	})
	return nums1, nums2
}

func absInt(n int) int {
	if n < 0 {
		return -n
	}
	return n
}
