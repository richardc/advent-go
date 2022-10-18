package day12

import (
	_ "embed"
	"testing"
)

//go:embed "example.txt"
var example string

func Test_connectedToZero(t *testing.T) {
	type args struct {
		puzzle string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"example", args{example}, 6},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := connectedToZero(tt.args.puzzle); got != tt.want {
				t.Errorf("connectedToZero() = %v, want %v", got, tt.want)
			}
		})
	}
}
