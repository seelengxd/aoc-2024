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

// turn right by 90 degrees each time
var changeDirectionMap = map[utils.Direction]utils.Direction{
	utils.DIRECTIONS["up"]:    utils.DIRECTIONS["right"],
	utils.DIRECTIONS["right"]: utils.DIRECTIONS["down"],
	utils.DIRECTIONS["down"]:  utils.DIRECTIONS["left"],
	utils.DIRECTIONS["left"]:  utils.DIRECTIONS["up"],
}

var directionToIndex = map[utils.Direction]int{
	utils.DIRECTIONS["up"]:    0,
	utils.DIRECTIONS["right"]: 1,
	utils.DIRECTIONS["down"]:  2,
	utils.DIRECTIONS["left"]:  3,
}

func processInput(input string) [][]rune {
	return parse.ParseGrid(input)
}

func isOutOfBounds(grid [][]rune, x int, y int) bool {
	return x < 0 || y < 0 || x >= len(grid) || y >= len(grid[0])
}

func isObstacle(grid [][]rune, x int, y int) bool {
	return !isOutOfBounds(grid, x, y) && grid[x][y] == '#'
}

func getCurrentPosition(grid [][]rune) (currX int, currY int) {
	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid[0]); j++ {
			if grid[i][j] == '^' {
				currX, currY = i, j
			}
		}
	}
	return currX, currY
}

func part1(grid [][]rune) (answer int) {
	currDirection := utils.DIRECTIONS["up"]
	currX, currY := getCurrentPosition(grid)

	for !isOutOfBounds(grid, currX, currY) {
		if grid[currX][currY] != 'X' {
			grid[currX][currY] = 'X'
			answer++
		}
		if isObstacle(grid, currX+currDirection.X, currY+currDirection.Y) {
			currDirection = changeDirectionMap[currDirection]
		}
		currX += currDirection.X
		currY += currDirection.Y
	}
	return answer
}

func hasLoop(grid [][]rune) bool {
	currDirection := utils.DIRECTIONS["up"]
	currX, currY := getCurrentPosition(grid)

	visited := make([][][]bool, len(grid))
	for i := range visited {
		visited[i] = make([][]bool, len(grid[0]))
		for j := range visited[i] {
			visited[i][j] = make([]bool, 4)
		}
	}

	for !isOutOfBounds(grid, currX, currY) {
		for isObstacle(grid, currX+currDirection.X, currY+currDirection.Y) {
			currDirection = changeDirectionMap[currDirection]
			// if you have changed direction here before to the same direction
			// it is a loop
			if visited[currX][currY][directionToIndex[currDirection]] {
				return true
			}
			visited[currX][currY][directionToIndex[currDirection]] = true
		}
		currX += currDirection.X
		currY += currDirection.Y
	}

	return false
}

func part2(grid [][]rune) (answer int) {
	// Optimisation: new obstacles must be on the old path to matter
	currX, currY := getCurrentPosition(grid)
	part1(grid)
	grid[currX][currY] = '^'

	for i := range grid {
		for j, c := range grid[i] {
			if c != 'X' {
				continue
			}
			grid[i][j] = '#'
			if hasLoop(grid) {
				answer++
			}
			grid[i][j] = '.'
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
