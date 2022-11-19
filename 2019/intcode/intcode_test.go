package intcode

import (
	"strings"
	"testing"

	"github.com/google/go-cmp/cmp"
	"github.com/richardc/advent-go/input"
	"github.com/richardc/advent-go/slices"
)

func Test_memory(t *testing.T) {
	type test struct {
		name    string
		program string
		want    []int
	}
	tests := []test{
		{"day 2 example 1", "1,9,10,3,2,3,11,0,99,30,40,50", []int{
			3500, 9, 10, 70,
			2, 3, 11, 0,
			99,
			30, 40, 50,
		}},
		{"day 2 example 2", "1,0,0,0,99", []int{2, 0, 0, 0, 99}},
		{"day 2 example 3", "2,4,4,5,99,0", []int{2, 4, 4, 5, 99, 9801}},
		{"day 2 example 4", "1,1,1,4,99,5,6,0,99", []int{30, 1, 1, 4, 2, 5, 6, 0, 99}},
		{"day 5 example 2", "1002,4,3,4,33", []int{1002, 4, 3, 4, 99}},
		{"day 5 example 4", "1101,100,-1,4,0", []int{1101, 100, -1, 4, 99}},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			cpu := New(tt.program)
			cpu.Run()
			got := cpu.Memory()

			if diff := cmp.Diff(got, tt.want); diff != "" {
				t.Errorf("Output() mismatch (-want +got):\n%s", diff)
			}
		})
	}
}

func Test_outputs(t *testing.T) {
	type test struct {
		name    string
		program string
		want    []int
	}
	quine := "109,1,204,-1,1001,100,1,100,1008,100,16,101,1006,101,0,99"
	tests := []test{
		{"quine", quine, slices.Map(strings.Split(quine, ","), input.MustAtoi)},
		{"16 digit", "1102,34915192,34915192,7,4,7,99,0", []int{1219070632396864}},
		{"middle", "104,1125899906842624,99", []int{1125899906842624}},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			cpu := New(tt.program)
			cpu.Run()
			got := cpu.Output()

			if diff := cmp.Diff(got, tt.want); diff != "" {
				t.Errorf("Output() mismatch (-want +got):\n%s", diff)
			}
		})
	}
}
