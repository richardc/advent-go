package day12

import (
	_ "embed"
	"regexp"

	"github.com/richardc/advent-go/input"
	"github.com/richardc/advent-go/maths"
	"github.com/richardc/advent-go/runner"
	"github.com/richardc/advent-go/slices"
)

//go:embed "input.txt"
var puzzle string

func init() {
	runner.Register(
		runner.Solution{
			Year:  2019,
			Day:   12,
			Part1: func(any) any { return energyAfter(puzzle, 1000) },
			Part2: func(any) any { return cyclesAt(puzzle) },
		},
	)
}

type Point3 struct {
	X, Y, Z int
}

func newPoint3(s string) Point3 {
	re := regexp.MustCompile(`<x=(.*?), y=(.*?), z=(.*?)>`)
	caps := re.FindStringSubmatch(s)
	return Point3{
		X: input.MustAtoi(caps[1]),
		Y: input.MustAtoi(caps[2]),
		Z: input.MustAtoi(caps[3]),
	}
}

func (p *Point3) Energy() int {
	return maths.Abs(p.X) + maths.Abs(p.Y) + maths.Abs(p.Z)
}

func (p *Point3) Add(other Point3) {
	p.X += other.X
	p.Y += other.Y
	p.Z += other.Z
}

type Moon struct {
	Position Point3
	Velocity Point3
}

func newMoon(s string) Moon {
	return Moon{
		Position: newPoint3(s),
	}
}

func (m *Moon) Energy() int {
	return m.Position.Energy() * m.Velocity.Energy()
}

type Moons struct {
	Moons []Moon
}

func newMoons(s string) Moons {
	return Moons{
		Moons: slices.Map(input.Lines(s), newMoon),
	}
}

func gravityDelta(a, b int) int {
	if a < b {
		return -1
	}
	if a > b {
		return 1
	}
	return 0
}

func (m *Moons) Step() {
	for i := 0; i < len(m.Moons); i++ {
		for j := i + 1; j < len(m.Moons); j++ {
			delta := gravityDelta(m.Moons[i].Position.X, m.Moons[j].Position.X)
			m.Moons[i].Velocity.X -= delta
			m.Moons[j].Velocity.X += delta

			delta = gravityDelta(m.Moons[i].Position.Y, m.Moons[j].Position.Y)
			m.Moons[i].Velocity.Y -= delta
			m.Moons[j].Velocity.Y += delta

			delta = gravityDelta(m.Moons[i].Position.Z, m.Moons[j].Position.Z)
			m.Moons[i].Velocity.Z -= delta
			m.Moons[j].Velocity.Z += delta
		}
	}

	for i := range m.Moons {
		m.Moons[i].Position.Add(m.Moons[i].Velocity)
	}
}

func (m *Moons) Energy() int {
	return slices.Sum(slices.Map(m.Moons, func(m Moon) int { return m.Energy() }))
}

func energyAfter(s string, steps int) int {
	moons := newMoons(s)
	for i := 0; i < steps; i++ {
		moons.Step()
	}
	return moons.Energy()
}

func cyclesAt(s string) int {
	start := newMoons(s)
	moons := newMoons(s)
	cyclex := 0
	cycley := 0
	cyclez := 0
	for step := 1; !(cyclex != 0 && cycley != 0 && cyclez != 0); step++ {
		moons.Step()
		if cyclex == 0 && slices.All(slices.Zip(start.Moons, moons.Moons), func(s slices.Zipped[Moon]) bool {
			return s.A.Position.X == s.B.Position.X && s.A.Velocity.X == s.B.Velocity.X
		}) {
			cyclex = step
		}
		if cycley == 0 && slices.All(slices.Zip(start.Moons, moons.Moons), func(s slices.Zipped[Moon]) bool {
			return s.A.Position.Y == s.B.Position.Y && s.A.Velocity.Y == s.B.Velocity.Y
		}) {
			cycley = step
		}
		if cyclez == 0 && slices.All(slices.Zip(start.Moons, moons.Moons), func(s slices.Zipped[Moon]) bool {
			return s.A.Position.Z == s.B.Position.Z && s.A.Velocity.Z == s.B.Velocity.Z
		}) {
			cyclez = step
		}
	}

	return maths.LCM(maths.LCM(cyclex, cycley), cyclez)
}
