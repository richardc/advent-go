package day16

import (
	_ "embed"
	"testing"
)

func Test_dance(t *testing.T) {
	type args struct {
		floor string
		moves []string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{"spin", args{"abcde", []string{"s1"}}, "eabcd"},
		{"exchange", args{"abcde", []string{"x2/3"}}, "abdce"},
		{"partner", args{"abcde", []string{"pb/e"}}, "aecdb"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := dance(tt.args.floor, tt.args.moves); got != tt.want {
				t.Errorf("dance() = %v, want %v", got, tt.want)
			}
		})
	}
}
