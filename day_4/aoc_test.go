package aoc2025

import (
	"os"
	"strings"
	"testing"
)

type Coord struct {
	X int
	Y int
}

type Grid [][]int

func (grid Grid) String() string {
	var sb strings.Builder
	sb.WriteRune('\n')

	for _, row := range [][]int(grid) {
		for _, v := range row {
			if v == 1 {
				sb.WriteRune('@')
			} else {
				sb.WriteRune('.')
			}
		}
		sb.WriteRune('\n')
	}

	return sb.String()
}

func (grid Grid) Get(x, y int) int {
	if y < 0 || x < 0 || y >= len(grid) || x >= len(grid[0]) {
		return 0
	}

	return grid[y][x]
}

func Accessible(grid Grid, x, y int) (bool, int) {
	if y < 0 || x < 0 || y > len(grid) || x > len(grid[0]) || grid[y][x] == 0 {
		return false, 0
	}

	neighbours := 0

	neighbours += grid.Get(x-1, y-1)
	neighbours += grid.Get(x, y-1)
	neighbours += grid.Get(x+1, y-1)
	neighbours += grid.Get(x-1, y)
	neighbours += grid.Get(x+1, y)
	neighbours += grid.Get(x-1, y+1)
	neighbours += grid.Get(x, y+1)
	neighbours += grid.Get(x+1, y+1)

	return neighbours < 4, neighbours
}

func CountAccessible(grid Grid) (int, []Coord) {
	count := 0
	coords := make([]Coord, 0)

	for y := 0; y < len(grid); y++ {
		for x := 0; x < len(grid[0]); x++ {
			accessible, _ := Accessible(grid, x, y)
			if accessible {
				count++
				coords = append(coords, Coord{X: x, Y: y})
			}
		}
	}

	return count, coords
}

func (grid *Grid) RemoveCoords(original Grid, coords []Coord) {
	for _, coord := range coords {
		(*grid)[coord.Y][coord.X] = 0
	}
}

func ParseInput(input string) Grid {
	tiles := make(Grid, 0)
	for _, line := range strings.Split(strings.TrimSpace(input), "\n") {
		tiles = append(tiles, ParseLine(line))
	}

	return tiles
}

func ParseLine(line string) []int {
	row := make([]int, len(line))
	for i, v := range line {
		if v == '@' {
			row[i] = 1
		}
	}

	return row
}

func TestExample(t *testing.T) {
	input := `
..@@.@@@@.
@@@.@.@.@@
@@@@@.@.@@
@.@@@@..@.
@@.@@@@.@@
.@@@@@@@.@
.@.@.@.@@@
@.@@@.@@@@
.@@@@@@@@.
@.@.@@@.@.
`
	result := 13
	resultWithRemove := 43

	grid := ParseInput(input)
	t.Log(grid)

	accessibleCount, accessibleCoords := CountAccessible(grid)
	t.Logf("Accessible tiles: %d", accessibleCount)

	if accessibleCount != result {
		t.Errorf("Expected %d accessible tiles, got %d", result, accessibleCount)
	}

	count := accessibleCount
	for count > 0 {
		grid.RemoveCoords(grid, accessibleCoords)
		//t.Logf("Grid after removal: %s", grid)
		count, accessibleCoords = CountAccessible(grid)
		accessibleCount += count
	}
	t.Logf("Total accessible tiles after removals: %d", accessibleCount)

	if accessibleCount != resultWithRemove {
		t.Errorf("Expected %d total accessible tiles after removals, got %d", resultWithRemove, accessibleCount)
	}
}

func TestInput(t *testing.T) {
	input, err := os.ReadFile("input")
	if err != nil {
		t.Fatal(err)
	}

	grid := ParseInput(string(input))
	t.Log(grid)

	accessibleCount, accessibleCoords := CountAccessible(grid)
	t.Logf("Accessible tiles: %d", accessibleCount)

	count := accessibleCount
	for count > 0 {
		grid.RemoveCoords(grid, accessibleCoords)
		count, accessibleCoords = CountAccessible(grid)
		accessibleCount += count
	}
	t.Logf("Total accessible tiles after removals: %d", accessibleCount)
}
