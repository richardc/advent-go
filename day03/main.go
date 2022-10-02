package main

import (
	_ "embed"
	"fmt"
	"strings"

	"advent2015/pkg/slices"
)

type location struct{ x, y int }

func calcDeliveries(moves string) []location {
	var pos location
	deliveries := []location{pos}
	for _, c := range moves {
		switch c {
		case '>':
			pos.x++
		case '<':
			pos.x--
		case '^':
			pos.y++
		case 'v':
			pos.y--
		}
		deliveries = append(deliveries, pos)

	}
	return deliveries
}

func countDeliveries(moves string) int {
	return len(slices.Unique(calcDeliveries(moves)))
}

func evenIndexes() func(rune) rune {
	i := 0
	return func(r rune) rune {
		i++
		if i%2 == 0 {
			return r
		}
		return -1
	}
}

func oddIndexes() func(rune) rune {
	i := 0
	return func(r rune) rune {
		i++
		if i%2 == 1 {
			return r
		}
		return -1
	}
}

func countRoboDeliveries(moves string) int {
	santa := calcDeliveries(strings.Map(evenIndexes(), moves))
	robot := calcDeliveries(strings.Map(oddIndexes(), moves))
	all := append(santa, robot...)
	return len(slices.Unique(all))
}

//go:embed input.txt
var input string

func main() {
	part1 := countDeliveries(input)
	fmt.Println("Part 1", part1)
	part2 := countRoboDeliveries(input)
	fmt.Println("Part 2", part2)
}
