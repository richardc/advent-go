package day18

import (
	_ "embed"
	"testing"
)

//go:embed "example.txt"
var example string

func TestIterateLit(t *testing.T) {
	life := newLife(example)
	got := iterateLit(life, 4)
	expected := 4
	if got != expected {
		t.Errorf("got %v, want %v", got, expected)
	}
}
