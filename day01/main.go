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

//go:embed input.txt
var input string

func main() {
	fmt.Println("Part1: ", whatFloor(input))
}
