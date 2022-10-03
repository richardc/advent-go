package day04

import (
	_ "embed"
	"fmt"
	"strings"

	"crypto/md5"

	"github.com/richardc/advent-go/runner"
)

//go:embed input.txt
var input string

func init() {
	runner.Register(runner.Solution{Day: 4,
		Input: func() any { return strings.Trim(input, "\n") },
		Part1: func(in any) any { return mineMD5(in.(string)) },
		Part2: func(in any) any { return mineMD56(in.(string)) },
	})
}

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
