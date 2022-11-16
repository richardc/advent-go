package day03

import (
	_ "embed"
	"testing"
)

func Test_solve(t *testing.T) {
	type args struct {
		puzzle string
	}
	tests := []struct {
		name  string
		args  args
		want  int
		want1 int
	}{
		{"example 1", args{"R8,U5,L5,D3\nU7,R6,D4,L4"}, 6, 30},
		{"example 2", args{"R75,D30,R83,U83,L12,D49,R71,U7,L72\nU62,R66,U55,R34,D71,R55,D58,R83"}, 159, 610},
		{"example 3", args{"R98,U47,R26,D63,R33,U87,L62,D20,R33,U53,R51\nU98,R91,D20,R16,D67,R40,U7,R15,U6,R7"}, 135, 410},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, got1 := solve(tt.args.puzzle)
			if got != tt.want {
				t.Errorf("solve() = %v, want %v", got, tt.want)
			}
			if got1 != tt.want1 {
				t.Errorf("solve()1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}
