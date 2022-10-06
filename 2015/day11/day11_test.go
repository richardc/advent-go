package day11

import (
	_ "embed"
	"testing"
)

func TestIncrementPassword(t *testing.T) {
	tests := []struct {
		input  string
		expect string
	}{
		{"fredisdead", "fredisdeae"},
	}

	for _, tt := range tests {
		t.Run(tt.input, func(t *testing.T) {
			got := incrementPassword(tt.input)
			if got != tt.expect {
				t.Errorf("got %v, expect %v", got, tt.expect)
			}
		})
	}
}

func TestSkipIOL(t *testing.T) {
	tests := []struct {
		input  string
		expect string
	}{
		{"fredisdead", "fredjaaaaa"},
	}

	for _, tt := range tests {
		t.Run(tt.input, func(t *testing.T) {
			got := skipIOL(tt.input)
			if got != tt.expect {
				t.Errorf("got %v, expect %v", got, tt.expect)
			}
		})
	}
}

func TestNextPassword(t *testing.T) {
	tests := []struct {
		input  string
		expect string
	}{
		{"abcdefgh", "abcdffaa"},
		{"ghijklmn", "ghjaabcc"},
	}

	for _, tt := range tests {
		t.Run(tt.input, func(t *testing.T) {
			if got := nextPassword(tt.input); got != tt.expect {
				t.Errorf("got %v, expect %v", got, tt.expect)
			}
		})
	}
}
