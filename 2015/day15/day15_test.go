package day15

import (
	_ "embed"
	"testing"

	"github.com/richardc/advent-go/input"
	"github.com/richardc/advent-go/slices"
)

//go:embed "example.txt"
var example string

func TestBestScore(t *testing.T) {
	ingredients := slices.Map(input.Lines(example), newIngredient)
	ans := bestScore2(ingredients)
	expected := 62842880
	if ans != expected {
		t.Errorf("got %v, expected %v", ans, expected)
	}
}
