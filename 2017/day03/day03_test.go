package day03

import (
	_ "embed"
	"testing"
)

func Test_spiralManhattan(t *testing.T) {
	type args struct {
		puzzle int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"1", args{1}, 0},
		{"12", args{12}, 3},
		{"23", args{23}, 2},
		{"1024", args{1024}, 31},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := spiralManhattan(tt.args.puzzle); got != tt.want {
				t.Errorf("spiralManhattan() = %v, want %v", got, tt.want)
			}
		})
	}
}
