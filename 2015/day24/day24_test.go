package day24

import (
	"fmt"
	"testing"
)

func TestSmallestQE(t *testing.T) {
	var tests = []struct {
		input   []int
		baskets int
		value   int
	}{
		{[]int{1, 2, 3, 4, 5, 7, 8, 9, 10, 11}, 3, 99},
		{[]int{1, 2, 3, 4, 5, 7, 8, 9, 10, 11}, 4, 44},
	}

	for _, testcase := range tests {
		testname := fmt.Sprintf("%v", testcase.input)
		t.Run(testname, func(t *testing.T) {
			ans := smallestQE(testcase.input, testcase.baskets)
			if ans != testcase.value {
				t.Errorf("got %v, want %v", ans, testcase.value)
			}
		})
	}
}
