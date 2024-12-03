package main

import (
	_ "embed"
	"flag"
	"fmt"
	"regexp"
	"strconv"
	"strings"
)

//go:embed input.txt
var input string

func processInput(input string) *[]string {
	lines := make([]string, 0)
	lines = append(lines, strings.TrimSpace(strings.ReplaceAll(input, "\n", "")))
	return &lines
}

func getAnswerNoFilter(line string) (answer int) {
	re := regexp.MustCompile(`mul\((\d{1,3}),(\d{1,3})\)`)
	results := re.FindAllStringSubmatch(line, -1)
	for _, result := range results {
		// result example [mul(509,507) 509 507]
		firstNum, err := strconv.Atoi(result[1])
		if err != nil {
			panic(result[1] + " was not a number")
		}
		secondNum, err := strconv.Atoi(result[2])
		if err != nil {
			panic(result[2] + " was not a number")
		}
		answer += firstNum * secondNum
	}
	return answer
}

func part1(lines *[]string) (answer int) {
	for _, line := range *lines {
		answer += getAnswerNoFilter(line)
	}
	return answer
}

func part2(lines *[]string) (answer int) {
	re := regexp.MustCompile(`do\(\)(.*?)don't\(\)`)
	for _, line := range *lines {
		line = "do()" + line + "don't()"
		results := re.FindAllStringSubmatch(line, -1)

		for _, result := range results {
			answer += getAnswerNoFilter(result[0])
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