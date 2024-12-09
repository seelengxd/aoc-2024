package main

import (
	_ "embed"
	"flag"
	"fmt"
	"slices"
	"strconv"
)

//go:embed input.txt
var input string

type Block struct {
	FileID int
}

func (b Block) IsEmpty() bool {
	return b.FileID == -1
}

func (b Block) String() string {
	if b.IsEmpty() {
		return "."
	}
	return strconv.Itoa(b.FileID)
}

type Group struct {
	Block Block
	Count int
}

func (g Group) IsEmpty() bool {
	return g.Block.IsEmpty()
}

var EMPTY_BLOCK = Block{-1}

func parseBlocks(input string) []Block {
	blocks := []Block{}

	currFileID := 0

	for i, n := range input {
		num := int(n - '0')
		var block Block
		if i%2 == 1 {
			block = EMPTY_BLOCK
		} else {
			block = Block{currFileID}
			currFileID++
		}
		for j := 0; j < num; j++ {
			blocks = append(blocks, block)
		}
	}
	return blocks
}

func parseGroups(input string) []Group {
	groups := []Group{}
	currFileID := 0

	for i, n := range input {
		num := int(n - '0')
		var block Block
		if i%2 == 1 {
			block = EMPTY_BLOCK
		} else {
			block = Block{currFileID}
			currFileID++
		}
		groups = append(groups, Group{block, num})

	}
	return groups
}

func part1(input string) (answer int) {
	blocks := parseBlocks(input)
	leftPointer, rightPointer := 0, len(blocks)-1
	for leftPointer < rightPointer {
		if !blocks[leftPointer].IsEmpty() {
			leftPointer++
			continue
		}
		if blocks[rightPointer].IsEmpty() {
			rightPointer--
			continue
		}
		blocks[leftPointer], blocks[rightPointer] = blocks[rightPointer], blocks[leftPointer]
		leftPointer++
		rightPointer--
	}
	for i, block := range blocks {
		if block.IsEmpty() {
			break
		}
		answer += i * block.FileID
	}
	return answer
}

func part2(input string) (answer int) {
	groups := parseGroups(input)

	for i := len(groups) - 1; i >= 0; i-- {
		if groups[i].IsEmpty() {
			continue
		}
		for j := 0; j < i; j++ {
			if !groups[j].IsEmpty() || groups[j].Count < groups[i].Count {
				continue
			}
			if groups[i].Count == groups[j].Count {
				groups[i], groups[j] = groups[j], groups[i]
				break
			}

			spaceGroup := groups[j]
			fileGroup := groups[i]
			spaceCount := spaceGroup.Count - fileGroup.Count

			groups[j] = Group{fileGroup.Block, fileGroup.Count}
			groups[i] = Group{EMPTY_BLOCK, fileGroup.Count}
			groups = slices.Concat(groups[:j+1], []Group{{EMPTY_BLOCK, spaceCount}}, groups[j+1:])
			i++
			break
		}
	}
	currIndex := 0
	for _, group := range groups {
		block := group.Block
		multiplier := block.FileID
		if block.IsEmpty() {
			multiplier = 0
		}
		for i := 0; i < group.Count; i++ {
			answer += currIndex * multiplier
			currIndex++
		}
	}

	return answer
}

func main() {
	var part int
	flag.IntVar(&part, "part", 1, "part 1 or 2")
	flag.Parse()

	if part == 1 {
		fmt.Printf("Part 1: %d\n", part1(input))
	} else {
		fmt.Printf("Part 2: %d\n", part2(input))
	}
}
