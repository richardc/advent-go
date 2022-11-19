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
			Year: 2019,
			Day:  9,
			Input: func() any {
				cpu := intcode.New(puzzle)
				return &cpu
			},
			Part1: func(i any) any { return testMode(i.(*intcode.Cpu)) },
		},
	)
}

func testMode(c *intcode.Cpu) int {
	c.Input([]int{1})
	c.Run()
	out := c.Output()
	if len(out) != 0 {
		fmt.Println(out)
	}
	return out[0]
}
