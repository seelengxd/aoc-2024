package main

import (
	_ "embed"
	"flag"
	"fmt"
	"strings"

	"github.com/seelengxd/aoc-2024/parse"
)

//go:embed input.txt
var input string

func processInput(input string) (grid [][]rune, moves string) {
	parts := strings.Split(strings.TrimSpace(input), "\n\n")
	grid = parse.ParseGrid(parts[0])
	moves = strings.ReplaceAll(parts[1], "\n", "")
	return grid, moves
}

func isEmpty(x, y int, grid [][]rune) bool {
	m, n := len(grid), len(grid[0])
	return x >= 0 && x < m && y >= 0 && y < n && grid[x][y] == '.'
}

func isWall(x, y int, grid [][]rune) bool {
	m, n := len(grid), len(grid[0])
	return x >= 0 && x < m && y >= 0 && y < n && grid[x][y] == '#'
}

func moveTo(x, y, dx, dy int, grid [][]rune, dryRun bool) bool {
	if isEmpty(x, y, grid) {
		return true
	}
	if isWall(x, y, grid) {
		return false
	}
	if dx != 0 && (grid[x][y] == '[' || grid[x][y] == ']') {
		x1, y1 := x, y
		var y2 int
		if grid[x][y] == '[' {
			y2 = y + 1
		} else {
			y1 = y - 1
			y2 = y
		}

		if dryRun {
			return moveTo(x1+dx, y1+dy, dx, dy, grid, dryRun) &&
				moveTo(x1+dx, y2+dy, dx, dy, grid, dryRun)
		}

		// if not dryrun, first "dryrun" it to see if both parts can be moved.
		if moveTo(x1+dx, y1+dy, dx, dy, grid, true) && moveTo(x1+dx, y2+dy, dx, dy, grid, true) {
			moveTo(x1+dx, y1+dy, dx, dy, grid, false)
			moveTo(x1+dx, y2+dy, dx, dy, grid, false)
			grid[x1][y1], grid[x1+dx][y1+dy] = grid[x1+dx][y1+dy], grid[x1][y1]
			grid[x1][y2], grid[x1+dx][y2+dy] = grid[x1+dx][y2+dy], grid[x1][y2]
			return true
		}
		return false
	}
	if isEmpty(x+dx, y+dy, grid) {
		if !dryRun {
			grid[x][y], grid[x+dx][y+dy] = grid[x+dx][y+dy], grid[x][y]
		}
		return true
	}
	if moveTo(x+dx, y+dy, dx, dy, grid, dryRun) {
		if !dryRun {
			grid[x][y], grid[x+dx][y+dy] = grid[x+dx][y+dy], grid[x][y]
		}
		return true
	}
	return false

}

func solve(grid [][]rune, moves string) (answer int) {
	// find starting positiion
	var x, y int
	for i, row := range grid {
		for j, cell := range row {
			if cell == '@' {
				x, y = i, j
				break
			}
		}
	}

	for _, move := range moves {
		var dx, dy int
		switch move {
		case '^':
			dx, dy = -1, 0
		case 'v':
			dx, dy = 1, 0
		case '<':
			dx, dy = 0, -1
		case '>':
			dx, dy = 0, 1
		}
		if moveTo(x, y, dx, dy, grid, false) {
			x, y = x+dx, y+dy
		}
	}

	for i, row := range grid {
		for j, cell := range row {
			if cell == 'O' || cell == '[' {
				answer += 100*i + j
			}
		}
	}

	return answer
}

func part1(grid [][]rune, moves string) (answer int) {
	return solve(grid, moves)
}

func part2(grid [][]rune, moves string) (answer int) {
	newGrid := make([][]rune, len(grid))
	for i, row := range grid {
		newRow := make([]rune, len(row)*2)
		for j, cell := range row {
			switch cell {
			case '#':
				newRow[j*2] = '#'
				newRow[j*2+1] = '#'
			case 'O':
				newRow[j*2] = '['
				newRow[j*2+1] = ']'
			case '.':
				newRow[j*2] = '.'
				newRow[j*2+1] = '.'
			case '@':
				newRow[j*2] = '@'
				newRow[j*2+1] = '.'
			}
		}
		newGrid[i] = newRow
	}
	return solve(newGrid, moves)
}

func main() {
	var part int
	flag.IntVar(&part, "part", 1, "part 1 or 2")
	flag.Parse()

	grid, moves := processInput(input)

	if part == 1 {
		fmt.Printf("Part 1: %d\n", part1(grid, moves))
	} else {
		fmt.Printf("Part 2: %d\n", part2(grid, moves))
	}
}
