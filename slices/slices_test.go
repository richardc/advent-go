package slices

import (
	"reflect"
	"testing"
)

func TestPermutations(t *testing.T) {
	p := Permutations([]int{1, 2, 3})
	expect := [][]int{
		{1, 2, 3},
		{1, 3, 2},
		{2, 1, 3},
		{2, 3, 1},
		{3, 2, 1},
		{3, 1, 2},
	}

	if !reflect.DeepEqual(p, expect) {
		t.Errorf("got %v, expected %v", p, expect)
	}
}
