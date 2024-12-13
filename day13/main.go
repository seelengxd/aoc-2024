package main

import (
	_ "embed"
	"flag"
	"fmt"
	"strings"
)

//go:embed input.txt
var input string

type Input struct {
	x1, x2, y1, y2, X, Y int
}

func processInput(input string) []Input {
	inputs := strings.Split(strings.TrimSpace(input), "\n\n")
	result := make([]Input, len(inputs))
	for i, inputStr := range inputs {
		newInput := &result[i]
		fmt.Sscanf(inputStr, "Button A: X+%d, Y+%d\nButton B: X+%d, Y+%d\nPrize: X=%d, Y=%d", &newInput.x1, &newInput.y1, &newInput.x2, &newInput.y2, &newInput.X, &newInput.Y)
	}
	return result
}

func GCD(a, b int) int {
	for b != 0 {
		a, b = b, a%b
	}
	return a
}

func solve(input Input) int {
	x1, x2, y1, y2, X, Y := input.x1, input.x2, input.y1, input.y2, input.X, input.Y
	// |x1 x2| |a| = |X|
	// |y1 y2| |b|   |Y|

	determinant := x1*y2 - x2*y1

	if determinant != 0 {
		// there is only one solution
		// | a | = | y2 -x2 | |X|
		// | b |   | y1  x1 | |Y|
		a, b := (y2*X - x2*Y), (-y1*X + x1*Y)
		if a%determinant == 0 && b%determinant == 0 && a/determinant >= 0 && b/determinant >= 0 {
			a, b = a/determinant, b/determinant
			return a*3 + b
		}
		return 0
	}

	// one is the multiple of the other
	// first calculate which one has more ROI
	if y2/y1 >= 3 {
		a, b := 0, Y/y2
		for i := 0; i < GCD(y2, y1)*10000000; i++ {
			for a*x1+b*x2 < X && a*y1+b*y2 < Y {
				a++
			}
			if a*x1+b*x2 == X && a*y1+b*y2 == Y {
				return a*3 + b
			}
			b -= 1
		}
		return 0
	}

	a, b := Y/y1, 0
	for i := 0; i < GCD(y2, y1)*10000000; i++ {
		for a*x1+b*x2 < X && a*y1+b*y2 < Y {
			b++
		}
		if a*x1+b*x2 == X && a*y1+b*y2 == Y {
			return a*3 + b
		}
		a -= 1
	}
	return 0

}

func part1(inputs []Input) (answer int) {
	for _, input := range inputs {
		answer += solve(input)
	}
	return answer
}

func part2(inputs []Input) (answer int) {
	for _, input := range inputs {
		input.X += 10000000000000
		input.Y += 10000000000000
		answer += solve(input)
	}
	return answer
}

func main() {
	var part int
	flag.IntVar(&part, "part", 1, "part 1 or 2")
	flag.Parse()

	inputs := processInput(input)

	if part == 1 {
		fmt.Printf("Part 1: %d\n", part1(inputs))
	} else {
		fmt.Printf("Part 2: %d\n", part2(inputs))
	}
}
