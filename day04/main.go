package main

import (
	_ "embed"
	"fmt"
	"strings"

	"crypto/md5"
)

func mineMD5(s string) int {
	guess := 0
	for {
		sum := fmt.Sprintf("%x", md5.Sum([]byte(fmt.Sprintf("%s%d", s, guess))))
		if sum[:5] == "00000" {
			return guess
		}

		guess++
	}
}

func mineMD56(s string) int {
	guess := 0
	for {
		sum := fmt.Sprintf("%x", md5.Sum([]byte(fmt.Sprintf("%s%d", s, guess))))
		if sum[:6] == "000000" {
			return guess
		}

		guess++
	}
}

//go:embed input.txt
var input string

func main() {
	part1 := mineMD5(strings.Trim(input, "\n"))
	fmt.Println("Part 1", part1)
	part2 := mineMD56(strings.Trim(input, "\n"))
	fmt.Println("Part 2", part2)
}
