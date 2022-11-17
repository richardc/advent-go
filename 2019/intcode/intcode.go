package intcode

import (
	"strings"

	"github.com/richardc/advent-go/input"
	"github.com/richardc/advent-go/slices"
)

type Intepreter struct {
	memory []int
	pc     int
	halted bool
}

func NewInterpreter(program string) Intepreter {
	return Intepreter{
		memory: slices.Map(strings.Split(strings.Trim(program, "\n"), ","), input.MustAtoi),
	}
}

func (i *Intepreter) Clone() Intepreter {
	return Intepreter{
		memory: append([]int{}, i.memory...),
		pc:     i.pc,
		halted: i.halted,
	}
}

func (i *Intepreter) Step() {
	switch i.memory[i.pc] {
	case 1:
		i.memory[i.memory[i.pc+3]] = i.memory[i.memory[i.pc+1]] + i.memory[i.memory[i.pc+2]]
		i.pc += 4
	case 2:
		i.memory[i.memory[i.pc+3]] = i.memory[i.memory[i.pc+1]] * i.memory[i.memory[i.pc+2]]
		i.pc += 4
	case 99:
		i.halted = true
	}
}

func (i *Intepreter) Halted() bool {
	return i.halted
}

func (i *Intepreter) Set(index, value int) {
	i.memory[index] = value
}

func (i *Intepreter) Get(index int) int {
	return i.memory[index]
}

func (i *Intepreter) Run() {
	for !i.Halted() {
		i.Step()
	}
}
