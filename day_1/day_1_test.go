package aoc2025

import (
	"os"
	"testing"
)

func TestExample(t *testing.T) {
	input := `
L68
L30
R48
L5
R60
L55
L1
L99
R14
L82
`

	n := LeftAtZero(input)
	if n != 3 {
		t.Errorf("Dial left at zero %d time(s)", n)
	}
}

func TestA(t *testing.T) {
	b, err := os.ReadFile("input_1a")
	if err != nil {
		panic(err)
	}

	input := string(b)

	n := LeftAtZero(input)
	t.Log(n)
}
