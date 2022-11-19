package day25

import (
	_ "embed"
	"strings"

	"github.com/richardc/advent-go/input"
	"github.com/richardc/advent-go/maths"
	"github.com/richardc/advent-go/runner"
)

//go:embed "input.txt"
var puzzle string

func init() {
	runner.Register(
		runner.Solution{
			Year:  2017,
			Day:   25,
			Part1: func(any) any { return checksum(puzzle) },
		},
	)
}

type StateID byte

type Action struct {
	Write byte
	Move  int
	Next  StateID
}

type State struct {
	When0 Action
	When1 Action
}

type Machine struct {
	State    StateID
	States   []State
	Tape     map[int]struct{}
	Position int
	Steps    int
}

func parseDirection(s string) int {
	if s == "right." {
		return 1
	}
	return -1
}

func NewMachine(s string) Machine {
	machine := Machine{Tape: map[int]struct{}{}, States: []State{}}
	lines := input.Lines(s)
	machine.State = StateID(lines[0][len(lines[0])-2] - byte('A'))
	machine.Steps = input.MustAtoi(strings.Fields(lines[1])[5])

	lines = lines[3:]
	for len(lines) > 1 {
		machine.States = append(machine.States, State{
			When0: Action{
				Write: lines[2][len(lines[2])-2] - byte('0'),
				Move:  parseDirection(strings.Fields(lines[3])[6]),
				Next:  StateID(lines[4][len(lines[4])-2] - byte('A')),
			},
			When1: Action{
				Write: lines[6][len(lines[6])-2] - byte('0'),
				Move:  parseDirection(strings.Fields(lines[7])[6]),
				Next:  StateID(lines[8][len(lines[8])-2] - byte('A')),
			},
		})
		lines = lines[maths.Min(10, len(lines)):]
	}
	return machine
}

func (m *Machine) Step() {
	state := m.States[m.State]
	var action Action
	if _, ok := m.Tape[m.Position]; ok {
		action = state.When1
	} else {
		action = state.When0
	}
	if action.Write == 0 {
		delete(m.Tape, m.Position)
	} else {
		m.Tape[m.Position] = struct{}{}
	}
	m.Position += action.Move
	m.State = action.Next
}

func (m Machine) Checksum() int {
	return len(m.Tape)
}

func checksum(puzzle string) int {
	machine := NewMachine(puzzle)
	for i := 0; i < machine.Steps; i++ {
		machine.Step()
	}
	return machine.Checksum()
}
