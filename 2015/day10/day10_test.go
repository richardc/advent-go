package day10

import (
	"testing"
)

func TestLookAndSay(t *testing.T) {
	var tests = []struct {
		input    string
		expected string
	}{
		{"1", "11"},
		{"11", "21"},
		{"21", "1211"},
		{"1211", "111221"},
		{"111221", "312211"},
	}

	for _, testcase := range tests {
		testname := testcase.input
		t.Run(testname, func(t *testing.T) {
			ans := lookAndSay(testcase.input)
			if ans != testcase.expected {
				t.Errorf("got %v, want %v", ans, testcase.expected)
			}
		})
	}
}
