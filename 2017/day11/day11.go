package day11

import (
	_ "embed"
	"strings"

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
			Day:   11,
			Part1: func(any) any { return hexDistance(strings.TrimSpace(puzzle)) },
		},
	)
}

// This method for mapping distances across a hex grid into 3 axes from
// https://archive.ph/20141214082648/http://keekerdc.com/2011/03/hexagon-grids-coordinate-systems-and-distance-calculations/#60%
func hexDistance(puzzle string) int {
	x, y, z := 0, 0, 0
	for _, step := range strings.Split(puzzle, ",") {
		switch step {
		case "nw":
			x--
			y++
		case "n":
			y++
			z--
		case "ne":
			x++
			z--
		case "se":
			x++
			y--
		case "s":
			y--
			z++
		case "sw":
			x--
			z++
		}
	}
	return slices.Max([]int{math.Abs(x), math.Abs(y), math.Abs(z)})
}
