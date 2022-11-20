package day11

import (
	_ "embed"

	"github.com/richardc/advent-go/2019/intcode"
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
			Day:   11,
			Part1: func(any) any { return numPainted(puzzle) },
			Part2: func(any) any { return codePainted(puzzle) },
		},
	)
}

type Point struct {
	X, Y int
}

type Direction = int

const (
	Up = iota
	Right
	Down
	Left
	DirectionEnd
)

type Robot struct {
	Cpu      intcode.Cpu
	Hull     map[Point]int
	Position Point
	Facing   Direction
	Seq      int
}

func newRobot(s string) *Robot {
	cpu := intcode.New(s)
	robot := &Robot{
		Cpu:  cpu,
		Hull: map[Point]int{},
	}
	robot.Cpu.OutputFunc(func(i int) { robot.handleOutput(i) })
	return robot
}

func (r *Robot) handleOutput(i int) {
	switch r.Seq % 2 {
	case 0:
		r.Hull[r.Position] = i
	case 1:
		switch i {
		case 0:
			r.Facing = (r.Facing + DirectionEnd - 1) % DirectionEnd
		case 1:
			r.Facing = (r.Facing + DirectionEnd + 1) % DirectionEnd
		}
		switch r.Facing {
		case Up:
			r.Position.Y--
		case Down:
			r.Position.Y++
		case Right:
			r.Position.X++
		case Left:
			r.Position.X--
		}
		r.Cpu.Input([]int{r.Hull[r.Position]})
	}
	r.Seq++
}

func (r *Robot) run() {
	r.Cpu.Input([]int{r.Hull[r.Position]})
	r.Cpu.Run()
}

func numPainted(s string) int {
	robot := newRobot(s)
	robot.run()
	// fmt.Printf("%+v\n", robot)
	return len(robot.Hull)
}

func codePainted(s string) string {
	robot := newRobot(s)
	robot.Hull[robot.Position] = 1
	robot.run()
	minx, maxx := slices.MinMax(slices.Map(maps.Keys(robot.Hull), func(p Point) int { return p.X }))
	miny, maxy := slices.MinMax(slices.Map(maps.Keys(robot.Hull), func(p Point) int { return p.Y }))
	out := "\n"
	for y := miny; y <= maxy; y++ {
		for x := minx; x <= maxx; x++ {
			if robot.Hull[Point{x, y}] == 1 {
				out += "#"
			} else {
				out += " "
			}
		}
		out += "\n"
	}
	return out
}
