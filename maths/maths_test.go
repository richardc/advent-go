package maths

import (
	"testing"
)

func TestAbsDiff(t *testing.T) {
	type args struct {
		a int
		b int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"1 2", args{1, 2}, 1},
		{"2 1", args{2, 1}, 1},
		{"100 80", args{100, 80}, 20},
		{"100 110", args{100, 110}, 10},
		{"-100 80", args{-100, 80}, 180},
		{"-100 -120", args{-100, -120}, 20},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := AbsDiff(tt.args.a, tt.args.b); got != tt.want {
				t.Errorf("AbsDiff() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestGCD(t *testing.T) {
	type args struct {
		a int
		b int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"1 1", args{1, 1}, 1},
		{"1 2", args{1, 2}, 1},
		{"2 4", args{2, 4}, 2},
		{"2 -4", args{2, -4}, 2},
		{"-2 4", args{-2, 4}, 2},
		{"-2 -4", args{-2, -4}, 2},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := GCD(tt.args.a, tt.args.b); got != tt.want {
				t.Errorf("GCD() = %v, want %v", got, tt.want)
			}
		})
	}
}
