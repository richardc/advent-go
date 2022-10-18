package day10

import (
	_ "embed"
	"fmt"
	"strings"

	"github.com/richardc/advent-go/2017/knothash"
	"github.com/richardc/advent-go/input"
	"github.com/richardc/advent-go/runner"
	"github.com/richardc/advent-go/slices"
)

//go:embed "input.txt"
var puzzle string

func init() {
	runner.Register(
		runner.Solution{
			Year: 2017,
			Day:  10,
			Part1: func(any) any {
				return knotHashOne(256, slices.Map(strings.Split(strings.TrimSpace(puzzle), ","), input.MustAtoi))
			},
			Part2: func(any) any {
				return knotHash(strings.TrimSpace(puzzle))
			},
		},
	)
}

func knotHashOne(size int, lengths []int) int {
	data := make([]int, size)
	for i := range data {
		data[i] = i
	}

	cur := 0
	skip := 0
	for _, length := range lengths {
		for i, j := cur, cur+length-1; i < j; i, j = i+1, j-1 {
			data[i%len(data)], data[j%len(data)] = data[j%len(data)], data[i%len(data)]
		}

		cur = (cur + length + skip) % len(data)
		skip++
	}

	return data[0] * data[1]
}

func knotHash(s string) string {
	dense := knothash.Hash([]byte(s))
	return fmt.Sprintf("%x", dense)
}
