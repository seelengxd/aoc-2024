package main

import (
	_ "embed"
	"flag"
	"fmt"
	"strconv"
	"strings"
)

//go:embed input.txt
var input string

func processInput(input string) *[][]int {
	grid := make([][]int, 0)
	for _, line := range strings.Split(input, "\n") {
		if line == "" {
			continue
		}
		row := strings.Fields(strings.TrimSpace(line))
		processed_row := make([]int, len(row))
		for i, num := range row {
			converted_num, err := strconv.Atoi(num)
			if err != nil {
				panic(fmt.Sprintf("%s was not a number", num))
			}
			processed_row[i] = converted_num
		}
		grid = append(grid, processed_row)

	}
	return &grid

}

type direction int

const (
	INCREASING direction = 1
	DECREASING direction = -1
)

func iterateReport(report *[]int, ignoredIndex int) func(yield func(int, int) bool) {
	return func(yield func(int, int) bool) {
		index := -1
		for i, v := range *report {
			if i != ignoredIndex {
				index++
				if !yield(index, v) {
					return
				}
			}
		}
		return
	}
}

func isReportSafeWithIgnore(report *[]int, ignoredIndex int) (safe bool) {
	direction := INCREASING
	prev := -1
	safe = true
	for i, num := range iterateReport(report, ignoredIndex) {
		if i == 0 {
			prev = num
			continue
		}
		if i == 1 {
			if num > prev {
				direction = INCREASING
			} else if num < prev {
				direction = DECREASING
			} else {
				return false
			}
		}
		diff := num - prev
		if !(1 <= diff*int(direction) && diff*int(direction) <= 3) {
			safe = false
			return safe
		}
		prev = num
	}
	return safe
}

func isReportSafe(report *[]int) bool {
	return isReportSafeWithIgnore(report, -1)
}

func part1(grid *[][]int) (answer int) {
	for _, row := range *grid {
		if isReportSafe(&row) {
			answer++
		}
	}
	return answer
}

func part2(grid *[][]int) (answer int) {
	for _, row := range *grid {
		for i := -1; i < len(row); i++ {
			if isReportSafeWithIgnore(&row, i) {
				answer++
				break
			}
		}
	}
	return answer
}

func main() {
	var part int
	flag.IntVar(&part, "part", 1, "part 1 or 2")
	flag.Parse()

	grid := processInput(input)
	if part == 1 {
		fmt.Printf("Part 1: %d\n", part1(grid))
	} else {
		fmt.Printf("Part 2: %d\n", part2(grid))
	}
}
