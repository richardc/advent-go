package day07

import (
	_ "embed"
	"testing"

	"github.com/go-test/deep"
	"github.com/richardc/advent-go/input"
)

//go:embed "example.txt"
var example string

func Test_buildTree(t *testing.T) {
	type args struct {
		in []string
	}
	tests := []struct {
		name string
		args args
		want Node
	}{
		{"example", args{input.Lines(example)}, Node{
			"tknk", 41, []*Node{
				{"ugml", 68, []*Node{
					{"gyxo", 61, nil},
					{"ebii", 61, nil},
					{"jptl", 61, nil},
				}},
				{"padx", 45, []*Node{
					{"pgba", 66, nil},
					{"havc", 66, nil},
					{"qoyq", 66, nil},
				}},
				{"fwft", 72, []*Node{
					{"ktlj", 57, nil},
					{"cntj", 57, nil},
					{"xhth", 57, nil},
				}},
			},
		}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := buildTree(tt.args.in)
			if diff := deep.Equal(got, tt.want); diff != nil {
				t.Error(diff)
			}
		})
	}
}

func Test_balanceTree(t *testing.T) {
	type args struct {
		n Node
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"example", args{buildTree(input.Lines(example))}, 60},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := balanceTree(tt.args.n); got != tt.want {
				t.Errorf("balanceTree() = %v, want %v", got, tt.want)
			}
		})
	}
}
