package day25

import (
	_ "embed"
	"testing"

	"github.com/google/go-cmp/cmp"
)

//go:embed "example.txt"
var example string

func TestNewMachine(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want Machine
	}{
		{"example", args{example}, Machine{
			State: 0,
			States: []State{
				{
					When0: Action{
						Write: 1,
						Move:  1,
						Next:  1,
					},
					When1: Action{
						Write: 0,
						Move:  -1,
						Next:  1,
					},
				},
				{
					When0: Action{
						Write: 1,
						Move:  -1,
						Next:  0,
					},
					When1: Action{
						Write: 1,
						Move:  1,
						Next:  0,
					},
				},
			},
			Tape:  map[int]struct{}{},
			Steps: 6,
		}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			got := NewMachine(tt.args.s)
			if diff := cmp.Diff(got, tt.want); diff != "" {
				t.Errorf("NewMachine() mismatch (-want +got):\n%s", diff)
			}
		})
	}
}

func Test_checksum(t *testing.T) {
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
			if got := checksum(tt.args.puzzle); got != tt.want {
				t.Errorf("checksum() = %v, want %v", got, tt.want)
			}
		})
	}
}
