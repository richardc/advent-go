package intcode

import (
	"strconv"
	"strings"

	"github.com/richardc/advent-go/input"
	"github.com/richardc/advent-go/slices"
	"golang.org/x/exp/maps"
)

type Cpu struct {
	memory       map[int]int
	input        []int
	output       []int
	outputFunc   func(int)
	inputFunc    func() int
	pc           int
	relativeBase int
	halted       bool
	blocked      bool
}

func New(program string) Cpu {
	memory := map[int]int{}
	for k, v := range slices.Map(strings.Split(strings.Trim(program, "\n"), ","), input.MustAtoi) {
		memory[k] = v
	}
	return Cpu{
		memory: memory,
	}
}

func (c *Cpu) Clone() Cpu {
	memory := map[int]int{}
	for k, v := range c.memory {
		memory[k] = v
	}
	return Cpu{
		memory: memory,
		pc:     c.pc,
		halted: c.halted,
	}
}

func (c *Cpu) Reset(other *Cpu) {
	memory := map[int]int{}
	for k, v := range other.memory {
		memory[k] = v
	}
	c.pc = 0
	c.relativeBase = 0
	c.halted = false
	c.blocked = false
	c.input = []int{}
	c.output = []int{}
	c.memory = memory
}

func (c *Cpu) InputFunc(f func() int) {
	c.inputFunc = f
}

func (c *Cpu) OutputFunc(f func(int)) {
	c.outputFunc = f
}

func (c *Cpu) Argument(offset int) int {
	return c.memory[c.Address(offset)]
}

func (c *Cpu) Address(offset int) int {
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
		return c.memory[c.pc+offset]
	case '1': // immediate
		return c.pc + offset
	case '2': // relative
		return c.memory[c.pc+offset] + c.relativeBase
	}
	return -1
}

func (c *Cpu) Step() {
	op := c.memory[c.pc]
	opcode := op % 100
	switch opcode {
	case 1: // Add
		c.memory[c.Address(3)] = c.Argument(1) + c.Argument(2)
		c.pc += 4
	case 2: // Mul
		c.memory[c.Address(3)] = c.Argument(1) * c.Argument(2)
		c.pc += 4
	case 3: // Input
		if c.inputFunc != nil {
			c.memory[c.Address(1)] = c.inputFunc()
			c.pc += 2
		} else {
			if len(c.input) > 0 {
				c.memory[c.Address(1)] = c.input[0]
				c.input = c.input[1:]
				c.pc += 2
			} else {
				c.blocked = true
			}
		}
	case 4: // Output
		if c.outputFunc != nil {
			c.outputFunc(c.Argument(1))
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
			c.memory[c.Address(3)] = 1
		} else {
			c.memory[c.Address(3)] = 0
		}
		c.pc += 4
	case 8: // Equals
		if c.Argument(1) == c.Argument(2) {
			c.memory[c.Address(3)] = 1
		} else {
			c.memory[c.Address(3)] = 0
		}
		c.pc += 4
	case 9: // Adjust relative base
		c.relativeBase += c.Argument(1)
		c.pc += 2
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

func (c *Cpu) Memory() []int {
	max := slices.Max(maps.Keys(c.memory))
	result := make([]int, max+1)
	for i := 0; i <= max; i++ {
		result[i] = c.memory[i]
	}
	return result
}
