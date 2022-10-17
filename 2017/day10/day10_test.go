package day10

import (
	_ "embed"
	"testing"
)

func Test_knotHashOne(t *testing.T) {
	type args struct {
		len     int
		lengths []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"example", args{5, []int{3, 4, 1, 5}}, 12},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := knotHashOne(tt.args.len, tt.args.lengths); got != tt.want {
				t.Errorf("knotHashOne() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_knotHash(t *testing.T) {
	tests := []struct {
		input string
		want  string
	}{
		{"", "a2582a3a0e66e6e86e3812dcb672a272"},
		{"AoC 2017", "33efeb34ea91902bb2f59c9920caa6cd"},
		{"1,2,3", "3efbe78a8d82f29979031a4aa0b16a9d"},
		{"1,2,4", "63960835bcdc130f0b66d7ff4f6a5a8e"},
	}
	for _, tt := range tests {
		t.Run(tt.input, func(t *testing.T) {
			if got := knotHash(tt.input); got != tt.want {
				t.Errorf("knotHash() = %v, want %v", got, tt.want)
			}
		})
	}
}
