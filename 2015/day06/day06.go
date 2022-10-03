package day06

import (
	_ "embed"
	"regexp"
	"strconv"

	"github.com/richardc/advent-go/input"
	"github.com/richardc/advent-go/runner"
	"github.com/richardc/advent-go/slices"
)

//go:embed input.txt
var puzzle string

func init() {
	runner.Register(runner.Solution{
		Day:   6,
		Input: func() any { return slices.Map(input.Lines(puzzle), newInstruction) },
		Part1: func(i any) any { return litLights(i.([]instruction)) },
		Part2: func(i any) any { return brightLights(i.([]instruction)) },
	})
}

func litLights(directions []instruction) int {
	lit := [1000 * 1000]byte{}

	for _, d := range directions {
		for x := d.area.x1; x < d.area.x2; x++ {
			for y := d.area.y1; y < d.area.y2; y++ {
				switch d.command {
				case on:
					lit[x+y*1000] = 1
				case off:
					lit[x+y*1000] = 0
				case toggle:
					lit[x+y*1000] ^= 1
				}
			}
		}
	}

	return len(slices.Filter(lit[:], func(x byte) bool { return x == 1 }))
}

func brightLights(directions []instruction) uint {
	lit := [1000 * 1000]uint{}

	for _, d := range directions {
		for x := d.area.x1; x < d.area.x2; x++ {
			for y := d.area.y1; y < d.area.y2; y++ {
				switch d.command {
				case on:
					lit[x+y*1000]++
				case off:
					if lit[x+y*1000] > 0 {
						lit[x+y*1000]--
					}
				case toggle:
					lit[x+y*1000] += 2
				}
			}
		}
	}

	return slices.Sum(lit[:])
}

type rectangle struct {
	x1, y1, x2, y2 int
}

type command int

const (
	invalid = iota
	on
	off
	toggle
)

type instruction struct {
	area    rectangle
	command command
}

func newInstruction(s string) instruction {
	re := regexp.MustCompile(`^(.*) (\d+),(\d+) through (\d+),(\d+)`)
	captures := re.FindStringSubmatch(s)

	var command command
	switch captures[1] {
	case "turn on":
		command = on
	case "turn off":
		command = off
	case "toggle":
		command = toggle
	}
	coord := slices.Map(captures[2:], func(s string) int { v, _ := strconv.Atoi(s); return v })

	return instruction{
		area:    rectangle{coord[0], coord[1], coord[2] + 1, coord[3] + 1},
		command: command,
	}
}
