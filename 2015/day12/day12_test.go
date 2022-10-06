package day12

import (
	"testing"
)

func TestSumNumbers(t *testing.T) {
	tests := []struct {
		input string
		value float64
	}{
		{`[1,2,3]`, 6},
	}

	for _, testcase := range tests {
		testname := testcase.input
		t.Run(testname, func(t *testing.T) {
			ans := sumNumbers(testcase.input)
			if ans != testcase.value {
				t.Errorf("got %v, want %v", ans, testcase.value)
			}
		})
	}
}
