package main

import (
	"strings"

	"github.com/csmith/aoc-2022/common"
)

const (
	sprite   = "▮"
	noSprite = " "
)

func main() {
	input := common.ReadFileAsStrings("10/input.txt")
	ascii := common.NewAsciiOutput(8)
	cycle := 0
	x := 1
	part1 := 0

	tick := func() {
		col := cycle % 40
		ascii.Write(x >= col-1 && x <= col+1)

		cycle++
		if (cycle-20)%40 == 0 {
			part1 += cycle * x
		}
	}

	for i := range input {
		tick()
		if input[i] != "noop" {
			tick()
			x += common.MustAtoi(strings.TrimPrefix(input[i], "addx "))
		}
	}

	println(part1)
	println(ascii.Read())
}
