package main

import (
	"strings"

	"github.com/csmith/aoc-2022/common"
)

func main() {
	pairs := common.ReadFileAsStrings("04/input.txt")
	overlap, contain := 0, 0
	for i := range pairs {
		parts := common.AtoiSlice(strings.Split(strings.Replace(pairs[i], ",", "-", 1), "-"))
		if parts[0] >= parts[2] && parts[0] <= parts[3] || parts[2] >= parts[0] && parts[2] <= parts[1] {
			overlap++
		}
		if parts[0] >= parts[2] && parts[1] <= parts[3] || parts[2] >= parts[0] && parts[3] <= parts[1] {
			contain++
		}
	}
	println(contain)
	println(overlap)
}
