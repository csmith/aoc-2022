package main

import (
	"math"

	"github.com/csmith/aoc-2022/common"
)

const ones = ^uint64(0)

func main() {
	bags := common.ReadFileAsStrings("03/input.txt")
	p1, p2, group := 0, 0, ones
	for i := range bags {
		group &= mask(bags[i])
		overlap := mask(bags[i][:len(bags[i])/2]) & mask(bags[i][len(bags[i])/2:])
		p1 += int(math.Log2(float64(overlap)))

		if i%3 == 2 {
			p2 += int(math.Log2(float64(group)))
			group = ones
		}
	}
	println(p1)
	println(p2)
}

func mask(input string) uint64 {
	var res uint64
	for i := range input {
		var m = 0
		if input[i] <= 'Z' {
			m = int(input[i] - 'A' + 27)
		} else {
			m = int(input[i] - 'a' + 1)
		}
		res |= 1 << m
	}
	return res
}
