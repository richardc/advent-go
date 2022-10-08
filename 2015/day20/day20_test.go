package day20

import (
	"fmt"
	"testing"
)

func TestLowestHouse(t *testing.T) {
	var tests = []struct {
		input  int
		expect int
	}{
		{40, 3},
		{50, 4},
	}

	for _, testcase := range tests {
		t.Run(fmt.Sprintf("%d", testcase.input), func(t *testing.T) {
			ans := lowestHouse(testcase.input)
			if ans != testcase.expect {
				t.Errorf("got %v, want %v", ans, testcase.expect)
			}
		})
	}
}
