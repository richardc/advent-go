package day02

import "testing"

func TestPaperNeeded(t *testing.T) {
	var tests = []struct {
		input  string
		expect int
	}{
		{"2x3x4", 58},
		{"1x1x10", 43},
	}

	for _, test := range tests {
		t.Run(test.input, func(t *testing.T) {
			ans := paperNeeded(newParcel(test.input))
			if ans != test.expect {
				t.Errorf("expected %v, got %v", test.expect, ans)
			}
		})
	}
}

func TestRibbonNeeded(t *testing.T) {
	var tests = []struct {
		input  string
		expect int
	}{
		{"2x3x4", 34},
		{"1x1x10", 14},
	}

	for _, test := range tests {
		t.Run(test.input, func(t *testing.T) {
			ans := ribbonNeeded(newParcel(test.input))
			if ans != test.expect {
				t.Errorf("expected %v, got %v", test.expect, ans)
			}
		})
	}
}
