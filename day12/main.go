package main

import (
	_ "embed"
	"flag"
	"fmt"

	"github.com/seelengxd/aoc-2024/parse"
)

//go:embed input.txt
var input string

func dfs(i, j, id int, plantType rune, regionGrid [][]int, grid [][]rune) {
	m, n := len(grid), len(grid[0])

	if i < 0 || i >= m || j < 0 || j >= n {
		return
	}
	if grid[i][j] != plantType {
		return
	}
	if regionGrid[i][j] != 0 {
		return
	}

	regionGrid[i][j] = id
	dfs(i+1, j, id, plantType, regionGrid, grid)
	dfs(i-1, j, id, plantType, regionGrid, grid)
	dfs(i, j+1, id, plantType, regionGrid, grid)
	dfs(i, j-1, id, plantType, regionGrid, grid)
}

func getRegionGrid(grid [][]rune) [][]int {
	m, n := len(grid), len(grid[0])
	regionGrid := make([][]int, m)
	for i := 0; i < m; i++ {
		regionGrid[i] = make([]int, n)
	}
	currId := 1

	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if regionGrid[i][j] != 0 {
				continue
			}
			dfs(i, j, currId, grid[i][j], regionGrid, grid)
			currId++
		}
	}
	return regionGrid

}

func isRegion(i, j, regionId int, regionGrid [][]int) bool {
	m, n := len(regionGrid), len(regionGrid[0])
	if i < 0 || i >= m || j < 0 || j >= n {
		return false
	}
	return regionGrid[i][j] == regionId
}

func part1(grid [][]rune) (answer int) {
	m, n := len(grid), len(grid[0])
	regionGrid := getRegionGrid(grid)

	areaMap := map[int]int{}
	perimeterMap := map[int]int{}

	incrementPerimeter := func(i, j, regionId int) {
		if !isRegion(i, j, regionId, regionGrid) {
			perimeterMap[regionId]++
		}
	}

	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			regionId := regionGrid[i][j]
			areaMap[regionId]++
			incrementPerimeter(i+1, j, regionId)
			incrementPerimeter(i-1, j, regionId)
			incrementPerimeter(i, j+1, regionId)
			incrementPerimeter(i, j-1, regionId)
		}
	}
	for regionId, area := range areaMap {
		answer += area * perimeterMap[regionId]
	}

	return answer
}

func part2(grid [][]rune) (answer int) {
	m, n := len(grid), len(grid[0])
	regionGrid := getRegionGrid(grid)

	areaMap := map[int]int{}
	sideMap := map[int]int{}

	incrementSide := func(i, j, regionId int) {
		if !isRegion(i, j, regionId, regionGrid) {
			sideMap[regionId]++
		}
	}

	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			regionId := regionGrid[i][j]
			areaMap[regionId]++

			// left side
			if !isRegion(i-1, j, regionId, regionGrid) || isRegion(i-1, j-1, regionId, regionGrid) {
				incrementSide(i, j-1, regionId)
			}

			// right side
			if !isRegion(i-1, j, regionId, regionGrid) || isRegion(i-1, j+1, regionId, regionGrid) {
				incrementSide(i, j+1, regionId)
			}

			// bottom
			if !isRegion(i, j-1, regionId, regionGrid) || isRegion(i+1, j-1, regionId, regionGrid) {
				incrementSide(i+1, j, regionId)
			}

			// top
			if !isRegion(i, j+1, regionId, regionGrid) || isRegion(i-1, j+1, regionId, regionGrid) {
				incrementSide(i-1, j, regionId)
			}
		}
	}

	for regionId, area := range areaMap {
		answer += area * sideMap[regionId]
	}

	return answer
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
