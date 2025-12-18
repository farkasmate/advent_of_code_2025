package aoc2025

import (
	"os"
	"slices"
	"strings"
	"testing"
)

func ParseInput(input string) [][]int {
	banks := make([][]int, 0)
	for _, b := range strings.Split(strings.TrimSpace(input), "\n") {
		banks = append(banks, ParseBank(b))
	}

	return banks
}

func ParseBank(input string) []int {
	batteries := make([]int, len(input))
	for i, r := range input {
		batteries[i] = int(r - '0')
	}

	return batteries
}

func Joltage(bank []int) int {
	c := make([]int, len(bank)-1)
	copy(c, bank[0:len(bank)-1])

	slices.Sort(c)
	first := c[len(c)-1]
	second := 0

	for i, v := range bank {
		if v < first {
			continue
		}

		c2 := make([]int, len(bank)-1-i)
		copy(c2, bank[i+1:len(bank)])
		slices.Sort(c2)
		second = c2[len(c2)-1]

		break
	}

	return 10*first + second
}

func GetJoltages(banks [][]int) []int {
	joltages := make([]int, len(banks))
	for i, bank := range banks {
		joltages[i] = Joltage(bank)
	}

	return joltages
}

func Sum(joltages []int) int {
	sum := 0
	for _, j := range joltages {
		sum += j
	}

	return sum
}

func TestExample(t *testing.T) {
	input := `
987654321111111
811111111111119
234234234234278
818181911112111
`
	result := 357

	banks := ParseInput(input)
	t.Log("Parsed banks:", banks)

	joltages := GetJoltages(banks)
	t.Log("Joltages:", joltages)

	sum := Sum(joltages)
	t.Log("Sum of joltages:", sum)

	if sum != result {
		t.Errorf("Expected %d, got %d", result, sum)
	}
}

func TestInput(t *testing.T) {
	input, err := os.ReadFile("input_3")
	if err != nil {
		t.Fatal(err)
	}

	banks := ParseInput(string(input))
	t.Log("Parsed banks:", banks)

	joltages := GetJoltages(banks)
	t.Log("Joltages:", joltages)

	sum := Sum(joltages)
	t.Log("Sum of joltages:", sum)
}
