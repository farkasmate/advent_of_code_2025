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

func Max(slice []int) int {
	if len(slice) == 0 {
		panic("Max of empty slice")
	}
	c := make([]int, len(slice))
	copy(c, slice)
	slices.Sort(c)

	return c[len(c)-1]
}

func IndexOf(slice []int, value int) int {
	for i, v := range slice {
		if v == value {
			return i
		}
	}

	return -1
}

func Joltage(bank []int, num int) int {
	joltage := 0

	start := 0
	end := len(bank) - num + 1
	last := 0

	for i := 0; i < num; i++ {
		r := bank[start:end]
		last = Max(r)
		joltage = 10*joltage + last

		lastIndex := IndexOf(r, last)
		start += lastIndex + 1
		end += 1
	}

	return joltage
}

func GetJoltages(banks [][]int, num int) []int {
	joltages := make([]int, len(banks))
	for i, bank := range banks {
		joltages[i] = Joltage(bank, num)
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
	resultTwelve := 3121910778619

	banks := ParseInput(input)
	t.Log("Parsed banks:", banks)

	joltages := GetJoltages(banks, 2)
	t.Log("Joltages:", joltages)

	sum := Sum(joltages)
	t.Log("Sum of joltages:", sum)

	if sum != result {
		t.Errorf("Expected %d, got %d", result, sum)
	}

	joltagesTwelve := GetJoltages(banks, 12)
	t.Log("Joltages for 12:", joltagesTwelve)

	sumTwelve := Sum(joltagesTwelve)
	t.Log("Sum of joltages for 12:", sumTwelve)

	if sumTwelve != resultTwelve {
		t.Errorf("Expected %d, got %d", resultTwelve, sumTwelve)
	}
}

func TestInput(t *testing.T) {
	input, err := os.ReadFile("input_3")
	if err != nil {
		t.Fatal(err)
	}

	banks := ParseInput(string(input))
	t.Log("Parsed banks:", banks)

	joltages := GetJoltages(banks, 2)
	t.Log("Joltages:", joltages)

	sum := Sum(joltages)
	t.Log("Sum of joltages:", sum)

	joltagesTwelve := GetJoltages(banks, 12)
	t.Log("Joltages for 12:", joltagesTwelve)

	sumTwelve := Sum(joltagesTwelve)
	t.Log("Sum of joltages for 12:", sumTwelve)
}
