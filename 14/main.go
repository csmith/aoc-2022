package main

import (
	"strings"

	"github.com/csmith/aoc-2022/common"
)

const (
	// Arbitrary bounds ahoy!
	xMin = 300
	xMax = 700
	yMin = 0
	yMax = 200

	dropX = 500
	dropY = 0

	empty common.Tile = '.'
	shelf common.Tile = '#'
	sand  common.Tile = 'o'
	drop  common.Tile = '+'
	oob   common.Tile = '@'
)

func main() {
	cave := common.NewMap(xMax-xMin, yMax-yMin, empty)

	cave[dropY-yMin][dropX-xMin] = drop
	floorHeight := 2 + plotShelves(common.ReadFileAsStrings("14/input.txt"), cave)

	rounds := dropSand(cave, dropX-xMin, dropY-yMin)
	println(rounds)

	for x := xMin; x < xMax; x++ {
		cave[floorHeight-yMin][x-xMin] = shelf
	}

	println(rounds + dropSand(cave, dropX-xMin, dropY-yMin))
}

func dropSand(cave common.Map, x, y int) int {
	bY, _, t := cave.ProjectUntil(y, x, 1, 0, oob, func(tile common.Tile) bool {
		return tile == sand || tile == shelf || tile == oob
	})

	if t == oob {
		// Bye bye, cruel world!
		return 0
	}

	// Pour sand down the diagonals for as long as we can
	rounds := 0
	for _, xOffset := range []int{-1, 1} {
		for {
			if t := cave.SafeTileAt(bY, x+xOffset, oob); t == empty {
				if r := dropSand(cave, x+xOffset, bY); r > 0 {
					rounds += r
				} else {
					return rounds
				}
			} else if t == oob {
				// Weeeeeeee
				return rounds
			} else {
				// Hit a shelf or otherwise blocked, carry on.
				break
			}
		}
	}

	// Settle on top
	cave[bY-1][x] = sand
	rounds++

	// Optimisation: see if we can place sand above us
	for nY := bY - 2; nY >= y; nY-- {
		if r := dropSand(cave, x, nY); r > 0 {
			rounds += r
		} else {
			return rounds
		}
	}

	return rounds
}

func plotShelves(shelves []string, cave common.Map) int {
	maxY := 0
	for i := range shelves {
		parts := strings.Split(shelves[i], " -> ")
		lastX, lastY := -1, -1
		for j := range parts {
			coords := strings.Split(parts[j], ",")
			x, y := common.MustAtoi(coords[0])-xMin, common.MustAtoi(coords[1])-yMin
			if y > maxY {
				maxY = y
			}

			if x == lastX {
				// Draw a vertical line
				for z := common.Min(lastY, y); z <= common.Max(lastY, y); z++ {
					cave[z][x] = shelf
				}
			} else if y == lastY {
				// Draw a horizontal line
				for z := common.Min(lastX, x); z <= common.Max(lastX, x); z++ {
					cave[y][z] = shelf
				}
			}

			lastX, lastY = x, y
		}
	}

	return maxY
}
