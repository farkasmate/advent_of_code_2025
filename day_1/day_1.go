package aoc2025

import (
	"strconv"
	"strings"
)

func LeftAtZero(input string) int {
	dial := 50
	at_zero := 0

	for line := range strings.Lines(input) {
		line = strings.TrimSpace(line)
		if line == "" {
			continue
		}

		val, err := strconv.Atoi(line[1:])
		if err != nil {
			panic(err)
		}

		if line[0] == 'L' {
			val = -val
		}

		dial += val
		dial %= 100
		for ;dial < 0; dial += 100 {}

		if dial == 0 {
			at_zero += 1
		}
	}

	return at_zero
}
