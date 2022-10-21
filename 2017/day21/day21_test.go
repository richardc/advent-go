package day21

import (
	_ "embed"
	"testing"

	"github.com/google/go-cmp/cmp"
)

func TestGame_Mirror(t *testing.T) {
	tests := []struct {
		name  string
		input Game
		want  Game
	}{
		{"glider", NewGameFrom(".#./..#/###"), NewGameFrom("###/..#/.#.")},
		{"L", NewGameFrom("##/.#"), NewGameFrom(".#/##")},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := tt.input.Mirror()
			if diff := cmp.Diff(got.String(), tt.want.String()); diff != "" {
				t.Errorf("Game.Mirror() mismatch (-want +got):\n%s", diff)
			}
		})
	}
}

func TestGame_Rotate(t *testing.T) {
	tests := []struct {
		name  string
		input Game
		want  Game
	}{
		{"glider", NewGameFrom(".#./..#/###"), NewGameFrom("#../#.#/##.")},
		{"L", NewGameFrom("##/#."), NewGameFrom("##/.#")},
		{"spot1", NewGameFrom("../.#"), NewGameFrom("../#.")},
		{"spot2", NewGameFrom("../#."), NewGameFrom("#./..")},
		{"spot3", NewGameFrom("#./.."), NewGameFrom(".#/..")},
		{"spot4", NewGameFrom(".#/.."), NewGameFrom("../.#")},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := tt.input.Rotate()
			if diff := cmp.Diff(got.Lines(), tt.want.Lines()); diff != "" {
				t.Errorf("Game.Rotate() mismatch (-want +got):\n%s", diff)
			}
		})
	}
}

//go:embed "example.txt"
var example string

func Test_countPixels(t *testing.T) {
	type args struct {
		r    Rules
		iter int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"example", args{NewRules(example), 2}, 12},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := countPixels(tt.args.r, tt.args.iter); got != tt.want {
				t.Errorf("countPixels() = %v, want %v", got, tt.want)
			}
		})
	}
}
