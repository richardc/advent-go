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
			Part1: func(any) any { return burstsAfter(puzzle, Alpha{}, 10_000) },
			Part2: func(any) any { return burstsAfter(puzzle, Omega{}, 10_000_000) },
		},
	)
}

type Direction int

const (
	Up Direction = iota
	Right
	Down
	Left
	MaxDirection
)

type Point struct {
	X, Y int
}

type State int

const (
	Clean State = iota
	Weakened
	Infected
	Flagged
)

type Game struct {
	Position   Point
	Facing     Direction
	Cells      map[Point]State
	Infections int
}

func NewGame(s string) Game {
	cells := map[Point]State{}
	lines := input.Lines(s)
	midpoint := len(lines) / 2
	for y, line := range lines {
		for x, b := range []byte(line) {
			if b == byte('#') {
				cells[Point{x - midpoint, y - midpoint}] = Infected
			}
		}
	}
	return Game{
		Cells: cells,
	}
}

func (g *Game) Burst(virus Virus) {
	virus.Burst(g)
	g.Facing = (g.Facing + MaxDirection) % MaxDirection
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

type Virus interface {
	Burst(*Game)
}

type Alpha struct{}

func (Alpha) Burst(g *Game) {
	switch g.Cells[g.Position] {
	case Infected:
		g.Facing++
		delete(g.Cells, g.Position)
	case Clean:
		g.Facing--
		g.Cells[g.Position] = Infected
		g.Infections++
	}
}

type Omega struct{}

func (Omega) Burst(g *Game) {
	switch g.Cells[g.Position] {
	case Clean:
		g.Facing--
		g.Cells[g.Position] = Weakened
	case Weakened:
		g.Cells[g.Position] = Infected
		g.Infections++
	case Infected:
		g.Facing++
		g.Cells[g.Position] = Flagged
	case Flagged:
		g.Facing += 2
		g.Cells[g.Position] = Clean
	}
}

func burstsAfter(puzzle string, virus Virus, bursts int) int {
	game := NewGame(puzzle)
	for i := 0; i < bursts; i++ {
		game.Burst(virus)
	}
	return game.Infections
}
