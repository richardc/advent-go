package main

import (
	_ "embed"
	"fmt"
)

func whatFloor(input string) int {
	count := 0
	for _, c := range input {
		switch c {
		case '(':
			count++
		case ')':
			count--
		}
	}
	return count
}

func goesNegative(input string) int {
	count := 0
	for i, c := range input {
		switch c {
		case '(':
			count++
		case ')':
			count--
		}

		if count < 0 {
			return i + 1
		}
	}
	return -1
}

//go:embed input.txt
var input string

func main() {
	fmt.Println("Part1: ", whatFloor(input))
	fmt.Println("Part2: ", goesNegative(input))
}
