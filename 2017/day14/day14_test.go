package day14

import (
	_ "embed"
	"testing"
)

func Test_countSet(t *testing.T) {
	type args struct {
		puzzle string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"example", args{"flqrgnkx"}, 8108},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := countSet(tt.args.puzzle); got != tt.want {
				t.Errorf("countSet() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_countIslands(t *testing.T) {
	type args struct {
		puzzle string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"example", args{"flqrgnkx"}, 1242},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := countIslands(tt.args.puzzle); got != tt.want {
				t.Errorf("countIslands() = %v, want %v", got, tt.want)
			}
		})
	}
}
