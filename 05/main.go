package main

import (
	"fmt"

	"github.com/csmith/aoc-2022/common"
)

func main() {
	input := common.ReadFileAsSectionedStrings("05/input.txt")
	p1state := parseState(input[0])
	p2state := make([][]uint8, len(p1state))
	for i := range p1state {
		p2state[i] = append([]uint8{}, p1state[i]...)
	}

	var num, from, to int
	for i := range input[1] {
		fmt.Sscanf(input[1][i], "move %d from %d to %d", &num, &from, &to)
		// Part 1
		selection := p1state[from-1][len(p1state[from-1])-num:]
		common.Reverse(selection)
		p1state[from-1] = p1state[from-1][:len(p1state[from-1])-num]
		p1state[to-1] = append(p1state[to-1], selection...)

		// Part 2
		selection = p2state[from-1][len(p2state[from-1])-num:]
		p2state[from-1] = p2state[from-1][:len(p2state[from-1])-num]
		p2state[to-1] = append(p2state[to-1], selection...)
	}

	for i := range p1state {
		fmt.Printf("%c", p1state[i][len(p1state[i])-1])
	}
	fmt.Println()

	for i := range p2state {
		fmt.Printf("%c", p2state[i][len(p2state[i])-1])
	}
	fmt.Println()
}

func parseState(lines []string) [][]uint8 {
	stacks := make([][]uint8, 1+len(lines[0])/4)
	for i := len(lines) - 2; i >= 0; i-- {
		for j := 1; j < len(lines[i]); j += 4 {
			el := lines[i][j]
			if el != ' ' {
				stacks[(j-1)/4] = append(stacks[(j-1)/4], el)
			}
		}
	}
	return stacks
}
