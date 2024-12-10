package main

import (
	_ "embed"
	"flag"
	"fmt"
	"maps"

	"github.com/seelengxd/aoc-2024/parse"
)

//go:embed input.txt
var input string

type Position struct {
	X int
	Y int
}

func part1(grid [][]rune) (answer int) {
	dp := make([][][]map[Position]bool, 10)
	for i := 0; i < 10; i++ {
		dp[i] = make([][]map[Position]bool, len(grid))
		for j := 0; j < len(grid); j++ {
			dp[i][j] = make([]map[Position]bool, len(grid[0]))
			for k := 0; k < len(grid[0]); k++ {
				dp[i][j][k] = map[Position]bool{}
			}
		}
	}

	for i, row := range grid {
		for j, cell := range row {
			if cell == '9' {
				dp[9][i][j][Position{i, j}] = true
			}
		}
	}

	get := func(k, i, j int) map[Position]bool {
		if i < 0 || j < 0 || i >= len(grid) || j >= len(grid[0]) {
			return map[Position]bool{}
		}
		return dp[k][i][j]

	}
	for k := 8; k >= 0; k-- {
		for i := 0; i < len(grid); i++ {
			for j := 0; j < len(grid[0]); j++ {
				if grid[i][j] == '0'+rune(k) {
					maps.Copy(dp[k][i][j], get(k+1, i+1, j))
					maps.Copy(dp[k][i][j], get(k+1, i-1, j))
					maps.Copy(dp[k][i][j], get(k+1, i, j+1))
					maps.Copy(dp[k][i][j], get(k+1, i, j-1))
				}

			}
		}
	}

	for _, row := range dp[0] {
		for _, cell := range row {
			answer += len(cell)
		}
	}
	return answer
}

func part2(grid [][]rune) (answer int) {
	dp := make([][][]int, 10)
	for i := 0; i < 10; i++ {
		dp[i] = make([][]int, len(grid))
		for j := 0; j < len(grid); j++ {
			dp[i][j] = make([]int, len(grid[0]))
		}
	}
	for i, row := range grid {
		for j, cell := range row {
			if cell == '9' {
				dp[9][i][j] = 1
			}
		}
	}

	get := func(k, i, j int) int {
		if i < 0 || j < 0 || i >= len(grid) || j >= len(grid[0]) {
			return 0
		}
		return dp[k][i][j]

	}
	for k := 8; k >= 0; k-- {
		for i := 0; i < len(grid); i++ {
			for j := 0; j < len(grid[0]); j++ {
				if grid[i][j] == '0'+rune(k) {
					dp[k][i][j] = get(k+1, i+1, j) + get(k+1, i-1, j) + get(k+1, i, j+1) + get(k+1, i, j-1)
				}

			}
		}
	}

	for _, row := range dp[0] {
		for _, cell := range row {
			answer += cell
		}
	}
	return answer
}

func main() {
	var part int
	flag.IntVar(&part, "part", 1, "part 1 or 2")
	flag.Parse()

	if part == 1 {
		fmt.Printf("Part 1: %d\n", part1(parse.ParseGrid(input)))
	} else {
		fmt.Printf("Part 2: %d\n", part2(parse.ParseGrid(input)))
	}
}
