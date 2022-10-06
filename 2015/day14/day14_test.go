package day14

import (
	_ "embed"
	"testing"

	"github.com/richardc/advent-go/input"
	"github.com/richardc/advent-go/slices"
)

//go:embed "example.txt"
var example string

func TestWinningDistanceAt(t *testing.T) {
	players := slices.Map(input.Lines(example), newReindeer)
	expected := 1120
	got := winningDistanceAt(players, 1000)
	if got != expected {
		t.Errorf("got %v, want %v", got, expected)
	}
}
