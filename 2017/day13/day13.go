package day13

import (
	_ "embed"
	"fmt"
	"strings"

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
			Day:   13,
			Part1: func(any) any { return severityAtTimeZero(puzzle) },
			Part2: func(any) any { return safeTimeToGo(puzzle) },
		},
	)
}

type Scanner struct {
	Column int
	Max    int
}

func NewScanner(s string) Scanner {
	column, max, _ := strings.Cut(s, ": ")
	return Scanner{
		Column: input.MustAtoi(column),
		Max:    input.MustAtoi(max),
	}
}

func (s Scanner) Position(time int) int {
	return (time + s.Column) % (2 * (s.Max - 1))
}

type Firewall struct {
	Scanners []Scanner
}

func NewFirewall(s string) Firewall {
	return Firewall{slices.Map(input.Lines(s), NewScanner)}
}

func (f Firewall) Severity(time int) int {
	total := 0
	for _, scanner := range f.Scanners {
		if scanner.Position(time) == 0 {
			total += scanner.Column * scanner.Max
		}
	}
	return total
}

func (f Firewall) Triggered(time int) bool {
	for _, scanner := range f.Scanners {
		if scanner.Position(time) == 0 {
			return true
		}
	}
	return false
}

func (f Firewall) At(time int) string {
	return strings.Join(slices.Map(f.Scanners, func(s Scanner) string {
		return fmt.Sprintf("%d:%d", s.Column, s.Position(time))
	}), ", ")
}

func severityAtTimeZero(puzzle string) int {
	fw := NewFirewall(puzzle)
	return fw.Severity(0)
}

func safeTimeToGo(puzzle string) int {
	fw := NewFirewall(puzzle)
	for time := 0; ; time++ {
		if !fw.Triggered(time) {
			return time
		}
	}
}
