package day06

import (
	_ "embed"
	"testing"
)

func Test_rebalance(t *testing.T) {
	type args struct {
		memory []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"example", args{[]int{0, 2, 7, 0}}, 5},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := rebalance(tt.args.memory); got != tt.want {
				t.Errorf("rebalance() = %v, want %v", got, tt.want)
			}
		})
	}
}
