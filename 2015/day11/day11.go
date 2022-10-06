package day11

import (
	_ "embed"
	"fmt"
	"strings"

	"github.com/richardc/advent-go/runner"
	"github.com/richardc/advent-go/slices"
)

//go:embed input.txt
var puzzle string

func init() {
	runner.Register(runner.Solution{
		Day:   11,
		Part1: func(i any) any { return nextPassword(strings.TrimSpace(puzzle)) },
	})
}

func pairIndex(s string) int {
	for i := range s[:len(s)-1] {
		if s[i] == s[i+1] {
			return i
		}
	}
	return -1
}

func twoPair(s string) bool {
	if firstPair := pairIndex(s); firstPair != -1 && firstPair < len(s)-2 {
		if secondPair := pairIndex(s[firstPair+2:]); secondPair != -1 {
			return true
		}
	}
	return false
}

func threeRun(s string) bool {
	for i := range s[:len(s)-2] {
		if s[i] == s[i+1]-1 && s[i] == s[i+2]-2 {
			return true
		}
	}
	return false
}

func isLegal(s string) bool {
	return twoPair(s) && threeRun(s)
}

func skipIOL(s string) string {
	if i := strings.IndexAny(s, "iol"); i != -1 {
		return fmt.Sprintf("%s%c%s", s[:i], s[i]+1, strings.Repeat("a", len(s)-i-1))
	}
	return s
}

func incrementPassword(s string) string {
	// base26 decode
	value := 0
	for i := range s {
		value = value*26 + int(s[i]-'a')
	}

	value++

	result := make([]byte, len(s))
	for i := range s {
		result[i] = byte('a') + byte(value%26)
		value /= 26

	}

	slices.Reverse(result)
	return string(result)
}

func nextPassword(s string) string {
	password := s
	for {
		password = skipIOL(incrementPassword(password))

		if isLegal(password) {
			return password
		}
	}
}
