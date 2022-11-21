package day12

import (
	_ "embed"
	"testing"

	"github.com/google/go-cmp/cmp"
)

//go:embed "example1.txt"
var example1 string

//go:embed "example2.txt"
var example2 string

func TestMoons_Step(t *testing.T) {
	type fields struct {
		moons Moons
		steps int
	}
	tests := []struct {
		name   string
		fields fields
		want   Moons
	}{
		{"example 1 1 step", fields{newMoons(example1), 1},
			Moons{[]Moon{
				{Point3{2, -1, 1}, Point3{3, -1, -1}},
				{Point3{3, -7, -4}, Point3{1, 3, 3}},
				{Point3{1, -7, 5}, Point3{-3, 1, -3}},
				{Point3{2, 2, 0}, Point3{-1, -3, 1}},
			}}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := tt.fields.moons
			for i := 0; i < tt.fields.steps; i++ {
				got.Step()
			}
			if diff := cmp.Diff(tt.want, got); diff != "" {
				t.Errorf("Moons.Step() mismatch (-want +got):\n%s", diff)
			}
		})
	}
}

func Test_energyAfter(t *testing.T) {
	type args struct {
		s     string
		steps int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"example 1 for 10", args{example1, 10}, 179},
		{"example 2 for 100", args{example2, 100}, 1940},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := energyAfter(tt.args.s, tt.args.steps); got != tt.want {
				t.Errorf("energyAfter() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_cyclesAt(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"example 1", args{example1}, 2772},
		{"example 2", args{example2}, 4686774924},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := cyclesAt(tt.args.s); got != tt.want {
				t.Errorf("cyclesAt() = %v, want %v", got, tt.want)
			}
		})
	}
}
