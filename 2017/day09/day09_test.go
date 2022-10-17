package day09

import (
	_ "embed"
	"testing"
)

func Test_score(t *testing.T) {
	tests := []struct {
		input string
		want  int
	}{
		{"{}", 1},
		{"{{{}}}", 6},
		{"{{},{}}", 5},
		{"{{{},{},{{}}}}", 16},
		{"{<a>,<a>,<a>,<a>}", 1},
		{"{{<ab>},{<ab>},{<ab>},{<ab>}}", 9},
		{"{{<!!>},{<!!>},{<!!>},{<!!>}}", 9},
		{"{{<a!>},{<a!>},{<a!>},{<ab>}}", 3},
	}
	for _, tt := range tests {
		t.Run(tt.input, func(t *testing.T) {
			if got := score(tt.input); got != tt.want {
				t.Errorf("score() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_garbage(t *testing.T) {
	tests := []struct {
		input string
		want  int
	}{
		{"<>", 0},
		{"<random characters>", 17},
		{"<<<<>", 3},
		{"<{!>}>", 2},
		{"<!!>", 0},
		{"<!!!>>", 0},
		{`<{o"i!a,<{i<a>`, 10},
	}
	for _, tt := range tests {
		t.Run(tt.input, func(t *testing.T) {
			if got := garbage(tt.input); got != tt.want {
				t.Errorf("garbage() = %v, want %v", got, tt.want)
			}
		})
	}
}
