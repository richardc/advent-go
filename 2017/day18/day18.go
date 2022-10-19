package day18

import (
	_ "embed"
	"strconv"
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
			Day:   18,
			Part1: func(any) any { return solve(puzzle) },
		},
	)
}

type Value interface {
	Eval(*Cpu) int
}

type Literal int

func (l Literal) Eval(*Cpu) int {
	return int(l)
}

type Register string

func (r Register) Eval(c *Cpu) int {
	return c.registers[string(r)]
}

func NewValue(s string) Value {
	if number, err := strconv.Atoi(s); err == nil {
		return Literal(number)
	}
	return Register(s)
}

type Instr struct {
	Op  string
	Op1 Value
	Op2 Value
}

func NewInstr(s string) Instr {
	toks := strings.Fields(s)
	instr := Instr{
		Op:  toks[0],
		Op1: NewValue(toks[1]),
	}
	if len(toks) == 3 {
		instr.Op2 = NewValue(toks[2])
	}
	return instr
}

type Cpu struct {
	registers map[string]int
	pc        int
	playing   int
	rcv       bool
	program   []Instr
}

func NewCpu(s string) Cpu {
	return Cpu{
		registers: map[string]int{},
		program:   slices.Map(input.Lines(s), NewInstr),
	}
}

func (c *Cpu) Apply(i Instr) {
	jump := 1
	switch i.Op {
	case "snd":
		c.playing = i.Op1.Eval(c)
	case "set":
		c.registers[string(i.Op1.(Register))] = i.Op2.Eval(c)
	case "add":
		c.registers[string(i.Op1.(Register))] += i.Op2.Eval(c)
	case "mul":
		c.registers[string(i.Op1.(Register))] *= i.Op2.Eval(c)
	case "mod":
		c.registers[string(i.Op1.(Register))] %= i.Op2.Eval(c)
	case "rcv":
		if i.Op1.Eval(c) != 0 {
			c.rcv = true
		}
	case "jgz":
		if i.Op1.Eval(c) > 0 {
			jump = i.Op2.Eval(c)
		}
	}
	c.pc += jump
}

func (c *Cpu) Run() {
	for c.pc >= 0 && c.pc <= len(c.program) {
		c.Apply(c.program[c.pc])
	}
}

func (c *Cpu) RunTillRcv() {
	for c.pc >= 0 && c.pc <= len(c.program) && !c.rcv {
		c.Apply(c.program[c.pc])
	}
}

func solve(puzzle string) int {
	cpu := NewCpu(puzzle)
	cpu.RunTillRcv()
	return cpu.playing
}
