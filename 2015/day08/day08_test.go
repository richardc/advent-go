package day08

import (
	"testing"
)

func TestCountChars(t *testing.T) {
	var tests = []struct {
		input string
		value int
	}{
		{`""`, 0},
		{`"abc"`, 3},
		{`"aaa\"aaa"`, 7},
		{`"\x27"`, 1},
	}

	for _, testcase := range tests {
		testname := testcase.input
		t.Run(testname, func(t *testing.T) {
			ans := countChars(testcase.input)
			if ans != testcase.value {
				t.Errorf("got %v, want %v", ans, testcase.value)
			}
		})
	}
}

func TestCountEscape(t *testing.T) {
	var tests = []struct {
		input string
		value int
	}{
		{`""`, 6},
		{`"abc"`, 9},
		{`"aaa\"aaa"`, 16},
		{`"\x27"`, 11},
	}

	for _, testcase := range tests {
		testname := testcase.input
		t.Run(testname, func(t *testing.T) {
			ans := countEscape(testcase.input)
			if ans != testcase.value {
				t.Errorf("got %v, want %v", ans, testcase.value)
			}
		})
	}
}
