package day14

import (
	_ "embed"
	"fmt"
	"strconv"
	"strings"

	"github.com/richardc/advent-go/input"
	"github.com/richardc/advent-go/runner"
	"github.com/richardc/advent-go/slices"
)

//go:embed input.txt
var puzzle string

func init() {
	runner.Register(runner.Solution{
		Day:   14,
		Input: func() any { return slices.Map(input.Lines(puzzle), newReindeer) },
		Part1: func(i any) any { return winningDistanceAt(i.([]reindeer), 2503) },
		Part2: func(i any) any { return winningScoreAt(i.([]reindeer), 2503) },
	})
}

type reindeer struct {
	name  string
	speed int
	burst int
	rest  int
}

func atoi(s string) int {
	i, _ := strconv.Atoi(s)
	return i
}

func newReindeer(s string) reindeer {
	// "Comet can fly 14 km/s for 10 seconds, but then must rest for 127 seconds."
	toks := strings.Fields(s)
	return reindeer{
		name:  toks[0],
		speed: atoi(toks[3]),
		burst: atoi(toks[6]),
		rest:  atoi(toks[13]),
	}
}

type state struct {
	deer     reindeer
	resting  bool
	timer    int
	position int
	score    int
}

func (s *state) Tick() {
	if s.timer == 0 {
		if s.resting { // Now time to burst
			s.timer = s.deer.burst
		} else {
			s.timer = s.deer.rest
		}
		s.resting = !s.resting
	}
	s.timer--

	if !s.resting {
		s.position += s.deer.speed
	}
}

func (s *state) String() string { return fmt.Sprintf("%v", *s) }

func winningDistanceAt(r []reindeer, time int) int {
	deer := slices.Map(r, func(r reindeer) *state {
		// Start with a zero timer, but rested and ready to burst
		return &state{
			deer:    r,
			resting: true,
		}
	})

	for now := 0; now < time; now++ {
		for _, d := range deer {
			d.Tick()
		}
		// fmt.Printf("now %d:  %v\n", now, deer)
	}

	return slices.Max(slices.Map(deer, func(d *state) int { return d.position }))
}

func winningScoreAt(r []reindeer, time int) int {
	deer := slices.Map(r, func(r reindeer) *state {
		// Start with a zero timer, but rested and ready to burst
		return &state{
			deer:    r,
			resting: true,
		}
	})

	for now := 0; now < time; now++ {
		for _, d := range deer {
			d.Tick()
		}
		// fmt.Printf("now %d:  %v\n", now, deer)
		front := slices.Max(slices.Map(deer, func(d *state) int { return d.position }))

		for _, d := range deer {
			if d.position == front {
				d.score++
			}
		}
	}

	return slices.Max(slices.Map(deer, func(d *state) int { return d.score }))
}
