package day19

import (
	_ "embed"
	"testing"
)

//go:embed "example.txt"
var example string

func Test_visited(t *testing.T) {
	type args struct {
		maze Maze
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{"example", args{NewMaze(example)}, "ABCDEF"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := visited(tt.args.maze); got != tt.want {
				t.Errorf("visited() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_steps(t *testing.T) {
	type args struct {
		maze Maze
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"example", args{NewMaze(example)}, 38},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := steps(tt.args.maze); got != tt.want {
				t.Errorf("steps() = %v, want %v", got, tt.want)
			}
		})
	}
}
