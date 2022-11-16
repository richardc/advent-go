package day04

import (
	_ "embed"
	"testing"
)

func Test_validPassword(t *testing.T) {
	tests := []struct {
		password string
		want     bool
	}{
		{"111111", true},
		{"223450", false},
		{"123789", false},
	}
	for _, tt := range tests {
		t.Run(tt.password, func(t *testing.T) {
			if got := validPassword(tt.password); got != tt.want {
				t.Errorf("validPassword() = %v, want %v", got, tt.want)
			}
		})
	}
}
