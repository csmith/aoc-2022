package main

import "github.com/csmith/aoc-2022/common"

var part1 = map[string]int{
	"A X": 3 + 1,
	"A Y": 6 + 2,
	"A Z": 0 + 3,
	"B X": 0 + 1,
	"B Y": 3 + 2,
	"B Z": 6 + 3,
	"C X": 6 + 1,
	"C Y": 0 + 2,
	"C Z": 3 + 3,
}

var part2 = map[string]int{
	"A X": 0 + 3,
	"A Y": 3 + 1,
	"A Z": 6 + 2,
	"B X": 0 + 1,
	"B Y": 3 + 2,
	"B Z": 6 + 3,
	"C X": 0 + 2,
	"C Y": 3 + 3,
	"C Z": 6 + 1,
}

func main() {
	matches := common.ReadFileAsStrings("02/input.txt")
	p1, p2 := 0, 0
	for i := range matches {
		p1 += part1[matches[i]]
		p2 += part2[matches[i]]
	}
	println(p1)
	println(p2)
}
