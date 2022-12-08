package main

import (
	"github.com/csmith/aoc-2022/common"
)

// oob is a placeholder tile to use when we go outside the range of our map. It is guaranteed to be "higher" than
// any actual tree.
const oob common.Tile = ':'

var cardinals = [][]int{
	{-1, 0},
	{0, -1},
	{0, 1},
	{1, 0},
}

func main() {
	m := common.ReadFileAsMap("08/input.txt")

	part1 := int64(0)
	part2 := int64(0)
	for y := range m {
		for x := range m[y] {
			outside := false
			visible := int64(1)

			height := m[y][x]
			blocks := func(t common.Tile) bool { return t >= height }

			for c := range cardinals {
				// Trace a path in each cardinal direction until we hit a tree or escape.
				// For part 2 it's slightly fiddly as a higher tree is included in the count, but we use imaginary
				// trees to handle going OOB so they have to be removed.
				dy, dx, t := m.ProjectUntil(y, x, cardinals[c][0], cardinals[c][1], oob, blocks)
				trees := common.Abs(int64(dx-x)) + common.Abs(int64(dy-y))
				if t == oob {
					trees--
					outside = true
				}
				visible *= trees
			}

			if outside {
				part1++
			}
			if visible > part2 {
				part2 = visible
			}
		}
	}

	println(part1)
	println(part2)
}
