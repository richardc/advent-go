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
		{"123444", true},
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

func Test_validPassword2(t *testing.T) {
	tests := []struct {
		password string
		want     bool
	}{
		{"111111", false},
		{"111122", true},
		{"112233", true},
		{"123444", false},
	}
	for _, tt := range tests {
		t.Run(tt.password, func(t *testing.T) {
			if got := validPassword2(tt.password); got != tt.want {
				t.Errorf("validPassword2() = %v, want %v", got, tt.want)
			}
		})
	}
}
