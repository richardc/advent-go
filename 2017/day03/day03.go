package day03

import (
	_ "embed"
	"strings"

	"github.com/richardc/advent-go/input"
	"github.com/richardc/advent-go/math"
	"github.com/richardc/advent-go/runner"
	"github.com/richardc/advent-go/slices"
	"golang.org/x/exp/maps"
)

//go:embed "input.txt"
var puzzle string

func init() {
	runner.Register(
		runner.Solution{
			Year:  2017,
			Day:   3,
			Part1: func(any) any { return spiralManhattan(input.MustAtoi(strings.TrimSpace(puzzle))) },
			Part2: func(any) any { return spiralManhattanSums(input.MustAtoi(strings.TrimSpace(puzzle))) },
		},
	)
}

func spiralManhattan(puzzle int) int {
	index := 1
	return spiralManhattanTours(puzzle, func(x, y int) bool {
		//fmt.Printf("%d %d %d,%d", index, puzzle, x, y)
		index++
		return index <= puzzle
	})
}

func spiralManhattanSums(puzzle int) int {
	sums := map[[2]int]int{}
	neighbours := [][]int{
		{-1, -1}, {-1, 0}, {-1, 1},
		{0, -1}, {0, 1},
		{1, -1}, {1, 0}, {1, 1},
	}

	spiralManhattanTours(puzzle, func(x, y int) bool {
		value := 0
		if x == 0 && y == 0 {
			// basecase, we have no neighbours, so just have the value 1
			value = 1
		} else {
			for _, neighbour := range neighbours {
				value += sums[[2]int{x + neighbour[0], y + neighbour[1]}]
			}
		}
		sums[[2]int{x, y}] = value

		return value <= puzzle
	})

	return slices.Max(maps.Values(sums))
}

func spiralManhattanTours(puzzle int, visit func(x, y int) bool) int {
	x, y := 0, 0
	minx, maxx := 0, 0
	miny, maxy := 0, 0
	direction := 0
	for visit(x, y) {
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
