package day13

import (
	_ "embed"
	"strings"

	"github.com/richardc/advent-go/input"
	"github.com/richardc/advent-go/runner"
)

//go:embed "input.txt"
var puzzle string

func init() {
	runner.Register(
		runner.Solution{
			Year:  2017,
			Day:   13,
			Part1: func(any) any { return severityAtTimeZero(puzzle) },
		},
	)
}

type Scanner struct {
	Column    int
	Max       int
	Position  int
	Direction int
}

func NewScanner(s string) Scanner {
	column, max, _ := strings.Cut(s, ": ")
	return Scanner{
		Column:    input.MustAtoi(column),
		Max:       input.MustAtoi(max),
		Direction: 1,
	}
}

func (s *Scanner) Tick() {
	s.Position += s.Direction
	switch s.Direction {
	case 1:
		if s.Position >= s.Max {
			s.Position = s.Max - 2
			s.Direction = -1
		}
	case -1:
		if s.Position < 0 {
			s.Position = 1
			s.Direction = 1
		}
	}
}

func (s Scanner) AtTop() bool {
	return s.Position == 0
}

type Firewall struct {
	Scanners []*Scanner
}

func NewFirewall(s string) Firewall {
	lines := input.Lines(s)
	last := NewScanner(lines[len(lines)-1])

	scanners := make([]*Scanner, last.Column+1)
	for _, scan := range input.Lines(s) {
		scanner := NewScanner(scan)
		scanners[scanner.Column] = &scanner
	}

	return Firewall{scanners}
}

func (f *Firewall) Tick() {
	for _, scanner := range f.Scanners {
		if scanner != nil {
			scanner.Tick()
		}
	}
}

func (f Firewall) Severity() int {
	total := 0
	for column, scanner := range f.Scanners {
		if scanner != nil && scanner.AtTop() {
			total += column * scanner.Max
		}
		f.Tick()
	}
	return total
}

func severityAtTimeZero(puzzle string) int {
	fw := NewFirewall(puzzle)
	return fw.Severity()
}
