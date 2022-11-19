package day18

import (
	_ "embed"
	"strings"

	"github.com/richardc/advent-go/input"
	"github.com/richardc/advent-go/maths"
	"github.com/richardc/advent-go/runner"
	"github.com/richardc/advent-go/slices"
)

//go:embed input.txt
var puzzle string

func init() {
	runner.Register(runner.Solution{
		Day:   18,
		Input: func() any { return newLife(puzzle) },
		Part1: func(i any) any { return iterateLit(i.(life), 100) },
		Part2: func(i any) any { return iterateLitCorners(i.(life), 100) },
	})
}

type life struct {
	x, y  int
	cells [][]bool
}

func newLife(s string) life {
	lines := input.Lines(s)

	return life{
		x: len(lines[0]),
		y: len(lines),
		cells: slices.Map(lines, func(s string) []bool {
			return slices.Map([]byte(s), func(b byte) bool { return b == '#' })
		}),
	}
}

func (l life) lit() int {
	count := 0
	for _, row := range l.cells {
		for _, cell := range row {
			if cell {
				count++
			}
		}
	}
	return count
}

func (l life) String() string {
	return strings.Join(slices.Map(l.cells, func(row []bool) string {
		return string(slices.Map(row, func(cell bool) rune {
			if cell {
				return '#'
			} else {
				return '.'
			}
		}))
	}), "\n")
}

func (l *life) step() {
	next := [][]bool{}
	for y, row := range l.cells {
		next = append(next, make([]bool, l.x))
		for x, cell := range row {
			nextcell := false
			switch l.neighbours(x, y) {
			case 2:
				nextcell = cell
			case 3:
				nextcell = true
			}
			next[y][x] = nextcell
		}
	}
	l.cells = next
}

func (l life) neighbours(cx, cy int) int {
	count := 0
	for x := maths.Max(cx-1, 0); x < maths.Min(cx+2, l.x); x++ {
		for y := maths.Max(cy-1, 0); y < maths.Min(cy+2, l.y); y++ {
			if l.cells[y][x] {
				count++
			}
		}
	}

	if l.cells[cy][cx] {
		count--
	}

	return count
}

func (l *life) turnOnCorners() {
	l.cells[0][0] = true
	l.cells[l.y-1][0] = true
	l.cells[0][l.x-1] = true
	l.cells[l.y-1][l.x-1] = true
}

func iterateLit(game life, iter int) int {
	for i := 0; i < iter; i++ {
		game.step()
	}
	return game.lit()
}

func iterateLitCorners(game life, iter int) int {
	game.turnOnCorners()
	for i := 0; i < iter; i++ {
		game.step()
		game.turnOnCorners()
	}
	return game.lit()
}
