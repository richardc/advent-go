package day25

import (
	_ "embed"
	"regexp"

	"github.com/richardc/advent-go/input"
	"github.com/richardc/advent-go/runner"
)

//go:embed input.txt
var puzzle string

func init() {
	runner.Register(runner.Solution{
		Day: 25,
		Input: func() any {
			re := regexp.MustCompile(`(\d+)`)
			captures := re.FindAllStringSubmatch(puzzle, -1)
			return []int{input.MustAtoi(captures[0][0]), input.MustAtoi(captures[1][0])}
		},
		Part1: func(i any) any { return grid(i.([]int)[0], i.([]int)[1]) },
	})
}

func grid(row, col int) int {
	r := 1
	c := 1
	ans := 20151125
	for {
		if row == r && col == c {
			return ans
		}
		// The next cell we have an answer for
		if r == 1 {
			r = c + 1
			c = 1
		} else {
			r--
			c++
		}
		ans = ans * 252533 % 33554393
	}
}
