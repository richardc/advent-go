package main

import (
	"testing"
)

func TestMineMD5(t *testing.T) {
	var tests = []struct {
		input string
		value int
	}{
		{"abcdef", 609043},
		{"pqrstuv", 1048970},
	}

	for _, testcase := range tests {
		testname := testcase.input
		t.Run(testname, func(t *testing.T) {
			ans := mineMD5(testcase.input)
			if ans != testcase.value {
				t.Errorf("got %v, want %v", ans, testcase.value)
			}
		})
	}
}
