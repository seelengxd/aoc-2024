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

type Stone int64

func MakeStone(value string) Stone {
	return Stone(int64(utils.Atoi(value)))
}

const STONE_MULTIPLIER = 2024

var dp = map[Stone][]Stone{}

func (s Stone) Replace() []Stone {
	if val, ok := dp[s]; ok {
		return val
	}

	// zero case
	if s == 0 {
		dp[s] = []Stone{1}
		return dp[s]
	}
	// split
	stoneStr := fmt.Sprintf("%v", s)
	if len(stoneStr)%2 == 0 {
		dp[s] = []Stone{MakeStone(stoneStr[:len(stoneStr)/2]), MakeStone(stoneStr[len(stoneStr)/2:])}
		return dp[s]
	}

	// otherwise, multiply by 2024
	dp[s] = []Stone{s * STONE_MULTIPLIER}
	return dp[s]
}

func StoneSlice(nums []string) []Stone {
	result := make([]Stone, len(nums))
	for i, num := range nums {
		result[i] = MakeStone(num)
	}
	return result
}

func CountStones(initialStones []Stone, iterations int) (answer int) {
	currStoneMap := map[Stone]int{}
	for _, stone := range initialStones {
		currStoneMap[stone]++
	}

	for i := 0; i < iterations; i++ {
		newStoneMap := map[Stone]int{}

		for stone, count := range currStoneMap {
			newStones := stone.Replace()
			for _, newStone := range newStones {
				if _, ok := newStoneMap[newStone]; !ok {
					newStoneMap[newStone] = 0
				}
				newStoneMap[newStone] += count
			}
		}
		currStoneMap = newStoneMap
	}

	for _, v := range currStoneMap {
		answer += v
	}
	return answer
}

func part1(stones []Stone) int {
	return CountStones(stones, 25)
}

func part2(stones []Stone) int {
	return CountStones(stones, 75)
}

func main() {
	nums := strings.Split(strings.TrimSpace(input), " ")
	stones := StoneSlice(nums)

	var part int
	flag.IntVar(&part, "part", 1, "part 1 or 2")
	flag.Parse()

	if part == 1 {
		fmt.Printf("Part 1: %d\n", part1(stones))
	} else {
		fmt.Printf("Part 2: %d\n", part2(stones))
	}

}
