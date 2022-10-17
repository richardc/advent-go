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
		name  string
		args  args
		want  int
		want1 int
	}{
		{"example", args{[]int{0, 2, 7, 0}}, 5, 4},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, got1 := rebalance(tt.args.memory)
			if got != tt.want {
				t.Errorf("rebalance() got = %v, want %v", got, tt.want)
			}
			if got1 != tt.want1 {
				t.Errorf("rebalance() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}
