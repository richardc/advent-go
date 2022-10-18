package day12

import (
	_ "embed"
	"strings"

	"github.com/richardc/advent-go/input"
	"github.com/richardc/advent-go/runner"
	"github.com/richardc/advent-go/slices"
)

//go:embed "input.txt"
var puzzle string

func init() {
	runner.Register(
		runner.Solution{
			Year:  2017,
			Day:   12,
			Part1: func(any) any { return connectedToZero(puzzle) },
		},
	)
}

type Node struct {
	Id   int
	Conn []*Node
}

func parseNodes(puzzle string) []*Node {
	lines := input.Lines(puzzle)
	graph := make([]*Node, len(lines))
	for i := range graph {
		graph[i] = &Node{Id: i}
	}

	for _, line := range lines {
		left, edges, _ := strings.Cut(line, " <-> ")
		idx := input.MustAtoi(left)
		for _, edge := range strings.Split(edges, ", ") {
			node := input.MustAtoi(edge)
			graph[node].AddEdge(graph[idx])
			graph[idx].AddEdge(graph[node])
		}
	}
	return graph
}

func (n *Node) AddEdge(other *Node) {
	n.Conn = append(n.Conn, other)
}

func (n *Node) ConnectsTo(id int) bool {
	visited := map[*Node]struct{}{}
	queue := []*Node{n}
	for len(queue) > 0 {
		next := []*Node{}
		for _, node := range queue {
			if node.Id == id {
				return true
			}
			for _, child := range node.Conn {
				if _, ok := visited[child]; ok {
					continue
				}
				next = append(next, child)
				visited[child] = struct{}{}
			}
		}
		queue = next
	}
	return false
}

func connectedToZero(puzzle string) int {
	graph := parseNodes(puzzle)

	return len(slices.Filter(graph, func(n *Node) bool { return n.ConnectsTo(0) }))
}
