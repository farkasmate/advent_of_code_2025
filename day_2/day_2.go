package aoc2025

import (
	"math"
	"strconv"
	"strings"
)

type Interval struct {
	Start int
	End   int
}

func ParseIntervals(input string) []Interval {
	splits := strings.Split(input, ",")
	intervals := make([]Interval, len(splits))

	for i, s := range splits {
		p := strings.Split(s, "-")
		start, err := strconv.Atoi(p[0])
		if err != nil {
			panic(err)
		}
		end, err := strconv.Atoi(p[1])
		if err != nil {
			panic(err)
		}

		intervals[i] = Interval{start, end}
	}

	return intervals
}

func NextInvalidId(id int) int {
	if id < 10 {
		return 11
	}

	strId := strconv.Itoa(id)
	lenId := len(strId)

	if lenId%2 != 0 {
		return NextInvalidId(int(math.Pow10(lenId)))
	}

	firstHalf := strId[:lenId/2]
	nextId, err := strconv.Atoi(firstHalf + firstHalf)
	if err != nil {
		panic(err)
	}

	if nextId >= id {
		return nextId
	}

	firstHalfInt, err := strconv.Atoi(firstHalf)
	if err != nil {
		panic(err)
	}
	firstHalfInt += 1
	firstHalf = strconv.Itoa(firstHalfInt)

	nextId, err = strconv.Atoi(firstHalf + firstHalf)
	if err != nil {
		panic(err)
	}

	return nextId
}

func GetAllInvalidIds(intervals []Interval) []int {
	invalidIds := make([]int, 0)
	for _, i := range intervals {
		s := i.Start
		for {
			id := NextInvalidId(s)
			if id > i.End {
				break
			}
			invalidIds = append(invalidIds, id)
			s = id + 1
		}
	}

	return invalidIds
}

func Sum(ids []int) int {
	sum := 0
	for _, id := range ids {
		sum += id
	}
	return sum
}
