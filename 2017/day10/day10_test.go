package day10

import (
	_ "embed"
	"testing"
)

func Test_knotHash(t *testing.T) {
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
			if got := knotHash(tt.args.len, tt.args.lengths); got != tt.want {
				t.Errorf("knotHash() = %v, want %v", got, tt.want)
			}
		})
	}
}
