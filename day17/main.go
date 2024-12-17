package main

import (
	_ "embed"
	"flag"
	"fmt"
	"math"
	"slices"
	"strings"

	"github.com/seelengxd/aoc-2024/parse"
)

//go:embed input.txt
var input string

type Program struct {
	A, B, C int64
	output  []int64
	program []int64
	pointer int64
}

func makeProgram(a, b, c int64, program []int64) *Program {
	return &Program{A: a, B: b, C: c, program: program, output: []int64{}, pointer: 0}
}

func (r *Program) combo(op int64) int64 {
	switch op {
	case 0, 1, 2, 3:
		return op
	case 4:
		return r.A
	case 5:
		return r.B
	case 6:
		return r.C
	default:
		panic("Invalid combo operand")
	}
}

func (r *Program) adv(operand int64) {
	r.A /= 1 << r.combo(operand)
}

func (r *Program) bxl(operand int64) {
	r.B ^= operand
}

func (r *Program) bst(operand int64) {
	r.B = r.combo(operand) % 8
}

func (r *Program) jnz(operand int64) {
	if r.A == 0 {
		return
	}
	r.pointer = operand - 2
}

func (r *Program) bxc(int64) {
	r.B ^= r.C
}

func (r *Program) out(operand int64) {
	r.output = append(r.output, r.combo(operand)%8)
}

func (r *Program) bdv(operand int64) {
	r.B = r.A / (1 << r.combo(operand))
}

func (r *Program) cdv(operand int64) {
	r.C = r.A / (1 << r.combo(operand))
}

func (r *Program) doOp(op int64, operand int64) {
	switch op {
	case 0:
		r.adv(operand)
	case 1:
		r.bxl(operand)
	case 2:
		r.bst(operand)
	case 3:
		r.jnz(operand)
	case 4:
		r.bxc(operand)
	case 5:
		r.out(operand)
	case 6:
		r.bdv(operand)
	case 7:
		r.cdv(operand)
	}
	r.pointer += 2
}

func (r *Program) validPointer() bool {
	return r.pointer >= 0 && r.pointer < int64(len(r.program))
}

func (r *Program) run() string {
	for r.validPointer() {
		op, operand := r.program[r.pointer], r.program[r.pointer+1]
		r.doOp(op, operand)
	}
	var outputStrings []string
	for _, num := range r.output {
		outputStrings = append(outputStrings, fmt.Sprintf("%d", num))
	}
	return strings.Join(outputStrings, ",")
}

func parseInput(input string) (registers Program) {
	parts := strings.Split(strings.TrimSpace(input), "\n\n")

	fmt.Sscanf(parts[0], "Register A: %d\nRegister B: %d\nRegister C: %d\n", &registers.A, &registers.B, &registers.C)
	registers.output = []int64{}
	registers.program = parse.Int64s(strings.TrimPrefix(parts[1], "Program: "), ",")
	registers.pointer = 0

	return registers
}

func part1(program Program) string {
	return program.run()
}

func part2(program Program) (answer int64) {
	curr := [][]int64{{}}
	next := [][]int64{}

	getAnswer := func(parts []int64) int64 {
		var answer int64
		for _, digit := range parts {
			answer <<= 3
			answer += digit
		}
		return answer

	}

	// observation: when u increase the answer by 3 bits, the output is the old output plus the new digit in front
	// lockpicking strategy, brute force 3 bits at a time
	for _, digit := range slices.Backward(program.program) {
		for _, parts := range curr {
			for guess := 0; guess < 8; guess++ {
				tempAnswer := getAnswer(parts) << 3
				tempAnswer += int64(guess)
				tempProgram := makeProgram(tempAnswer, 0, 0, program.program)
				tempProgram.run()
				if tempProgram.output[0] == digit {
					next = append(next, slices.Concat([]int64{}, parts, []int64{int64(guess)}))
				}
			}
		}
		curr, next = next, [][]int64{}
	}

	answer = math.MaxInt64
	for _, parts := range curr {
		answer = min(answer, getAnswer(parts))
	}
	return answer
}

func main() {
	var part int
	flag.IntVar(&part, "part", 1, "part 1 or 2")
	flag.Parse()

	program := parseInput(input)

	if part == 1 {
		fmt.Printf("Part 1: %s\n", part1(program))
	} else {
		fmt.Printf("Part 2: %d\n", part2(program))
	}
}
