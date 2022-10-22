package day23

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
			Part1: func(any) any { return debugMulCalled(puzzle) },
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
	program   []Instr
	mulCalled int
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
	case "set":
		c.registers[string(i.Op1.(Register))] = i.Op2.Eval(c)
	case "sub":
		c.registers[string(i.Op1.(Register))] -= i.Op2.Eval(c)
	case "mul":
		c.registers[string(i.Op1.(Register))] *= i.Op2.Eval(c)
		c.mulCalled++
	case "jnz":
		if i.Op1.Eval(c) != 0 {
			jump = i.Op2.Eval(c)
		}
	default:
		panic(i.Op)
	}
	c.pc += jump
}

func (c Cpu) Runnable() bool {
	return c.pc >= 0 && c.pc < len(c.program)
}

func (c *Cpu) Run() {
	for c.Runnable() {
		c.Apply(c.program[c.pc])
	}
}

func debugMulCalled(puzzle string) int {
	cpu := NewCpu(puzzle)

	cpu.Run()
	return cpu.mulCalled
}
