package day03

import (
	_ "embed"
	"strings"

	"github.com/richardc/advent-go/input"
	"github.com/richardc/advent-go/math"
	"github.com/richardc/advent-go/runner"
)

//go:embed "input.txt"
var puzzle string

func init() {
	runner.Register(
		runner.Solution{
			Year:  2017,
			Day:   3,
			Part1: func(any) any { return spiralManhattan(input.MustAtoi(strings.TrimSpace(puzzle))) },
		},
	)
}

func spiralManhattan(puzzle int) int {
	x, y := 0, 0
	minx, maxx := 0, 0
	miny, maxy := 0, 0
	direction := 0
	for n := 1; n < puzzle; n++ {
		//fmt.Printf("%d: %d,%d %d\n", n, x, y, math.Abs(x)+math.Abs(y))
		switch direction {
		case 0:
			x++
			if x > maxx {
				maxx++
				direction++
			}
		case 1:
			y++
			if y > maxy {
				maxy++
				direction++
			}
		case 2:
			x--
			if x < minx {
				minx--
				direction++
			}
		case 3:
			y--
			if y < miny {
				miny--
				direction = 0
			}
		}
	}
	return math.Abs(x) + math.Abs(y)
}
