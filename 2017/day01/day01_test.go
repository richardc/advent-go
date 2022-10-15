package day01

import (
	_ "embed"
	"testing"
)

func Test_uncaptcha(t *testing.T) {
	type args struct {
		puzzle string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"1122 = 3", args{"1122"}, 3},
		{"1111 = 4", args{"1111"}, 4},
		{"1234 = 0", args{"1234"}, 0},
		{"91212129 = 9", args{"91212129"}, 9},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := uncaptcha(tt.args.puzzle); got != tt.want {
				t.Errorf("uncaptcha() = %v, want %v", got, tt.want)
			}
		})
	}
}
