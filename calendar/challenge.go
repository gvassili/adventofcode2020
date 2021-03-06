package calendar

import (
	"fmt"
	"github.com/gvassili/adventofcode2020/calendar/day01"
	"github.com/gvassili/adventofcode2020/calendar/day02"
	"github.com/gvassili/adventofcode2020/calendar/day03"
	"github.com/gvassili/adventofcode2020/calendar/day04"
	"github.com/gvassili/adventofcode2020/calendar/day05"
	"github.com/gvassili/adventofcode2020/calendar/day06"
	"github.com/gvassili/adventofcode2020/calendar/day07"
	"github.com/gvassili/adventofcode2020/calendar/day08"
	"github.com/gvassili/adventofcode2020/calendar/day09"
	"github.com/gvassili/adventofcode2020/calendar/day10"
	"github.com/gvassili/adventofcode2020/calendar/day18"
	"io"
	"sort"
)

var challenges = map[int]func() Challenge{
	1:  func() Challenge { return new(day01.Challenge) },
	2:  func() Challenge { return new(day02.Challenge) },
	3:  func() Challenge { return new(day03.Challenge) },
	4:  func() Challenge { return new(day04.Challenge) },
	5:  func() Challenge { return new(day05.Challenge) },
	6:  func() Challenge { return new(day06.Challenge) },
	7:  func() Challenge { return new(day07.Challenge) },
	8:  func() Challenge { return new(day08.Challenge) },
	9:  func() Challenge { return new(day09.Challenge) },
	10: func() Challenge { return new(day10.Challenge) },
	18: func() Challenge { return new(day18.Challenge) },
}

func Load(day int) (Challenge, error) {
	loader, ok := challenges[day]
	if !ok {
		return nil, fmt.Errorf("could not find challenge %d", day)
	}
	return loader(), nil
}

type Challenge interface {
	Day() int
	Prepare(r io.Reader) error
	Part1() (string, error)
	Part2() (string, error)
}

func LoadAllChallenges() []Challenge {
	challengeNames := make([]int, 0, len(challenges))
	for name := range challenges {
		challengeNames = append(challengeNames, name)
	}
	sort.Ints(challengeNames)
	result := make([]Challenge, 0, len(challenges))
	for _, day := range challengeNames {
		result = append(result, challenges[day]())
	}
	return result
}
