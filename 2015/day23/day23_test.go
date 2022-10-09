package day23

import (
	"testing"
)

func TestNewInstr(t *testing.T) {
	var tests = []struct {
		input  string
		expect instr
	}{
		{"hlf b", hlf("b")},
		{"tpl b", tpl("b")},
		{"inc b", inc("b")},
		{"jmp 10", jmp(10)},
		{"jie b, 20", jie{"b", 20}},
		{"jio b, 30", jio{"b", 30}},
	}

	for _, testcase := range tests {
		testname := testcase.input
		t.Run(testname, func(t *testing.T) {
			ans := newInstr(testcase.input)
			if ans != testcase.expect {
				t.Errorf("got %v, want %v", ans, testcase.expect)
			}
		})
	}
}
