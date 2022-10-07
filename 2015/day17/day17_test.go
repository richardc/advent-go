package day17

import (
	"testing"
)

func TestContainers(t *testing.T) {
	ans := containers([]int{20, 15, 10, 5, 5}, 25)
	expected := 4
	if ans != expected {
		t.Errorf("got %v, want %v", ans, expected)
	}
}

func TestSmallContainers(t *testing.T) {
	ans := smallContainers([]int{20, 15, 10, 5, 5}, 25)
	expected := 3
	if ans != expected {
		t.Errorf("got %v, want %v", ans, expected)
	}
}
