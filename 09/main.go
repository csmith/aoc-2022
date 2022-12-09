package main

import (
	"fmt"

	"github.com/csmith/aoc-2022/common"
)

var directions = map[rune][]int{
	'U': {0, -1},
	'D': {0, +1},
	'L': {-1, 0},
	'R': {+1, 0},
}

const knots = 10

func main() {
	var positions = [knots][2]int{}
	var visitedFirst = map[string]bool{}
	var visitedLast = map[string]bool{}
	var dir rune
	var num int
	input := common.ReadFileAsStrings("09/input.txt")
	for i := range input {
		fmt.Sscanf(input[i], "%c %d", &dir, &num)
		for j := 0; j < num; j++ {
			positions[0][0] += directions[dir][0]
			positions[0][1] += directions[dir][1]

			for z := 1; z < knots; z++ {
				positions[z][0], positions[z][1] = chase(positions[z][0], positions[z][1], positions[z-1][0], positions[z-1][1])
			}

			visitedFirst[fmt.Sprintf("%d,%d", positions[1][0], positions[1][1])] = true
			visitedLast[fmt.Sprintf("%d,%d", positions[knots-1][0], positions[knots-1][1])] = true
		}
	}

	println(len(visitedFirst))
	println(len(visitedLast))
}

func chase(tx int, ty int, hx int, hy int) (int, int) {
	if tx == hx && ty == hy {
		// Already on the same cell
		return tx, ty
	}

	dx := hx - tx
	dy := hy - ty
	if dx >= -1 && dx <= 1 && dy >= -1 && dy <= 1 {
		// Adjacent, no need to move
		return tx, ty
	}

	// Otherwise, chase
	return tx + common.Sign(dx), ty + common.Sign(dy)
}
