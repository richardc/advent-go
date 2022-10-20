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
