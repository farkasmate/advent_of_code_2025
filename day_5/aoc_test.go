package aoc2025

import (
	"os"
	"strconv"
	"strings"
	"testing"
)

type Range struct {
	Min int
	Max int
}

func (r Range) String() string {
	return "(" + strconv.Itoa(r.Min) + ".." + strconv.Itoa(r.Max) + ")"
}

func (r Range) Contains(n int) bool {
	return n >= r.Min && n <= r.Max
}

func ParseInput(input string) ([]Range, []int) {
	parseRanges := true
	ranges := make([]Range, 0)
	ids := make([]int, 0)

	for _, line := range strings.Split(strings.TrimSpace(input), "\n") {
		if line == "" {
			parseRanges = false
			continue
		}

		if parseRanges {
			s := strings.Split(line, "-")
			min, err := strconv.Atoi(s[0])
			if err != nil {
				panic(err)
			}
			max, err := strconv.Atoi(s[1])
			if err != nil {
				panic(err)
			}
			ranges = append(ranges, Range{Min: min, Max: max})
		} else {
			n, err := strconv.Atoi(line)
			if err != nil {
				panic(err)
			}
			ids = append(ids, n)
		}
	}

	return ranges, ids
}

func GetFreshIngredients(ranges []Range, ids []int) []int {
	fresh := make([]int, 0)

	for _, id := range ids {
		for _, r := range ranges {
			if r.Contains(id) {
				fresh = append(fresh, id)
				break
			}
		}
	}

	return fresh
}

func TestExample(t *testing.T) {
	input := `
3-5
10-14
16-20
12-18

1
5
8
11
17
32
`
	result := 3

	ranges, ids := ParseInput(input)
	t.Log("Ranges:", ranges)
	t.Log("IDs:", ids)

	fresh := GetFreshIngredients(ranges, ids)
	t.Log("Fresh Ingredients:", fresh)
	t.Log("Number of Fresh Ingredients:", len(fresh))

	if len(fresh) != result {
		t.Errorf("Expected %d fresh ingredients, got %d", result, len(fresh))
	}
}

func TestInput(t *testing.T) {
	input, err := os.ReadFile("input")
	if err != nil {
		t.Fatal(err)
	}

	ranges, ids := ParseInput(string(input))
	t.Log("Ranges:", ranges)
	t.Log("IDs:", ids)

	fresh := GetFreshIngredients(ranges, ids)
	t.Log("Fresh Ingredients:", fresh)
	t.Log("Number of Fresh Ingredients:", len(fresh))
}
