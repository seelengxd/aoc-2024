package main

import (
	_ "embed"
	"flag"
	"fmt"
	"strings"
)

//go:embed input.txt
var input string

func parseInput(input string) (towels []string, designs []string) {
	parts := strings.Split(strings.TrimSpace(input), "\n\n")
	towels = strings.Split(parts[0], ", ")
	designs = strings.Split(parts[1], "\n")
	return towels, designs
}

func part1(towels, designs []string) (answer int) {
	possible := map[string]bool{"": true}
	var dp func(string) bool
	dp = func(design string) bool {
		if _, ok := possible[design]; ok {
			return possible[design]
		}

		for _, towel := range towels {
			if strings.HasSuffix(design, towel) && dp(design[:len(design)-len(towel)]) {
				possible[design] = true
				return true
			}
		}

		possible[design] = false
		return false
	}
	for _, design := range designs {
		if dp(design) {
			answer++
		}
	}
	return answer
}

func part2(towels, designs []string) (answer int) {
	possible := map[string]int{"": 1}
	var dp func(string) int
	dp = func(design string) int {
		if _, ok := possible[design]; ok {
			return possible[design]
		}

		possible[design] = 0
		for _, towel := range towels {
			if strings.HasSuffix(design, towel) && dp(design[:len(design)-len(towel)]) > 0 {
				possible[design] += dp(design[:len(design)-len(towel)])
			}
		}
		return possible[design]
	}
	for _, design := range designs {
		answer += dp(design)
	}
	return answer

}

func main() {
	var part int
	flag.IntVar(&part, "part", 1, "part 1 or 2")
	flag.Parse()

	towels, designs := parseInput(input)

	if part == 1 {
		fmt.Printf("Part 1: %d\n", part1(towels, designs))
	} else {
		fmt.Printf("Part 2: %d\n", part2(towels, designs))
	}
}
