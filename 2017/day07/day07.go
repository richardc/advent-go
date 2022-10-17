package day07

import (
	_ "embed"
	"strings"

	"github.com/richardc/advent-go/input"
	"github.com/richardc/advent-go/runner"
	"golang.org/x/exp/maps"
)

//go:embed "input.txt"
var puzzle string

func init() {
	runner.Register(
		runner.Solution{
			Year:  2017,
			Day:   7,
			Input: func() any { return buildTree(input.Lines(puzzle)) },
			Part1: func(i any) any { return i.(Node).name },
		},
	)
}

type Node struct {
	name     string
	weight   int
	children []*Node
}

func newNode(s string) Node {
	toks := strings.Fields(s)
	weight := input.MustAtoi(toks[1][1 : len(toks[1])-1])
	return Node{
		name:   toks[0],
		weight: weight,
	}
}

func buildTree(in []string) Node {
	nodes := map[string]*Node{}

	// Build node objects
	for _, description := range in {
		node := newNode(description)
		nodes[node.name] = &node
	}

	// Resolve children
	children := []string{}
	for _, description := range in {
		if node, kids, ok := strings.Cut(description, " -> "); ok {
			parent, _, _ := strings.Cut(node, " ")
			for _, name := range strings.Split(kids, ", ") {
				child := nodes[name]
				nodes[parent].children = append(nodes[parent].children, child)
				children = append(children, name)
			}
		}
	}

	// Delete all the children nodes, only root will remain
	for _, c := range children {
		delete(nodes, c)
	}

	if len(nodes) > 1 {
		panic("Should only have root node")
	}

	return *maps.Values(nodes)[0]
}
