package day22

import (
	_ "embed"

	"github.com/richardc/advent-go/input"
	"github.com/richardc/advent-go/runner"
)

//go:embed "input.txt"
var puzzle string

func init() {
	runner.Register(
		runner.Solution{
			Year:  2017,
			Day:   22,
			Part1: func(any) any { return burstsAfter(puzzle, 10_000) },
		},
	)
}

type Direction int

const (
	Up Direction = iota
	Right
	Down
	Left
	Max
)

type Point struct {
	X, Y int
}

type Game struct {
	Position   Point
	Facing     Direction
	Infected   map[Point]struct{}
	Infections int
}

func NewGame(s string) Game {
	infected := map[Point]struct{}{}
	lines := input.Lines(s)
	midpoint := len(lines) / 2
	for y, line := range lines {
		for x, b := range []byte(line) {
			if b == byte('#') {
				infected[Point{x - midpoint, y - midpoint}] = struct{}{}
			}
		}
	}
	return Game{
		Infected: infected,
	}
}

func (g *Game) Burst() {
	if _, ok := g.Infected[g.Position]; ok {
		g.Facing++
		delete(g.Infected, g.Position)
	} else {
		g.Facing--
		g.Infected[g.Position] = struct{}{}
		g.Infections++
	}
	g.Facing = (g.Facing + Max) % Max
	switch g.Facing {
	case Up:
		g.Position.Y--
	case Down:
		g.Position.Y++
	case Left:
		g.Position.X--
	case Right:
		g.Position.X++
	}
}

func burstsAfter(puzzle string, bursts int) int {
	game := NewGame(puzzle)
	for i := 0; i < bursts; i++ {
		game.Burst()
	}
	return game.Infections
}
