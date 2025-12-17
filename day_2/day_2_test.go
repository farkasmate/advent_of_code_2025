package aoc2025

import (
	"os"
	"slices"
	"strings"
	"testing"
)

func TestExample(t *testing.T) {
	input := "11-22,95-115,998-1012,1188511880-1188511890,222220-222224,1698522-1698528,446443-446449,38593856-38593862,565653-565659,824824821-824824827,2121212118-2121212124"
	result := 1227775554

	intervals := ParseIntervals(input)
	t.Log("Intervals:", intervals)

	invalidIds := GetAllInvalidIds(intervals)
	t.Log("Invalid IDs", invalidIds)

	sum := Sum(invalidIds)
	t.Log("Sum of invalid IDs:", sum)

	if sum != result {
		t.Errorf("Expected %d, got %d", result, sum)
	}
}

func TestInput(t *testing.T) {
	b, err := os.ReadFile("input_2")
	if err != nil {
		panic(err)
	}

	input := strings.TrimSpace(string(b))
	intervals := ParseIntervals(input)
	t.Log("Intervals:", intervals)

	invalidIds := GetAllInvalidIds(intervals)
	slices.Sort(invalidIds)
	invalidIds = slices.Compact(invalidIds)
	t.Log("Invalid IDs", invalidIds)

	sum := Sum(invalidIds)
	t.Log("Sum of invalid IDs:", sum)
}
