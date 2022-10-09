package slices

import (
	"reflect"
	"strconv"
	"testing"
)

func TestCombinations(t *testing.T) {
	c := Combinations([]int{1, 2, 3}, 2)
	expect := [][]int{
		{1, 2},
		{1, 3},
		{2, 1},
		{2, 3},
		{3, 2},
		{3, 1},
	}

	if !reflect.DeepEqual(c, expect) {
		t.Errorf("got %v, expected %v", c, expect)
	}
}

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

func TestGroupByFunc(t *testing.T) {
	groups := GroupByFunc([]int{1, 1, 2, 2, 2, 1}, func(i int) string { return strconv.Itoa(i) })
	expect := []Group[string, int]{
		{"1", []int{1, 1}},
		{"2", []int{2, 2, 2}},
		{"1", []int{1}},
	}

	if !reflect.DeepEqual(groups, expect) {
		t.Errorf("got %v, expected %v", groups, expect)
	}
}
