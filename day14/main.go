package main

import (
	_ "embed"
	"flag"
	"fmt"
	"os"
	"strings"

	tea "github.com/charmbracelet/bubbletea"
)

//go:embed input.txt
var input string

type Input struct {
	px, py, vx, vy int
}

func processInput(input string) []Input {
	inputs := strings.Split(strings.TrimSpace(input), "\n")
	result := make([]Input, len(inputs))
	for i, inputStr := range inputs {
		newInput := &result[i]
		fmt.Sscanf(inputStr, "p=%d,%d v=%d,%d", &newInput.py, &newInput.px, &newInput.vy, &newInput.vx)
	}
	return result
}

const WIDTH, HEIGHT = 101, 103

func quadrant(x, y int) int {
	// 1 2
	// 3 4
	switch {
	case x < HEIGHT/2 && y < WIDTH/2:
		return 1
	case x < HEIGHT/2 && y > WIDTH/2:
		return 2
	case x > HEIGHT/2 && y < WIDTH/2:
		return 3
	case x > HEIGHT/2 && y > WIDTH/2:
		return 4
	default:
		return 0
	}
}

func finalPosition(input Input, seconds int) (x, y int) {
	x = (input.px + input.vx*seconds) % HEIGHT
	y = (input.py + input.vy*seconds) % WIDTH
	if x < 0 {
		x += HEIGHT
	}
	if y < 0 {
		y += WIDTH
	}
	return x, y
}

func part1(inputs []Input) (answer int) {
	quadrants := [5]int{}
	for _, input := range inputs {
		quadrants[quadrant(finalPosition(input, 100))]++
	}
	answer = quadrants[1] * quadrants[2] * quadrants[3] * quadrants[4]
	return answer
}

func calculateGrid(inputs []Input, seconds int) (grid [HEIGHT][WIDTH]bool) {
	for _, input := range inputs {
		x, y := finalPosition(input, seconds)
		grid[x][y] = true
	}
	return grid
}

type model struct {
	seconds int
	inputs  []Input
}

func initialModel() model {
	return model{
		seconds: 0,
		inputs:  processInput(input),
	}
}

func (m model) Init() tea.Cmd {
	return nil
}

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "ctrl+c", "q":
			return m, tea.Quit
		case "left":
			found := false
			for !found {
				m.seconds--
				grid := calculateGrid(m.inputs, m.seconds)
				for _, row := range grid {
					s := ""
					for _, cell := range row {
						if cell {
							s += "#"
						} else {
							s += " "
						}
					}
					if strings.Contains(s, "#####") {
						found = true
						break
					}
				}
			}
		case "right":
			found := false
			for !found {
				m.seconds++
				grid := calculateGrid(m.inputs, m.seconds)
				for _, row := range grid {
					s := ""
					for _, cell := range row {
						if cell {
							s += "#"
						} else {
							s += " "
						}
					}
					if strings.Contains(s, "#####") {
						found = true
						break
					}
				}
			}

		}
	}
	return m, nil
}

func (m model) View() string {
	grid := [HEIGHT][WIDTH]bool{}
	for _, input := range m.inputs {
		x, y := finalPosition(input, m.seconds)
		grid[x][y] = true
	}
	s := fmt.Sprintf("Seconds: %d\n\n", m.seconds)
	for i := 0; i < HEIGHT; i++ {
		for j := 0; j < WIDTH; j++ {
			if grid[i][j] {
				s += "\033[32m#\033[0m"
			} else {
				s += "."
			}
		}
		s += "\n"
	}
	return s
}

func part2(inputs []Input) (answer int) {
	p := tea.NewProgram(initialModel())
	m, err := p.Run()
	if err != nil {
		fmt.Printf("Alas, there's been an error: %v", err)
		os.Exit(1)
	}
	return m.(model).seconds
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
