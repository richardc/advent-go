package day10

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
		Day:   10,
		Part1: func(i any) any { return lookAndSayIterated(strings.TrimSpace(puzzle), 40) },
		Part2: func(i any) any { return lookAndSayIterated(strings.TrimSpace(puzzle), 50) },
	})
}

func lookAndSay(s string) string {
	groups := slices.GroupBy([]byte(s))
	parts := slices.Map(groups, func(g slices.Group[byte, byte]) string {
		return fmt.Sprintf("%d%c", len(g.Group), g.Key)
	})
	return strings.Join(parts, "")
}

func lookAndSayIterated(s string, count int) int {
	for i := 0; i < count; i++ {
		s = lookAndSay(s)
	}
	return len(s)
}
