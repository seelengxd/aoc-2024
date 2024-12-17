package parse

import (
	"fmt"
	"strconv"
	"strings"
)

func ParseLines(input string) (lines []string) {
	return strings.Split(strings.TrimSpace(input), "\n")
}

func ParseIntSlice(line string, delimiter string) (nums []int) {
	parts := strings.Split(line, delimiter)
	for _, part := range parts {
		if part == "" {
			continue
		}
		num, err := strconv.Atoi(part)
		if err != nil {
			panic(fmt.Sprintf("%v was not a number", part))
		}
		nums = append(nums, num)
	}
	return nums
}

func Int64s(line string, delimiter string) (nums []int64) {
	parts := strings.Split(line, delimiter)
	for _, part := range parts {
		if part == "" {
			continue
		}
		numInt, err := strconv.Atoi(part)
		if err != nil {
			panic(fmt.Sprintf("%v was not a number", part))
		}
		num := int64(numInt)
		if err != nil {
			panic(fmt.Sprintf("%v was not a number", part))
		}
		nums = append(nums, num)
	}
	return nums
}

func ParseGrid(input string) [][]rune {
	lines := make([][]rune, 0)
	for _, line := range strings.Split(strings.TrimSpace(input), "\n") {
		row := make([]rune, len(line))
		for i, char := range line {
			row[i] = char
		}
		lines = append(lines, row)
	}
	return lines
}
