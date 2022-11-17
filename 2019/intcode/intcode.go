package intcode

import (
	"strconv"
	"strings"

	"github.com/richardc/advent-go/input"
	"github.com/richardc/advent-go/slices"
)

type Cpu struct {
	memory      []int
	input       []int
	output      []int
	output_func func(int)
	pc          int
	halted      bool
	blocked     bool
}

func NewCpu(program string) Cpu {
	return Cpu{
		memory: slices.Map(strings.Split(strings.Trim(program, "\n"), ","), input.MustAtoi),
	}
}

func (c *Cpu) Clone() Cpu {
	return Cpu{
		memory: append([]int{}, c.memory...),
		pc:     c.pc,
		halted: c.halted,
	}
}

func (c *Cpu) Reset(other *Cpu) {
	c.pc = 0
	c.halted = false
	c.blocked = false
	c.input = []int{}
	c.output = []int{}
	c.memory = append([]int{}, other.memory...)
}

func (c *Cpu) OutputFunc(f func(int)) {
	c.output_func = f
}

func (c *Cpu) Argument(offset int) int {
	op := c.memory[c.pc]
	digits := []byte(strconv.FormatInt(int64(op), 10))
	slices.Reverse(digits)
	var mode byte
	if offset+1 > len(digits)-1 {
		mode = '0'
	} else {
		mode = digits[offset+1]
	}
	switch mode {
	case '0': // position
		return c.memory[c.memory[c.pc+offset]]
	case '1': // immediate
		return c.memory[c.pc+offset]
	}
	return 0
}

func (c *Cpu) Step() {
	op := c.memory[c.pc]
	opcode := op % 100
	switch opcode {
	case 1: // Add
		c.memory[c.memory[c.pc+3]] = c.Argument(1) + c.Argument(2)
		c.pc += 4
	case 2: // Mul
		c.memory[c.memory[c.pc+3]] = c.Argument(1) * c.Argument(2)
		c.pc += 4
	case 3: // Input
		if len(c.input) > 0 {
			c.memory[c.memory[c.pc+1]] = c.input[0]
			c.input = c.input[1:]
			c.pc += 2
		} else {
			c.blocked = true
		}
	case 4: // Output
		if c.output_func != nil {
			c.output_func(c.Argument(1))
		} else {
			c.output = append(c.output, c.Argument(1))
		}
		c.pc += 2
	case 5: // Jump if true
		if c.Argument(1) != 0 {
			c.pc = c.Argument(2)
		} else {
			c.pc += 3
		}
	case 6: // Jump if false
		if c.Argument(1) == 0 {
			c.pc = c.Argument(2)
		} else {
			c.pc += 3
		}
	case 7: // Less than
		if c.Argument(1) < c.Argument(2) {
			c.memory[c.memory[c.pc+3]] = 1
		} else {
			c.memory[c.memory[c.pc+3]] = 0
		}
		c.pc += 4
	case 8: // Equals
		if c.Argument(1) == c.Argument(2) {
			c.memory[c.memory[c.pc+3]] = 1
		} else {
			c.memory[c.memory[c.pc+3]] = 0
		}
		c.pc += 4
	case 99: // Halt
		c.halted = true
	default:
		panic("Unhandled opcode")
	}
}

func (c *Cpu) Halted() bool {
	return c.halted
}

func (c *Cpu) Blocked() bool {
	return c.blocked
}

func (c *Cpu) Set(index, value int) {
	c.memory[index] = value
}

func (c *Cpu) Get(index int) int {
	return c.memory[index]
}

func (c *Cpu) Run() {
	for !c.Blocked() && !c.Halted() {
		c.Step()
	}
}

func (c *Cpu) Input(input []int) {
	c.input = append(c.input, input...)
	c.blocked = false
}

func (c *Cpu) Output() []int {
	return c.output
}
