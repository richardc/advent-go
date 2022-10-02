package main

import (
	_ "embed"
	"fmt"
)

func XXX(s string) int {
	return 0
}

//go:embed input.txt
var input string

func main() {
	part1 := 0
	fmt.Println("Part 1", part1)
	part2 := 0
	fmt.Println("Part 2", part2)
}
