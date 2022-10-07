package day19

import (
	_ "embed"
	"testing"

	"github.com/richardc/advent-go/input"
)

//go:embed "example.txt"
var example string

func TestCountOneReplacement(t *testing.T) {
	var tests = []struct {
		input string
		value int
	}{
		{"HOH", 4},
		{"HOHOHO", 7},
	}

	m := newMachine(input.Lines(example))

	for _, testcase := range tests {
		testname := testcase.input
		t.Run(testname, func(t *testing.T) {
			ans := countOneReplacement(m, testcase.input)
			if ans != testcase.value {
				t.Errorf("got %v, want %v", ans, testcase.value)
			}
		})
	}
}
