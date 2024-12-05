package main

import (
	_ "embed"
	"flag"
	"fmt"
	"sort"
	"strconv"
	"strings"
)

//go:embed input.txt
var input string

func processInput(input string) (left []int, right []int) {
	for _, line := range strings.Split(input, "\n") {
		if line == "" {
			continue
		}
		line = strings.TrimSpace(line)
		numbers := strings.Fields(line)
		leftNumber, err := strconv.Atoi(numbers[0])
		if err != nil {
			panic(fmt.Sprintf("%s was not a number", numbers[0]))
		}
		left = append(left, leftNumber)
		rightNumber, err := strconv.Atoi(numbers[1])
		if err != nil {
			panic(fmt.Sprintf("%s was not a number", numbers[1]))
		}
		right = append(right, rightNumber)
	}
	return left, right
}

func part1(left []int, right []int) (answer int) {
	sort.Ints(left)
	sort.Ints(right)
	for i, leftNumber := range left {
		rightNumber := right[i]
		difference := rightNumber - leftNumber
		if difference < 0 {
			answer += -difference
		} else {
			answer += difference
		}
	}
	return answer
}

func part2(left []int, right []int) (answer int) {
	rightCounter := make(map[int]int)
	for _, rightNumber := range right {
		rightCounter[rightNumber]++
	}
	for _, number := range left {
		answer += number * rightCounter[number]
	}
	return answer
}

func main() {
	var part int
	flag.IntVar(&part, "part", 1, "part 1 or 2")
	flag.Parse()

	left, right := processInput(input)

	if part == 1 {
		fmt.Printf("Part 1: %d\n", part1(left, right))
	} else {
		fmt.Printf("Part 2: %d\n", part2(left, right))
	}
}
