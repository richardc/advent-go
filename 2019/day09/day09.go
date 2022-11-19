package day09

import (
	_ "embed"
	"fmt"

	"github.com/richardc/advent-go/2019/intcode"
	"github.com/richardc/advent-go/runner"
)

//go:embed "input.txt"
var puzzle string

func init() {
	runner.Register(
		runner.Solution{
			Year:  2019,
			Day:   9,
			Part1: func(i any) any { return testMode(puzzle) },
			Part2: func(i any) any { return boost(puzzle) },
		},
	)
}

func testMode(program string) int {
	c := intcode.New(program)
	c.Input([]int{1})
	c.Run()
	out := c.Output()
	if len(out) > 1 {
		fmt.Println(out)
	}
	return out[0]
}

func boost(program string) int {
	c := intcode.New(program)
	c.Input([]int{2})
	c.Run()
	return c.Output()[0]
}
