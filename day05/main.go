package main

import (
	_ "embed"
	"flag"
	"fmt"
	"strings"

	"github.com/seelengxd/aoc-2024/parse"
)

//go:embed input.txt
var input string

type OrderingRule struct {
	Before int
	After  int
}

type Update []int

func processInput(input string) ([]OrderingRule, []Update) {
	parts := strings.Split(input, "\n\n")
	orderingRuleString, updateString := parts[0], parts[1]
	orderingRules := make([]OrderingRule, 0)
	updates := make([]Update, 0)

	for _, line := range strings.Split(orderingRuleString, "\n") {
		var orderingRule OrderingRule
		fmt.Sscanf(line, "%d|%d", &orderingRule.Before, &orderingRule.After)
		orderingRules = append(orderingRules, orderingRule)
	}

	for _, line := range strings.Split(updateString, "\n") {
		updates = append(updates, parse.ParseIntSlice(line, ","))
	}

	return orderingRules, updates
}

func isOrderingValid(orderingRules []OrderingRule, update Update) bool {
	for i := 0; i < len(update)-1; i++ {
		for j := i + 1; j < len(update); j++ {
			for _, orderingRule := range orderingRules {
				if orderingRule.Before == update[j] && orderingRule.After == update[i] {
					return false
				}
			}
		}
	}
	return true
}

func fixUpdate(orderingRules []OrderingRule, update Update) {
	for i := 0; i < len(update)-1; i++ {
		for j := i + 1; j < len(update); j++ {
			for _, orderingRule := range orderingRules {
				if orderingRule.Before == update[j] && orderingRule.After == update[i] {
					update[i], update[j] = update[j], update[i]
				}
			}
		}
	}

}

func part1(orderings []OrderingRule, updates []Update) (answer int) {
	for _, update := range updates {
		if isOrderingValid(orderings, update) {
			answer += update[len(update)/2]
		}
	}
	return answer
}

func part2(orderings []OrderingRule, updates []Update) (answer int) {
	for _, update := range updates {
		if !isOrderingValid(orderings, update) {
			fixUpdate(orderings, update)
			answer += update[len(update)/2]
		}
	}
	return answer
}

func main() {
	var part int
	flag.IntVar(&part, "part", 1, "part 1 or 2")
	flag.Parse()

	orderings, updates := processInput(input)
	if part == 1 {
		fmt.Printf("Part 1: %d\n", part1(orderings, updates))
	} else {
		fmt.Printf("Part 2: %d\n", part2(orderings, updates))
	}
}
