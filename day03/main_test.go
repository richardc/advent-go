package main

import (
	"testing"
)

func TestCountDeliveries(t *testing.T) {
	var tests = []struct {
		input string
		value int
	}{
		{">", 2},
		{"^>v<", 4},
		{"^v^v^v^v^v", 2},
	}

	for _, testcase := range tests {
		testname := testcase.input
		t.Run(testname, func(t *testing.T) {
			ans := countDeliveries(testcase.input)
			if ans != testcase.value {
				t.Errorf("got %v, want %v", ans, testcase.value)
			}
		})
	}
}
