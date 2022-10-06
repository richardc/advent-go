package day13

import (
	_ "embed"
	"testing"

	"github.com/richardc/advent-go/input"
)

//go:embed "example.txt"
var example string

func TestHappiestSeating(t *testing.T) {
	expected := 330
	got := happiestSeating(input.Lines(example), false)
	if got != expected {
		t.Errorf("got %v, expected %v", got, expected)
	}
}
