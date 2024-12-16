package main

import (
	"container/heap"
	_ "embed"
	"flag"
	"fmt"

	"github.com/seelengxd/aoc-2024/ds"
	"github.com/seelengxd/aoc-2024/parse"
	"github.com/seelengxd/aoc-2024/utils"
)

//go:embed input.txt
var input string

type Reindeer struct {
	x, y      int
	direction utils.Direction
}

type ReindeerItem = ds.Item[Reindeer]

const MOVE_COST = 1
const TURN_COST = 1000

func solve(grid [][]rune) (cache map[Reindeer]int, answer int) {
	cache = map[Reindeer]int{}
	q := ds.PriorityQueue[Reindeer]{}

	// get initial reindeer position and target position
	var sx, sy, tx, ty int
	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid[0]); j++ {
			if grid[i][j] == 'S' {
				sx, sy = i, j
			}
			if grid[i][j] == 'E' {
				tx, ty = i, j
			}
		}
	}

	heap.Init(&q)
	heap.Push(&q, ds.CreateItem(Reindeer{x: sx, y: sy, direction: utils.DIRECTIONS["right"]}, 0))

	for q.Len() != 0 {
		item := heap.Pop(&q).(*ReindeerItem)
		pos := item.Value()
		if _, ok := cache[pos]; ok {
			continue
		}
		if pos.x < 0 || pos.y < 0 || pos.x >= len(grid) || pos.y >= len(grid[0]) || grid[pos.x][pos.y] == '#' {
			continue
		}

		cost := item.Priority()
		if pos.x == tx && pos.y == ty {
			cache[pos] = cost
			if answer == 0 {
				answer = cost
			}
		}

		cache[pos] = cost

		// move forward
		heap.Push(&q, ds.CreateItem(Reindeer{x: pos.x + pos.direction.X, y: pos.y + pos.direction.Y, direction: pos.direction}, cost+MOVE_COST))
		// turn 2 directions
		heap.Push(&q, ds.CreateItem(Reindeer{x: pos.x, y: pos.y, direction: utils.ChangeDirectionAntiClockwiseMap[pos.direction]}, cost+TURN_COST))
		heap.Push(&q, ds.CreateItem(Reindeer{x: pos.x, y: pos.y, direction: utils.ChangeDirectionClockwiseMap[pos.direction]}, cost+TURN_COST))
	}
	return cache, answer
}

func part1(grid [][]rune) (answer int) {
	_, answer = solve(grid)
	return answer
}

func part2(grid [][]rune) (answer int) {
	cache, cost := solve(grid)
	posCache := map[Reindeer]bool{}
	visited := map[ds.Item[Reindeer]]bool{}
	stack := ds.Stack[*ds.Item[Reindeer]]{}

	var tx, ty int
	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid[0]); j++ {
			if grid[i][j] == 'E' {
				tx, ty = i, j
			}
		}
	}
	for _, direction := range utils.NON_DIAGONAL_DIRECTIONS {
		pos := Reindeer{x: tx, y: ty, direction: direction}
		if cache[pos] == cost {
			stack.Push(ds.CreateItem(pos, cost))
		}

	}

	for len(stack) != 0 {
		item := stack.Pop()
		pos := item.Value()
		if _, ok := visited[*item]; ok {
			continue
		}

		visited[*item] = true
		posCache[Reindeer{pos.x, pos.y, utils.DIRECTIONS["up"]}] = true
		cost := item.Priority()

		dir1, dir2 := utils.ChangeDirectionAntiClockwiseMap[pos.direction], utils.ChangeDirectionClockwiseMap[pos.direction]

		nextPos := Reindeer{x: pos.x, y: pos.y, direction: dir1}
		if cache[nextPos] == cost-TURN_COST {
			stack.Push(ds.CreateItem(nextPos, cost-TURN_COST))
		}

		nextPos = Reindeer{x: pos.x, y: pos.y, direction: dir2}
		if cache[Reindeer{x: pos.x, y: pos.y, direction: dir2}] == cost-TURN_COST {
			stack.Push(ds.CreateItem(nextPos, cost-TURN_COST))
		}

		nextPos = Reindeer{x: pos.x - pos.direction.X, y: pos.y - pos.direction.Y, direction: pos.direction}
		if cache[nextPos] == cost-MOVE_COST {
			stack.Push(ds.CreateItem(nextPos, cost-MOVE_COST))
		}
	}

	return len(posCache)
}

func main() {
	var part int
	flag.IntVar(&part, "part", 1, "part 1 or 2")
	flag.Parse()

	grid := parse.ParseGrid(input)

	if part == 1 {
		fmt.Printf("Part 1: %d\n", part1(grid))
	} else {
		fmt.Printf("Part 2: %d\n", part2(grid))
	}
}
