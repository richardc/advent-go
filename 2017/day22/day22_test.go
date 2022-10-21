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
			Cells: map[Point]State{
				{-1, 0}: Infected,
				{1, -1}: Infected,
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
			Cells: map[Point]State{
				{-1, 0}: Infected,
				{0, 0}:  Infected,
				{1, -1}: Infected,
			},
		}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.game.Burst(Alpha{})
			if diff := cmp.Diff(tt.game, tt.want); diff != "" {
				t.Errorf("Game.Burst() mismatch (-want +got):\n%s", diff)
			}
		})
	}
}

func Test_burstsAfter(t *testing.T) {
	type args struct {
		puzzle string
		virus  Virus
		bursts int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"alpha 70==41", args{example, Alpha{}, 70}, 41},
		{"alpha 10000==5587", args{example, Alpha{}, 10_000}, 5_587},
		{"omega 70==41", args{example, Omega{}, 100}, 26},
		{"omega 10000==5587", args{example, Omega{}, 10_000_000}, 2_511_944},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := burstsAfter(tt.args.puzzle, tt.args.virus, tt.args.bursts); got != tt.want {
				t.Errorf("burstsAfter() = %v, want %v", got, tt.want)
			}
		})
	}
}
