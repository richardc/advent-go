package day22

import (
	_ "embed"

	"github.com/richardc/advent-go/input"
	"github.com/richardc/advent-go/math"
	"github.com/richardc/advent-go/runner"
)

//go:embed input.txt
var puzzle string

func init() {
	runner.Register(runner.Solution{
		Day:   22,
		Input: func() any { return newBoss(input.Lines(puzzle)) },
		Part1: func(i any) any { return cheapestWin(game{w: wizard{health: 50, mana: 500}, b: i.(boss)}) },
	})
}

type game struct {
	b boss
	w wizard
}

type boss struct {
	health int
	damage int
}

func newBoss(s []string) boss {
	return boss{}
}

type wizard struct {
	health   int
	mana     int
	shield   int
	poison   int
	recharge int
	spent    int
}

func (w *wizard) spend(m int) {
	w.mana -= m
	w.spent += m
}

func (g game) missile() game {
	g.b.health -= 4
	g.w.spend(53)
	return g
}

func (g game) drain() game {
	g.b.health -= 2
	g.w.health += 2
	g.w.spend(73)
	return g
}

func (g game) shield() game {
	if g.w.shield != 0 {
		return game{}
	}
	g.w.shield = 6
	g.w.spend(113)
	return g
}

func (g game) poison() game {
	if g.w.poison != 0 {
		return game{}
	}
	g.w.poison = 6
	g.w.spend(173)
	return g
}

func (g game) recharge() game {
	if g.w.recharge != 0 {
		return game{}
	}
	g.w.recharge = 5
	g.w.spend(229)
	return g
}

func (g game) tick() game {
	if g.w.poison > 0 {
		g.b.health -= 3
	}
	if g.w.recharge > 0 {
		g.w.mana += 101
	}
	g.w.shield -= math.Signum(g.w.shield)
	g.w.poison -= math.Signum(g.w.poison)
	g.w.recharge -= math.Signum(g.w.recharge)
	return g
}

func (g game) moves() []game {
	return nil
}

func cheapestWin(g game) int {
	return 0
}
