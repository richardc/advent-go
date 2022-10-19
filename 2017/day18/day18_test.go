package day18

import (
	_ "embed"
	"testing"
)

//go:embed "example.txt"
var example string

func Test_duet(t *testing.T) {
	type args struct {
		puzzle string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"example", args{example}, 3},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := duet(tt.args.puzzle); got != tt.want {
				t.Errorf("duet() = %v, want %v", got, tt.want)
			}
		})
	}
}
