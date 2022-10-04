package day09

import (
	_ "embed"
	"testing"

	"github.com/richardc/advent-go/input"
)

//go:embed example.txt
var example string

func TestShortestRoute(t *testing.T) {
	distances := paths(input.Lines(example))
	shortest := shortestRoute(distances)
	expected := 605

	if shortest != expected {
		t.Errorf("got %v, want %v", shortest, expected)
	}
}
