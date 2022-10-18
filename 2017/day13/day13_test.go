package day13

import (
	_ "embed"
	"testing"
)

//go:embed "example.txt"
var example string

func Test_severityAtTimeZero(t *testing.T) {
	type args struct {
		puzzle string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"example", args{example}, 24},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := severityAtTimeZero(tt.args.puzzle); got != tt.want {
				t.Errorf("severityAtTimeZero() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_safeTimeToGo(t *testing.T) {
	type args struct {
		puzzle string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"example", args{example}, 10},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := safeTimeToGo(tt.args.puzzle); got != tt.want {
				t.Errorf("safeTimeToGo() = %v, want %v", got, tt.want)
			}
		})
	}
}
