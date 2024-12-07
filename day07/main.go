package main

import (
	_ "embed"
	"flag"
	"fmt"
	"strconv"
	"strings"

	"github.com/seelengxd/aoc-2024/parse"
)

//go:embed input.txt
var input string

type Calibration struct {
	target int
	nums   []int
}

func Add(a, b int) int {
	return a + b
}

func Multiply(a, b int) int {
	return a * b
}

func Concat(a, b int) int {
	leftStr := strconv.Itoa(a)
	rightStr := strconv.Itoa(b)
	nextNumber, _ := strconv.Atoi(leftStr + rightStr)
	return nextNumber
}

func checkCalibration(calibration Calibration, operators []func(int, int) int) bool {
	possible := map[int]bool{calibration.nums[0]: true}
	nextPossible := map[int]bool{}
	for _, num := range calibration.nums[1:] {
		for possibleNum := range possible {
			for _, operator := range operators {
				if nextNum := operator(possibleNum, num); nextNum <= calibration.target {
					nextPossible[nextNum] = true
				}
			}
		}
		possible, nextPossible = nextPossible, map[int]bool{}
	}
	_, ok := possible[calibration.target]
	return ok
}

func part1(calibrations []Calibration) (answer int) {
	var operators = []func(int, int) int{Add, Multiply}

	for _, calibration := range calibrations {
		if checkCalibration(calibration, operators) {
			answer += calibration.target
		}
	}
	return answer
}

func part2(calibrations []Calibration) (answer int) {
	var operators = []func(int, int) int{Add, Multiply, Concat}

	for _, calibration := range calibrations {
		if checkCalibration(calibration, operators) {
			answer += calibration.target
		}
	}
	return answer
}

func processInput(input string) []Calibration {
	calibrations := make([]Calibration, 0)
	for _, line := range strings.Split(strings.TrimSpace(input), "\n") {
		split := strings.Split(line, ": ")
		left, right := split[0], split[1]
		target, _ := strconv.Atoi(left)
		nums := parse.ParseIntSlice(right, " ")
		calibrations = append(calibrations, Calibration{target, nums})
	}
	return calibrations
}

func main() {
	var part int
	flag.IntVar(&part, "part", 1, "part 1 or 2")
	flag.Parse()

	inputs := processInput(input)

	if part == 1 {
		fmt.Printf("Part 1: %d\n", part1(inputs))
	} else {
		fmt.Printf("Part 2: %d\n", part2(inputs))
	}
}
