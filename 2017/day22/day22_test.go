package day22

import (
	_ "embed"
	"reflect"
	"testing"

	"github.com/google/go-cmp/cmp"
)

//go:embed "example.txt"
var example string

func TestNewGame(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want Game
	}{
		{"example", args{example}, Game{
			Infected: map[Point]struct{}{
				{-1, 0}: {},
				{1, -1}: {},
			},
			Facing: Up,
		}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewGame(tt.args.s); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewGame() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestGame_Burst(t *testing.T) {
	tests := []struct {
		name string
		game Game
		want Game
	}{
		{"step1", NewGame(example), Game{
			Position:   Point{-1, 0},
			Facing:     Left,
			Infections: 1,
			Infected: map[Point]struct{}{
				{-1, 0}: {},
				{0, 0}:  {},
				{1, -1}: {},
			},
		}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.game.Burst()
			if diff := cmp.Diff(tt.game, tt.want); diff != "" {
				t.Errorf("Game.Burst() mismatch (-want +got):\n%s", diff)
			}
		})
	}
}

func Test_burstsAfter(t *testing.T) {
	type args struct {
		puzzle string
		bursts int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"example 70==41", args{example, 70}, 41},
		{"example 10000==5587", args{example, 10_000}, 5_587},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := burstsAfter(tt.args.puzzle, tt.args.bursts); got != tt.want {
				t.Errorf("burstsAfter() = %v, want %v", got, tt.want)
			}
		})
	}
}
