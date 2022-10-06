package day12

import (
	_ "embed"
	"encoding/json"
	"fmt"

	"github.com/richardc/advent-go/runner"
)

//go:embed input.txt
var puzzle string

func init() {
	runner.Register(runner.Solution{
		Day:   12,
		Part1: func(any) any { return sumNumbers(puzzle) },
	})
}

func eval(d any) float64 {
	switch d := d.(type) {
	case float64:
		return d
	case []any:
		total := 0.0
		for _, v := range d {
			total += eval(v)
		}
		return total
	case map[string]any:
		total := 0.0
		for _, v := range d {
			total += eval(v)
		}
		return total
	}
	return 0
}

func sumNumbers(s string) float64 {
	var document any
	if err := json.Unmarshal([]byte(s), &document); err != nil {
		fmt.Println(err)
	}

	return eval(document)
}
