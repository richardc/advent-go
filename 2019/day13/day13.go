package day13

import (
	_ "embed"

	"github.com/richardc/advent-go/2019/intcode"
	"github.com/richardc/advent-go/runner"
)

//go:embed "input.txt"
var puzzle string

func init() {
	runner.Register(
		runner.Solution{
			Year:  2019,
			Day:   13,
			Part1: func(any) any { return startingBlocks(puzzle) },
			Part2: func(any) any { return winningScore(puzzle) },
		},
	)
}

func startingBlocks(puzzle string) int {
	cpu := intcode.New(puzzle)
	cpu.Run()

	blocks := 0
	for i, val := range cpu.Output() {
		if i%3 == 2 {
			if val == 2 {
				blocks++
			}
		}
	}

	return blocks
}

func winningScore(puzzle string) int {
	cpu := intcode.New(puzzle)
	var step, x, y, ball, bat, score int

	// Insert 2 coins
	cpu.Set(0, 2)

	// Input - make the bat be under the ball as much as possible
	cpu.InputFunc(func() int {
		// fmt.Println("move", ball, bat)
		if ball > bat {
			return 1
		}
		if ball < bat {
			return -1
		}
		return 0
	})

	// Track the ball, bat, and score
	cpu.OutputFunc(func(in int) {
		switch step {
		case 0:
			x = in
		case 1:
			y = in
		case 2:
			// fmt.Println("[", x, y, in, "]")
			if x == -1 && y == 0 {
				score = in
			} else {
				switch in {
				case 3:
					bat = x
				case 4:
					ball = x
				}
			}
		}
		step++
		step %= 3
	})

	cpu.Run()
	return score
}
