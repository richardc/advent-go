package day07

import (
	_ "embed"
	"strings"

	"github.com/richardc/advent-go/input"
	"github.com/richardc/advent-go/runner"
	slcs "github.com/richardc/advent-go/slices"
	"golang.org/x/exp/maps"
	"golang.org/x/exp/slices"
)

//go:embed "input.txt"
var puzzle string

func init() {
	runner.Register(
		runner.Solution{
			Year:  2017,
			Day:   7,
			Input: func() any { return buildTree(input.Lines(puzzle)) },
			Part1: func(i any) any { return i.(Node).Name },
			Part2: func(i any) any { return balanceTree(i.(Node)) },
		},
	)
}

type Node struct {
	Name     string
	Weight   int
	Children []*Node
}

func newNode(s string) Node {
	toks := strings.Fields(s)
	weight := input.MustAtoi(toks[1][1 : len(toks[1])-1])
	return Node{
		Name:   toks[0],
		Weight: weight,
	}
}

func buildTree(in []string) Node {
	nodes := map[string]*Node{}

	// Build node objects
	for _, description := range in {
		node := newNode(description)
		nodes[node.Name] = &node
	}

	// Resolve children
	children := []string{}
	for _, description := range in {
		if node, kids, ok := strings.Cut(description, " -> "); ok {
			parent, _, _ := strings.Cut(node, " ")
			for _, name := range strings.Split(kids, ", ") {
				child := nodes[name]
				nodes[parent].Children = append(nodes[parent].Children, child)
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

func (n Node) TotalWeight() int {
	total := n.Weight
	for _, child := range n.Children {
		total += child.TotalWeight()
	}
	return total
}

func balanceTree(n Node) int {
	if len(n.Children) == 0 {
		// No children, implicitly balanced
		return 0
	}

	// Depth first, find imbalanced tree
	for _, child := range n.Children {
		answer := balanceTree(*child)
		if answer > 0 {
			return answer
		}
	}

	childs := slcs.Map(n.Children, func(n *Node) int { return n.TotalWeight() })
	weights := slcs.Counts(childs)

	if len(weights) == 1 {
		// Children are balanced
		return 0
	}

	// Rebalance the outlier
	keys := maps.Keys(weights)
	slices.SortFunc(keys, func(a, b int) bool { return weights[a] < weights[b] })

	correcting := keys[0]
	delta := keys[1] - keys[0]

	index := slices.Index(childs, correcting)
	return n.Children[index].Weight + delta
}
