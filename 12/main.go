package main

import (
	"github.com/csmith/aoc-2022/common"
)

const (
	start common.Tile = 'S'
	end   common.Tile = 'E'
	oob               = '#'
)

var directions = [][]int{
	{0, -1},
	{0, +1},
	{-1, 0},
	{+1, 0},
}

type step struct {
	distance int
	row      int
	col      int
	height   int
}

func main() {
	grid := common.ReadFileAsMap("12/input.txt")

	part1 := 0
	part2 := 0

	startingRow, startingCol := grid.First(func(tile common.Tile) bool {
		return tile == end
	})

	visited := make([][]bool, len(grid))
	for i := range grid {
		visited[i] = make([]bool, len(grid[i]))
	}
	visited[startingRow][startingCol] = true

	steps := make(chan step, 100)
	steps <- step{
		distance: 0,
		row:      startingRow,
		col:      startingCol,
		height:   height(end),
	}

	for s := range steps {
		if part1 > 0 && part2 > 0 {
			break
		}

		for i := range directions {
			newRow, newCol := s.row+directions[i][0], s.col+directions[i][1]
			if newRow < 0 || newCol < 0 || newRow >= len(grid) || newCol >= len(grid[0]) || visited[newRow][newCol] {
				continue
			}

			t := grid.TileAt(newRow, newCol)
			h := height(t)

			// We're working backwards, so you can go down a single step or up any amount...
			if diff := h - s.height; diff >= -1 {
				visited[newRow][newCol] = true

				if t == start && part1 == 0 {
					part1 = s.distance + 1
				}
				if t == 'a' && part2 == 0 {
					part2 = s.distance + 1
				}

				steps <- step{
					distance: s.distance + 1,
					row:      newRow,
					col:      newCol,
					height:   h,
				}
			}
		}
	}

	println(part1)
	println(part2)
}

func height(tile common.Tile) int {
	if tile == start {
		return 'a'
	} else if tile == end {
		return 'z'
	} else {
		return int(tile)
	}
}
