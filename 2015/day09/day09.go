package day09

import (
	_ "embed"
	"fmt"
	"strconv"
	"strings"

	"github.com/richardc/advent-go/input"
	"github.com/richardc/advent-go/runner"
	"github.com/richardc/advent-go/slices"
	"golang.org/x/exp/maps"
)

//go:embed input.txt
var puzzle string

func init() {
	runner.Register(runner.Solution{
		Day:   9,
		Input: func() any { return paths(input.Lines(puzzle)) },
		Part1: func(i any) any { return shortestRoute(i.([]path)) },
	})
}

type path struct {
	from  string
	to    string
	miles int
}

func shortestRoute(paths []path) int {
	graph := map[string]int{}
	places := map[string]struct{}{}
	for _, path := range paths {
		places[path.from] = struct{}{}
		places[path.to] = struct{}{}
		graph[fmt.Sprintf("%s -> %s", path.from, path.to)] = path.miles
		graph[fmt.Sprintf("%s -> %s", path.to, path.from)] = path.miles
	}

	routes := slices.Permutations(maps.Keys(places))

	costs := slices.Map(routes, func(route []string) int {
		result := 0
		for i, here := range route[:len(route)-1] {
			there := route[i+1]
			result += graph[fmt.Sprintf("%s -> %s", here, there)]
		}
		return result
	})

	return slices.Min(costs)
}

func newPath(s string) path {
	// London to Dublin = 464
	chunks := strings.Fields(s)
	miles, _ := strconv.Atoi(chunks[4])

	return path{
		from:  chunks[0],
		to:    chunks[2],
		miles: miles,
	}
}

func paths(s []string) []path {
	return slices.Map(s, newPath)
}
