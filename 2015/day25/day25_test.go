package day25

import (
	_ "embed"
	"testing"
)

func Test_grid(t *testing.T) {
	type args struct {
		x int
		y int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"1,1", args{1, 1}, 20151125},
		{"1,2", args{1, 2}, 18749137},
		{"2,1", args{2, 1}, 31916031},
		{"3,1", args{3, 1}, 16080970},
		{"2,2", args{2, 2}, 21629792},
		{"1,3", args{1, 3}, 17289845},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := grid(tt.args.x, tt.args.y); got != tt.want {
				t.Errorf("grid() = %v, want %v", got, tt.want)
			}
		})
	}
}
