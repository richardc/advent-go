package day14

import (
	_ "embed"
	"fmt"
	"strings"

	"github.com/richardc/advent-go/2017/knothash"
	"github.com/richardc/advent-go/runner"
	"github.com/richardc/advent-go/slices"
)

//go:embed "input.txt"
var puzzle string

func init() {
	runner.Register(
		runner.Solution{
			Year:  2017,
			Day:   14,
			Part1: func(any) any { return countSet(strings.TrimSpace(puzzle)) },
			Part2: func(any) any { return countIslands(strings.TrimSpace(puzzle)) },
		},
	)
}

func expandMap(puzzle string) [128][128]bool {
	result := [128][128]bool{}
	for row := 0; row < 128; row++ {
		hash := knothash.Hash([]byte(fmt.Sprintf("%s-%d", puzzle, row)))
		bits := strings.Join(slices.Map(hash[:], func(b byte) string { return fmt.Sprintf("%08b", b) }), "")
		for col := 0; col < 128; col++ {
			result[row][col] = bits[col] == '1'
		}
	}
	return result
}

// destrucitively turn off cells
func markIsland(memory *[128][128]bool, row, col int) {
	if !memory[row][col] {
		return
	}
	memory[row][col] = false
	if row > 0 {
		markIsland(memory, row-1, col)
	}
	if row < 127 {
		markIsland(memory, row+1, col)
	}
	if col > 0 {
		markIsland(memory, row, col-1)
	}
	if col < 127 {
		markIsland(memory, row, col+1)
	}
}

func countIslands(puzzle string) int {
	islands := 0
	memory := expandMap(puzzle)
	for row := range memory {
		for col := range memory[row] {
			if memory[row][col] {
				islands++
				markIsland(&memory, row, col)
			}
		}
	}
	return islands
}

func countSet(puzzle string) int {
	total := 0
	for row := 0; row < 128; row++ {
		hash := knothash.Hash([]byte(fmt.Sprintf("%s-%d", puzzle, row)))
		bits := strings.Join(slices.Map(hash[:], func(b byte) string { return fmt.Sprintf("%08b", b) }), "")
		total += len(slices.Filter([]byte(bits), func(c byte) bool { return c == byte('1') }))
	}
	return total
}
