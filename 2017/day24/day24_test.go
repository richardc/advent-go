package day24

import (
	_ "embed"
	"testing"
)

//go:embed "example.txt"
var example string

func Test_strongestBridge(t *testing.T) {
	type args struct {
		puzzle string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"example", args{example}, 31},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := strongestBridge(tt.args.puzzle); got != tt.want {
				t.Errorf("strongestBridge() = %v, want %v", got, tt.want)
			}
		})
	}
}
