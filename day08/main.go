package main

import (
	_ "embed"
	"flag"
	"fmt"

	"github.com/seelengxd/aoc-2024/parse"
)

type Position struct {
	X int
	Y int
}

//go:embed input.txt
var input string

func withinBounds(grid [][]rune, position Position) bool {
	return position.X >= 0 && position.X < len(grid) && position.Y >= 0 && position.Y < len(grid[0])
}

func getAntennas(grid [][]rune) map[rune][]Position {
	antennas := map[rune][]Position{}
	for i, row := range grid {
		for j, cell := range row {
			if cell == '.' {
				continue
			}
			if _, ok := antennas[cell]; !ok {
				antennas[cell] = []Position{}
			}
			antennas[cell] = append(antennas[cell], Position{i, j})
		}
	}
	return antennas
}

func markAntinodeIfValid(antinodes map[Position]bool, grid [][]rune, position Position) bool {
	valid := withinBounds(grid, position)
	if valid {
		antinodes[position] = true
	}
	return valid
}

func part1(grid [][]rune) int {
	antinodes := map[Position]bool{}
	antennas := getAntennas(grid)

	for _, positions := range antennas {
		n := len(positions)
		for i := 0; i < n; i++ {
			for j := i + 1; j < n; j++ {
				diffX := positions[i].X - positions[j].X
				diffY := positions[i].Y - positions[j].Y

				firstAntinode := Position{positions[i].X + diffX, positions[i].Y + diffY}
				secondAntinode := Position{positions[j].X - diffX, positions[j].Y - diffY}
				markAntinodeIfValid(antinodes, grid, firstAntinode)
				markAntinodeIfValid(antinodes, grid, secondAntinode)
			}
		}
	}

	return len(antinodes)
}

func part2(grid [][]rune) int {
	antinodes := map[Position]bool{}
	antennas := getAntennas(grid)

	for _, positions := range antennas {
		n := len(positions)
		for i := 0; i < n; i++ {
			for j := i + 1; j < n; j++ {
				diffX := positions[i].X - positions[j].X
				diffY := positions[i].Y - positions[j].Y

				firstAntinode := positions[i]
				secondAntinode := positions[j]

				for markAntinodeIfValid(antinodes, grid, firstAntinode) {
					firstAntinode.X += diffX
					firstAntinode.Y += diffY
				}

				for markAntinodeIfValid(antinodes, grid, secondAntinode) {
					secondAntinode.X -= diffX
					secondAntinode.Y -= diffY
				}
			}
		}
	}

	return len(antinodes)

}

func main() {
	var part int
	flag.IntVar(&part, "part", 1, "part 1 or 2")
	flag.Parse()

	inputs := parse.ParseGrid(input)

	if part == 1 {
		fmt.Printf("Part 1: %d\n", part1(inputs))
	} else {
		fmt.Printf("Part 2: %d\n", part2(inputs))
	}
}
