package day06

import (
	_ "embed"
	"testing"
)

//go:embed "example.txt"
var example string

func TestStarmap_Orbits(t *testing.T) {
	type fields struct {
		starmap string
	}
	tests := []struct {
		name  string
		input fields
		want  int
	}{
		{"example", fields{example}, 42},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := newStarmap(tt.input.starmap)
			if got := s.Orbits(); got != tt.want {
				t.Errorf("Starmap.Orbits() = %v, want %v", got, tt.want)
			}
		})
	}
}

//go:embed "example2.txt"
var example2 string

func TestStarmap_Transfers(t *testing.T) {
	type fields struct {
		starmap string
	}
	tests := []struct {
		name  string
		input fields
		want  int
	}{
		{"example2", fields{example2}, 4},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := newStarmap(tt.input.starmap)
			if got := s.Transfers(); got != tt.want {
				t.Errorf("Starmap.Transfers() = %v, want %v", got, tt.want)
			}
		})
	}
}
