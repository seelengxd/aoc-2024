package main

import (
	_ "embed"
	"flag"
	"fmt"

	"github.com/seelengxd/aoc-2024/parse"
	"github.com/seelengxd/aoc-2024/utils"
)

//go:embed input.txt
var input string

func processInput(input string) [][]rune {
	return parse.ParseGrid(input)
}

const XMAS string = "XMAS"
const MAS string = "MAS"

func traverse(lines [][]rune, index int, direction utils.Direction, x int, y int, target string) bool {
	if x < 0 || y < 0 || x >= len(lines) || y >= len(lines[0]) {
		return false
	}
	if lines[x][y] != rune(target[index]) {
		return false
	}
	if index == len(target)-1 {
		return true
	} else {
		return traverse(lines, index+1, direction, x+direction.X, y+direction.Y, target)
	}
}

func traverseXMAS(lines [][]rune, index int, direction utils.Direction, x int, y int) bool {
	return traverse(lines, index, direction, x, y, XMAS)
}

func traverseMAS(lines [][]rune, index int, direction utils.Direction, x int, y int) bool {
	return traverse(lines, index, direction, x, y, MAS)
}

func part1(lines [][]rune) (answer int) {
	for i, line := range lines {
		for j, char := range line {
			if char == rune(XMAS[0]) {
				for _, direction := range utils.DIRECTIONS {
					if traverseXMAS(lines, 0, direction, i, j) {
						answer++
					}
				}
			}
		}
	}
	return answer
}

func part2(lines [][]rune) (answer int) {
	for i, line := range lines {
		for j, char := range line {
			// suppose (i, j) is the center of MAS cross
			if char == 'A' {
				first_line := traverseMAS(lines, 0, utils.DIRECTIONS["up_left"], i+1, j+1) || traverseMAS(lines, 0, utils.DIRECTIONS["down_right"], i-1, j-1)
				second_line := traverseMAS(lines, 0, utils.DIRECTIONS["up_right"], i+1, j-1) || traverseMAS(lines, 0, utils.DIRECTIONS["down_left"], i-1, j+1)
				if first_line && second_line {
					answer++
				}
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
