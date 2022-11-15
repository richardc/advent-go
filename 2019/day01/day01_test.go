package day01

import (
	_ "embed"
	"testing"
)

func Test_moduleFuel(t *testing.T) {
	type args struct {
		mass int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"12", args{12}, 2},
		{"14", args{14}, 2},
		{"1969", args{1969}, 654},
		{"100756", args{100756}, 33583},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := moduleFuel(tt.args.mass); got != tt.want {
				t.Errorf("moduleFuel() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_moduleFuelWithFuel(t *testing.T) {
	type args struct {
		mass int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"14", args{14}, 2},
		{"1969", args{1969}, 966},
		{"100756", args{100756}, 50346},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := moduleFuelWithFuel(tt.args.mass); got != tt.want {
				t.Errorf("moduleFuelWithFuel() = %v, want %v", got, tt.want)
			}
		})
	}
}
