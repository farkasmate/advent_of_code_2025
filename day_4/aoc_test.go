package aoc2025

import (
	"os"
	"strings"
	"testing"
)

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

func CountAccessible(grid Grid) int {
	count := 0
	for y := 0; y < len(grid); y++ {
		for x := 0; x < len(grid[0]); x++ {
			accessible, _ := Accessible(grid, x, y)
			if accessible {
				count++
			}
		}
	}

	return count
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

	grid := ParseInput(input)
	t.Log(grid)

	accessibleCount := CountAccessible(grid)
	t.Logf("Accessible tiles: %d", accessibleCount)

	if accessibleCount != result {
		t.Errorf("Expected %d accessible tiles, got %d", result, accessibleCount)
	}
}

func TestInput(t *testing.T) {
	input, err := os.ReadFile("input")
	if err != nil {
		t.Fatal(err)
	}

	grid := ParseInput(string(input))
	t.Log(grid)

	accessibleCount := CountAccessible(grid)
	t.Logf("Accessible tiles: %d", accessibleCount)
}
