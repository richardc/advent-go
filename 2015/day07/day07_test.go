package day07

import (
	_ "embed"
	"reflect"
	"testing"

	"github.com/richardc/advent-go/input"
)

//go:embed example.txt
var example string

func TestResolveWiring(t *testing.T) {
	wiring := resolveWiring(input.Lines(example))
	expected := map[string]Value{
		"d": 72,
		"e": 507,
		"f": 492,
		"g": 114,
		"h": 65412,
		"i": 65079,
		"x": 123,
		"y": 456,
	}

	if !reflect.DeepEqual(wiring, expected) {
		t.Errorf("got %v, want %v", wiring, expected)
	}
}
