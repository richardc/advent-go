package day07

import (
	_ "embed"
	"strconv"
	"strings"

	"github.com/richardc/advent-go/input"
	"github.com/richardc/advent-go/runner"
)

//go:embed input.txt
var puzzle string

func init() {
	runner.Register(runner.Solution{
		Day:   7,
		Input: func() any { return input.Lines(puzzle) },
		Part1: func(i any) any { return valueOfA(i.([]string)) },
	})
}

func value(v string) Eval {
	if val, err := strconv.Atoi(v); err == nil {
		return &Constant{Value(val)}
	} else {
		return &Named{v}
	}
}

func resolveWiring(c []string) map[string]Value {
	wirings := map[string]Eval{}

	for _, conn := range c {
		parts := strings.Split(conn, " -> ")
		target := parts[1]
		args := strings.Split(parts[0], " ")
		switch len(args) {
		case 1:
			wirings[target] = value(args[0])
		case 2: // NOT
			wirings[target] = &Not{value(args[1])}
		case 3:
			left := value(args[0])
			right := value(args[2])
			switch args[1] {
			case "AND":
				wirings[target] = &And{left, right}
			case "OR":
				wirings[target] = &Or{left, right}
			case "LSHIFT":
				wirings[target] = &Lshift{left, right}
			case "RSHIFT":
				wirings[target] = &Rshift{left, right}
			}
		}
	}

	flat := map[string]Value{}
	for k, v := range wirings {
		flat[k] = v.Eval(Bindings{
			table: wirings,
			cache: flat,
		})
	}

	return flat
}

type Value uint16
type Bindings struct {
	cache map[string]Value
	table map[string]Eval
}

type Eval interface {
	Eval(Bindings) Value
}

type Named struct {
	identifier string
}

func (n *Named) Eval(b Bindings) Value {
	if val, ok := b.cache[n.identifier]; ok {
		return val
	}
	res := b.table[n.identifier].Eval(b)
	b.cache[n.identifier] = res
	return res
}

type Constant struct {
	value Value
}

func (c *Constant) Eval(Bindings) Value { return c.value }

type Not struct {
	Op Eval
}

func (n *Not) Eval(b Bindings) Value {
	return ^n.Op.Eval(b)
}

type And struct {
	Left  Eval
	Right Eval
}

func (a *And) Eval(t Bindings) Value {
	return a.Left.Eval(t) & a.Right.Eval(t)
}

type Or struct {
	Left  Eval
	Right Eval
}

func (a *Or) Eval(t Bindings) Value {
	return a.Left.Eval(t) | a.Right.Eval(t)
}

type Lshift struct {
	Left  Eval
	Right Eval
}

func (a *Lshift) Eval(t Bindings) Value {
	return a.Left.Eval(t) << a.Right.Eval(t)
}

type Rshift struct {
	Left  Eval
	Right Eval
}

func (a *Rshift) Eval(t Bindings) Value {
	return a.Left.Eval(t) >> a.Right.Eval(t)
}

func valueOfA(c []string) Value {
	return resolveWiring(c)["a"]
}
