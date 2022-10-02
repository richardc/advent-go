package main

import (
	"testing"
)

func TestXXX(t *testing.T) {
	var tests = []struct {
		input string
		value int
	}{
		{"", 0},
	}

	for _, testcase := range tests {
		testname := testcase.input
		t.Run(testname, func(t *testing.T) {
			ans := XXX(testcase.input)
			if ans != testcase.value {
				t.Errorf("got %v, want %v", ans, testcase.value)
			}
		})
	}
}
