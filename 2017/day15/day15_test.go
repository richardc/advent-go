package day15

import (
	_ "embed"
	"testing"
)

func Test_judge(t *testing.T) {
	type args struct {
		genA   func() int
		genB   func() int
		rounds int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"5 rounds", args{generatorA(65), generatorB(8921), 5}, 1},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := judge(tt.args.genA, tt.args.genB, tt.args.rounds); got != tt.want {
				t.Errorf("judge() = %v, want %v", got, tt.want)
			}
		})
	}
}
