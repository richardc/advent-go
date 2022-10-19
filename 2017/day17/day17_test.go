package day17

import (
	_ "embed"
	"reflect"
	"testing"
)

func Test_spinlock(t *testing.T) {
	type args struct {
		step       int
		iterations int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{"one", args{3, 1}, []int{0, 1}},
		{"two", args{3, 2}, []int{0, 2, 1}},
		{"three", args{3, 3}, []int{0, 2, 3, 1}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := spinlock(tt.args.step, tt.args.iterations); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("spinlock() = %v, want %v", got, tt.want)
			}
		})
	}
}
