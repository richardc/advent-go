package day23

import (
	_ "embed"
	"strings"

	"github.com/richardc/advent-go/input"
	"github.com/richardc/advent-go/runner"
	"github.com/richardc/advent-go/slices"
)

//go:embed input.txt
var puzzle string

func init() {
	runner.Register(runner.Solution{
		Day:   23,
		Input: func() any { return slices.Map(input.Lines(puzzle), newInstr) },
		Part1: func(i any) any { return valueOfB(cpu{}, i.([]instr)) },
		Part2: func(i any) any { return valueOfB(cpu{a: 1}, i.([]instr)) },
	})
}

type cpu struct {
	a, b int
	pc   int
}

func reg(c *cpu, r string) *int {
	if r == "a" {
		return &c.a
	}
	return &c.b
}

type instr interface {
	apply(*cpu)
}

type hlf string

func (h hlf) apply(c *cpu) {
	*reg(c, string(h)) /= 2
	c.pc++
}

type tpl string

func (t tpl) apply(c *cpu) {
	*reg(c, string(t)) *= 3
	c.pc++
}

type inc string

func (i inc) apply(c *cpu) {
	*reg(c, string(i))++
	c.pc++
}

type jmp int

func (j jmp) apply(c *cpu) {
	c.pc += int(j)
}

type jie struct {
	r string
	o int
}

func (j jie) apply(c *cpu) {
	if *reg(c, j.r)%2 == 0 {
		c.pc += j.o
	} else {
		c.pc++
	}
}

type jio struct {
	r string
	o int
}

func (j jio) apply(c *cpu) {
	if *reg(c, j.r) == 1 {
		c.pc += j.o
	} else {
		c.pc++
	}
}

func newInstr(s string) instr {
	toks := strings.FieldsFunc(s, func(r rune) bool { return r == ' ' || r == ',' })
	switch toks[0] {
	case "hlf":
		return hlf(toks[1])
	case "tpl":
		return tpl(toks[1])
	case "inc":
		return inc(toks[1])
	case "jmp":
		return jmp(input.MustAtoi(toks[1]))
	case "jie":
		return jie{toks[1], input.MustAtoi(toks[2])}
	case "jio":
		return jio{toks[1], input.MustAtoi(toks[2])}
	}
	panic("Unknown")
}

func valueOfB(c cpu, prog []instr) int {
	for c.pc >= 0 && c.pc < len(prog) {
		prog[c.pc].apply(&c)
	}
	return c.b
}
