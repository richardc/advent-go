package day06

import (
	"testing"
)

func TestLitLights(t *testing.T) {
	var tests = []struct {
		name  string
		input []instruction
		value int
	}{
		{"all", []instruction{
			newInstruction("turn on 0,0 through 999,999"),
		}, 1_000_000},
		{"toggle", []instruction{
			newInstruction("turn on 0,0 through 999,999"),
			newInstruction("toggle 0,0 through 999,0"),
		}, 999_000},
		{"mini", []instruction{
			newInstruction("turn on 0,0 through 999,999"),
			newInstruction("toggle 0,0 through 999,0"),
			newInstruction("turn off 499,499 through 500,500"),
		}, 998_996},
	}

	for _, testcase := range tests {
		t.Run(testcase.name, func(t *testing.T) {
			ans := litLights(testcase.input)
			if ans != testcase.value {
				t.Errorf("got %v, want %v", ans, testcase.value)
			}
		})
	}
}
