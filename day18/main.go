package main

import (
	_ "embed"
	"flag"
	"fmt"
	"strings"

	"github.com/seelengxd/aoc-2024/utils"
)

//go:embed input.txt
var input string

type Position struct {
	X, Y int
}

func parseInput(input string) []Position {
	inputs := []Position{}
	for _, line := range strings.Split(strings.TrimSpace(input), "\n") {
		var position Position
		fmt.Sscanf(line, "%d,%d", &position.Y, &position.X)
		inputs = append(inputs, position)
	}
	return inputs
}

const GRID_SIZE = 71
const FIRST_X_BYTES = 1024

func solve(program []Position, byteCount int) (answer int) {
	set := map[Position]bool{}
	for i, pos := range program {
		if i < byteCount {
			set[pos] = true
		} else {
			break
		}
	}

	isSafe := func(x, y int) bool {
		if x < 0 || y < 0 || x >= GRID_SIZE || y >= GRID_SIZE {
			return false
		}
		_, ok := set[Position{x, y}]
		return !ok
	}

	curr := []Position{{0, 0}}
	next := []Position{}

	steps := 0
	visited := map[Position]bool{}
	found := false
	for len(curr) > 0 && !found {
		for _, pos := range curr {
			if visited[pos] {
				continue
			}
			if !isSafe(pos.X, pos.Y) {
				continue
			}

			visited[pos] = true
			if pos.X == GRID_SIZE-1 && pos.Y == GRID_SIZE-1 {
				found = true
				break
			}
			for _, direction := range utils.NON_DIAGONAL_DIRECTIONS {
				nextPos := Position{pos.X + direction.X, pos.Y + direction.Y}
				_, ok := visited[nextPos]
				if isSafe(nextPos.X, nextPos.Y) && !ok {
					next = append(next, nextPos)
				}
			}
		}
		if found {
			answer = steps
			break
		}
		steps++
		curr, next = next, []Position{}
	}

	return answer

}

func part1(program []Position) int {
	return solve(program, FIRST_X_BYTES)
}

func part2(program []Position) string {
	bytes := FIRST_X_BYTES + 1
	for {
		answer := solve(program, bytes)
		if answer == 0 {
			return fmt.Sprintf("%d,%d", program[bytes-1].Y, program[bytes-1].X)
		}
		bytes++
	}
}

func main() {
	var part int
	flag.IntVar(&part, "part", 1, "part 1 or 2")
	flag.Parse()

	program := parseInput(input)

	if part == 1 {
		fmt.Printf("Part 1: %d\n", part1(program))
	} else {
		fmt.Printf("Part 2: %s\n", part2(program))
	}
}
