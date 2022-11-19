package day03

import (
	_ "embed"
	"strings"

	"github.com/richardc/advent-go/input"
	"github.com/richardc/advent-go/maths"
	"github.com/richardc/advent-go/runner"
	"github.com/richardc/advent-go/slices"
	"golang.org/x/exp/maps"
)

//go:embed "input.txt"
var puzzle string

func init() {
	runner.Register(
		runner.Solution{
			Year:  2019,
			Day:   3,
			Input: func() any { closest, smallest := solve(puzzle); return []int{closest, smallest} },
			Part1: func(i any) any { return i.([]int)[0] },
			Part2: func(i any) any { return i.([]int)[1] },
		},
	)
}

type Link struct {
	direction byte
	distance  int
}

func newLink(s string) Link {
	return Link{
		direction: s[0],
		distance:  input.MustAtoi(s[1:]),
	}
}

type Point [2]int

func solve(puzzle string) (int, int) {
	wires := slices.Map(input.Lines(puzzle), func(s string) []Link { return slices.Map(strings.Split(s, ","), newLink) })
	points := map[Point]int{}
	steps := [2]map[Point]int{}
	for i, wire := range wires {
		bit := i + 1
		x, y := 0, 0
		step := 0
		steps[i] = map[Point]int{}
		for _, link := range wire {
			for j := 0; j < link.distance; j++ {
				switch link.direction {
				case 'D':
					y--
				case 'U':
					y++
				case 'L':
					x--
				case 'R':
					x++
				}
				step++
				if _, ok := steps[i][Point{x, y}]; !ok {
					steps[i][Point{x, y}] = step
				}
				points[Point{x, y}] |= bit
			}
		}
	}

	crossed := slices.Filter(maps.Keys(points), func(p Point) bool { return points[p] == 3 })
	closest := slices.Min(slices.Map(crossed, func(p Point) int { return maths.Abs(p[0]) + maths.Abs(p[1]) }))
	smallest := slices.Min(slices.Map(crossed, func(p Point) int { return steps[0][p] + steps[1][p] }))
	return closest, smallest
}
