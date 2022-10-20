package day20

import (
	_ "embed"
	"strings"

	"github.com/richardc/advent-go/input"
	"github.com/richardc/advent-go/math"
	"github.com/richardc/advent-go/runner"
	"github.com/richardc/advent-go/slices"
)

//go:embed "input.txt"
var puzzle string

func init() {
	runner.Register(
		runner.Solution{
			Year:  2017,
			Day:   20,
			Input: func() any { return slices.Map(input.Lines(puzzle), NewParticle) },
			Part1: func(i any) any { return staysClosest(i.([]Particle)) },
		},
	)
}

type Vector struct {
	X, Y, Z int
}

func NewVector(s string) Vector {
	toks := strings.Split(s[3:len(s)-1], ",")
	return Vector{
		X: input.MustAtoi(toks[0]),
		Y: input.MustAtoi(toks[1]),
		Z: input.MustAtoi(toks[2]),
	}
}

func (v Vector) Distance(other Vector) int {
	return math.Abs(v.X-other.X) + math.Abs(v.Y-other.Y) + math.Abs(v.Z-other.Z)
}

type Particle struct {
	Position     Vector
	Velocity     Vector
	Acceleration Vector
}

func NewParticle(s string) Particle {
	toks := strings.Split(s, ", ")
	return Particle{
		Position:     NewVector(toks[0]),
		Velocity:     NewVector(toks[1]),
		Acceleration: NewVector(toks[2]),
	}
}

// Over time, the dominating factor is the Acceleration
func staysClosest(particles []Particle) int {
	index := 0
	zero := Vector{}
	max := particles[0].Acceleration.Distance(zero)
	for i, p := range particles {
		a := p.Acceleration.Distance(zero)
		if a < max {
			index = i
			max = a
		}
	}
	return index
}
