package day24

import (
	_ "embed"
	"testing"
)

//go:embed "example.txt"
var example string

func Test_strongestBridge(t *testing.T) {
	type args struct {
		chains []Chain
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"example", args{makeChains(example)}, 31},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := strongestBridge(tt.args.chains); got != tt.want {
				t.Errorf("strongestBridge() = %v, want %v", got, tt.want)
			}
		})
	}
}
