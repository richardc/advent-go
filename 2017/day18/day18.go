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
			Part2: func(any) any { return duet(puzzle) },
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
	snd       func(int)
	rcv       func(Value)
	sleep     string
	program   []Instr
	channel   []int
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
		c.snd(i.Op1.Eval(c))
	case "set":
		c.registers[string(i.Op1.(Register))] = i.Op2.Eval(c)
	case "add":
		c.registers[string(i.Op1.(Register))] += i.Op2.Eval(c)
	case "mul":
		c.registers[string(i.Op1.(Register))] *= i.Op2.Eval(c)
	case "mod":
		c.registers[string(i.Op1.(Register))] %= i.Op2.Eval(c)
	case "rcv":
		c.rcv(i.Op1)
	case "jgz":
		if i.Op1.Eval(c) > 0 {
			jump = i.Op2.Eval(c)
		}
	}
	c.pc += jump
}

func (c Cpu) Runnable() bool {
	return c.pc >= 0 && c.pc <= len(c.program) && !c.Blocked()
}

func (c *Cpu) Run() {
	for c.Runnable() {
		c.Apply(c.program[c.pc])
	}
}

func (c Cpu) Blocked() bool {
	return c.sleep != ""
}

func (c *Cpu) TryWake() {
	if c.Blocked() && len(c.channel) > 0 {
		c.registers[c.sleep] = c.channel[0]
		c.channel = c.channel[1:]
		c.sleep = ""
	}
}

func (c *Cpu) Recv(v Value) {
	register := string(v.(Register))
	if len(c.channel) == 0 {
		c.sleep = register
		return
	}

	c.registers[register] = c.channel[0]
	c.channel = c.channel[1:]
}

func solve(puzzle string) int {
	cpu := NewCpu(puzzle)

	playing := 0
	cpu.snd = func(val int) {
		if val > 0 {
			playing = val
		}
	}
	cpu.rcv = func(v Value) {
		cpu.sleep = string(v.(Register))
	}
	cpu.Run()
	return playing
}

func duet(puzzle string) int {
	cpu0 := NewCpu(puzzle)
	cpu0.registers["p"] = 0
	cpu1 := NewCpu(puzzle)
	cpu1.registers["p"] = 1

	answer := 0
	cpu0.snd = func(val int) {
		cpu1.channel = append(cpu1.channel, val)
	}

	cpu1.snd = func(val int) {
		answer++
		cpu0.channel = append(cpu0.channel, val)
	}

	cpu0.rcv = func(v Value) { cpu0.Recv(v) }
	cpu1.rcv = func(v Value) { cpu1.Recv(v) }

	for cpu0.Runnable() || cpu1.Runnable() {
		cpu0.Run()
		cpu1.Run()
		cpu0.TryWake()
		cpu1.TryWake()
	}

	return answer
}
