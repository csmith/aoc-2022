package common

// Tile is a single space within a 2D map, as read by ReadFileAsMap.
type Tile rune

// Map is a 2d map.
type Map [][]Tile

// ReadFileAsMap reads all lines from the given path and returns them as a two-dimensional map.
// If an error occurs, the function will panic.
func ReadFileAsMap(path string) Map {
	var res [][]Tile
	lines := ReadFileAsStrings(path)
	for i := range lines {
		res = append(res, []Tile(lines[i]))
	}
	return res
}

// TileAt returns the tile at the given co-ordinates in the map, wrapping around on both axes.
func (m Map) TileAt(row, col int) Tile {
	line := m[row%len(m)]
	return line[col%len(line)]
}

// SafeTileAt returns the tile at the given row and column, or if those co-ordinates are out of bounds, the fallback.
func (m Map) SafeTileAt(row, col int, fallback Tile) Tile {
	if row < 0 || col < 0 || row >= len(m) || col >= len(m[row]) {
		return fallback
	}
	return m[row][col]
}

var directions = [][]int{
	{-1, -1},
	{-1, 0},
	{-1, 1},
	{0, -1},
	{0, 1},
	{1, -1},
	{1, 0},
	{1, 1},
}

// Neighbours returns the eight neighbours of the given cell, using the given fallback tile if those neighbours
// are out of bounds.
func (m Map) Neighbours(row, col int, fallback Tile) []Tile {
	var res []Tile
	for _, n := range directions {
		res = append(res, m.SafeTileAt(row+n[0], col+n[1], fallback))
	}
	return res
}

// Count counts the number of cells in the map that match the given func.
func (m Map) Count(matcher func(Tile)bool) int {
	res := 0
	for y := range m {
		res += CountTiles(m[y], matcher)
	}
	return res
}

// ProjectUntil projects a line from (row, col) that moves by (dy, dx) tiles each iteration until a tile is found
// matching the given func. Returns the row, column and tile of the match.
func (m Map) ProjectUntil(row, col int, dy, dx int, fallback Tile, matcher func(Tile)bool) (int, int, Tile) {
	y := row + dy
	x := col + dx
	for !matcher(m.SafeTileAt(y, x, fallback)) {
		y += dy
		x += dx
	}
	return y, x, m.SafeTileAt(y, x, fallback)
}

// Starburst projects lines out in the eight cardinal directions from the given cell until the given matcher is
// satisfied.
func (m Map) Starburst(row, col int, fallback Tile, matcher func(Tile)bool) []Tile {
	var res []Tile
	for _, n := range directions {
		_, _, t := m.ProjectUntil(row, col, n[0], n[1], fallback, matcher)
		res = append(res, t)
	}
	return res
}

// CountTiles counts the number of tiles in the slice that match the given matcher.
func CountTiles(tiles []Tile, matcher func(Tile)bool) int {
	res := 0
	for i := range tiles {
		if matcher(tiles[i]) {
			res++
		}
	}
	return res
}
