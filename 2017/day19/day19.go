package day19

import (
	_ "embed"

	"github.com/richardc/advent-go/input"
	"github.com/richardc/advent-go/runner"
	"github.com/richardc/advent-go/slices"
)

//go:embed "input.txt"
var puzzle string

func init() {
	runner.Register(
		runner.Solution{
			Year:  2017,
			Day:   19,
			Input: func() any { return NewMaze(puzzle) },
			Part1: func(i any) any { return visited(i.(Maze)) },
		},
	)
}

type Maze [][]byte

func NewMaze(s string) Maze {
	maze := [][]byte{}
	lines := input.Lines(s)
	for _, line := range lines {
		maze = append(maze, []byte(line))
	}

	// Pad spaced - example.txt keeps getting eaten by editor
	longest := slices.Max(slices.Map(maze, func(l []byte) int { return len(l) }))
	for i := range maze {
		maze[i] = append(maze[i], slices.Map(slices.Range(0, 1+longest-len(maze[i])), func(i int) byte { return byte(' ') })...)
	}
	return maze
}

const (
	down = iota
	up
	left
	right
)

func visited(maze Maze) string {
	seen := []byte{}
	direction := down
	x, y := 0, 0
	for scan := 0; scan < len(maze[0]); scan++ {
		if maze[0][scan] == byte('|') {
			x = scan
		}
	}

walk:
	for {
		switch direction {
		case down:
			y++
		case up:
			y--
		case left:
			x--
		case right:
			x++
		}

		over := maze[y][x]
		switch {
		case over == byte(' '):
			break walk
		case over >= byte('A') && over <= byte('Z'):
			seen = append(seen, over)
		case over == byte('+'):
			switch direction {
			case up, down:
				switch {
				case maze[y][x-1] != ' ':
					direction = left
				case maze[y][x+1] != ' ':
					direction = right
				}
			case left, right:
				switch {
				case maze[y-1][x] != ' ':
					direction = up
				case maze[y+1][x] != ' ':
					direction = down
				}
			}
		}

	}

	return string(seen)
}
