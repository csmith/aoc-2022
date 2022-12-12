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

	println(solve(grid, false, start, func(tile common.Tile) bool {
		return tile == end
	}))

	println(solve(grid, true, end, func(tile common.Tile) bool {
		return tile == start || tile == 'a'
	}))
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

func solve(grid common.Map, invert bool, from common.Tile, to func(common.Tile) bool) int {
	startingRow, startingCol := grid.First(func(tile common.Tile) bool {
		return tile == from
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
		height:   height(from),
	}

	for s := range steps {
		for i := range directions {
			newRow, newCol := s.row+directions[i][0], s.col+directions[i][1]
			if newRow < 0 || newCol < 0 || newRow >= len(grid) || newCol >= len(grid[0]) || visited[newRow][newCol] {
				continue
			}

			t := grid.TileAt(newRow, newCol)
			h := height(t)

			if diff := h - s.height; (invert && diff >= -1) || (!invert && diff <= 1) {
				visited[newRow][newCol] = true

				if to(t) {
					return s.distance + 1
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

	return -1
}
