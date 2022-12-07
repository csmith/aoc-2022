package main

import "github.com/csmith/aoc-2022/common"

func main() {
	input := common.ReadFileAsStrings("06/input.txt")[0]

	p1, p2 := 0, 0
	for i := range input {
		if p1 == 0 && unique(input[i:i+4]) {
			p1 = i + 4
		}
		// Minor optimisation - we can never start 14 unique chars if we haven't started 4 unique chars
		if p1 > 0 && p2 == 0 && unique(input[i:i+14]) {
			p2 = i + 14
			break
		}
	}

	println(p1)
	println(p2)
}

func unique(input string) bool {
	for i := 0; i < len(input)-1; i++ {
		for j := i + 1; j < len(input); j++ {
			if input[i] == input[j] {
				return false
			}
		}
	}
	return true
}
