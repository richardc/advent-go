package day21

import (
	_ "embed"
	"fmt"
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
			Day:   21,
			Input: func() any { return NewRules(puzzle) },
			Part1: func(i any) any { return countPixels(i.(Rules), 5, false) },
			Part2: func(i any) any { return countPixels(i.(Rules), 18, false) },
		},
	)
}

type Rules map[string]Game

func NewRules(s string) Rules {
	rules := Rules{}
	for _, l := range input.Lines(s) {
		from, to, _ := strings.Cut(l, " => ")
		for _, game := range NewGameFrom(from).Rotations() {
			rules[game.String()] = NewGameFrom(to)
		}
	}
	return rules
}

type Game struct {
	Width int
	Data  []byte
}

func NewGame() Game {
	return NewGameFrom(".#./..#/###")
}

func NewGameFrom(s string) Game {
	return Game{
		Width: strings.Index(s, "/"),
		Data:  slices.Filter([]byte(s), func(b byte) bool { return b != byte('/') }),
	}
}

func (g Game) Clone() Game {
	copy := g
	// need to make a deep copy of Data
	copy.Data = nil
	copy.Data = append(copy.Data, g.Data...)
	return copy
}

func (g Game) Mirror() Game {
	copy := g.Clone()
	for i := 0; i < g.Width; i++ {
		j := i + (g.Width-1)*g.Width
		copy.Data[i], copy.Data[j] = copy.Data[j], copy.Data[i]
	}
	return copy
}

func (g Game) Rotate() Game {
	// Simpler than doing the software was to draw digrams for the two sizes
	copy := g.Clone()
	if g.Width == 2 {
		copy.Data[0] = g.Data[2]
		copy.Data[1] = g.Data[0]
		copy.Data[2] = g.Data[3]
		copy.Data[3] = g.Data[1]
	} else {
		copy.Data[0] = g.Data[6]
		copy.Data[1] = g.Data[3]
		copy.Data[2] = g.Data[0]
		copy.Data[3] = g.Data[7]
		copy.Data[4] = g.Data[4]
		copy.Data[5] = g.Data[1]
		copy.Data[6] = g.Data[8]
		copy.Data[7] = g.Data[5]
		copy.Data[8] = g.Data[2]
	}
	return copy
}

func (g Game) Rotations() []Game {
	return []Game{
		g,
		g.Rotate(),
		g.Rotate().Rotate(),
		g.Rotate().Rotate().Rotate(),
		g.Mirror(),
		g.Mirror().Rotate(),
		g.Mirror().Rotate().Rotate(),
		g.Mirror().Rotate().Rotate().Rotate(),
	}
}

func (g Game) Split(side int) [][]Game {
	res := [][]Game{}
	for row := 0; row < g.Width; row += side {
		grow := []Game{}
		for col := 0; col < g.Width; col += side {
			extracted :=
				strings.Join(slices.Map(slices.Range(0, side), func(i int) string {
					start := g.Width*(row+i) + col
					// fmt.Printf("row %d col %d i %d start %d\n", row, col, i, start)
					return string(g.Data[start : start+side])
				}), "/")
			// fmt.Println("extracted", extracted)
			grow = append(grow, NewGameFrom(extracted))
		}
		res = append(res, grow)
	}
	return res
}

func (g *Game) Step(rules Rules) {
	var chunk int
	switch {
	case g.Width%2 == 0:
		chunk = 2
	case g.Width%3 == 0:
		chunk = 3
	}

	newWidth := g.Width + g.Width/chunk
	next := make([]byte, newWidth*newWidth)
	parts := g.Split(chunk)
	for row, line := range parts {
		for col, game := range line {
			nextGame := rules[game.String()]
			// fmt.Println("from", game, "to", nextGame, row, col, nextGame.Width, newWidth, g.Width)
			for i, b := range nextGame.Data {
				x := i % nextGame.Width
				y := i / nextGame.Width
				into := row*newWidth*nextGame.Width + col*nextGame.Width + y*newWidth + x
				// fmt.Println("row", row, "col", col, "i", i, "x", x, "y", y, "into", into)
				next[into] = b
			}
		}
	}
	g.Width = newWidth
	g.Data = next
}

func (g Game) String() string {
	result := ""
	for y := 0; y < len(g.Data)/g.Width; y++ {
		result += string(g.Data[y*g.Width:(y+1)*g.Width]) + "\n"
	}
	return result
}

func (g Game) Lines() []string {
	result := []string{}
	for y := 0; y < len(g.Data)/g.Width; y++ {
		result = append(result, string(g.Data[y*g.Width:(y+1)*g.Width]))
	}
	return result
}

func (g Game) CountLit() int {
	return len(slices.Filter(g.Data, func(b byte) bool { return b == byte('#') }))
}

func countPixels(r Rules, iter int, draw bool) int {
	game := NewGame()
	if draw {
		fmt.Println(game)
	}
	for i := 0; i < iter; i++ {
		game.Step(r)
		if draw {
			fmt.Println(game)
		}
	}
	return game.CountLit()
}
