package day22

import (
	"container/heap"
	_ "embed"
	"strings"

	"github.com/richardc/advent-go/input"
	"github.com/richardc/advent-go/math"
	"github.com/richardc/advent-go/runner"
	"github.com/richardc/advent-go/slices"
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
	boss := boss{}
	for _, line := range s {
		field, value, _ := strings.Cut(line, ": ")
		switch field {
		case "Hit Points":
			boss.health = input.MustAtoi(value)
		case "Damage":
			boss.damage = input.MustAtoi(value)
		}
	}
	return boss
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
	return g.bossmove()
}

func (g game) drain() game {
	g.b.health -= 2
	g.w.health += 2
	g.w.spend(73)
	return g.bossmove()
}

func (g game) shield() game {
	if g.w.shield != 0 {
		return game{}
	}
	g.w.shield = 6
	g.w.spend(113)
	return g.bossmove()
}

func (g game) poison() game {
	if g.w.poison != 0 {
		return game{}
	}
	g.w.poison = 6
	g.w.spend(173)
	return g.bossmove()
}

func (g game) recharge() game {
	if g.w.recharge != 0 {
		return game{}
	}
	g.w.recharge = 5
	g.w.spend(229)
	return g.bossmove()
}

func (g *game) tick() game {
	if g.w.poison > 0 {
		g.b.health -= 3
	}
	if g.w.recharge > 0 {
		g.w.mana += 101
	}
	g.w.shield -= math.Signum(g.w.shield)
	g.w.poison -= math.Signum(g.w.poison)
	g.w.recharge -= math.Signum(g.w.recharge)
	return *g
}

func (g game) bossmove() game {
	g.tick()
	if g.b.health < 1 {
		return g
	}
	if g.w.shield > 0 {
		g.w.health -= math.Max(0, g.b.damage-7)
	} else {
		g.w.health -= g.b.damage
	}
	return g
}

func (g game) win() bool {
	return g.b.health < 1
}

func (g game) legal() bool {
	return g.w.health > 0 && g.w.mana > 0
}

func (g game) moves() []game {
	g.tick()
	moves := []game{
		g.drain(),
		g.missile(),
		g.shield(),
		g.poison(),
		g.recharge(),
	}

	return slices.Filter(moves, func(g game) bool { return g.legal() })
}

func (g game) state() game {
	g.w.spent = 0
	return g
}

func (g game) cost() int {
	return g.w.spent
}

type queue []*game

func (q queue) Len() int           { return len(q) }
func (q queue) Less(i, j int) bool { return q[i].w.spent < q[j].w.spent }
func (q queue) Swap(i, j int)      { q[i], q[j] = q[j], q[i] }
func (q *queue) Push(x any) {
	*q = append(*q, x.(*game))
}
func (q *queue) Pop() any {
	old := *q
	n := len(old)
	item := old[n-1]
	old[n-1] = nil
	*q = old[:n-1]
	return item
}

func cheapestWin(g game) int {
	q := queue{&g}
	heap.Init(&q)
	costs := map[game]int{g.state(): 0}

	for q.Len() > 0 {
		curr := heap.Pop(&q).(*game)

		if curr.win() {
			return curr.cost()
		}

		if cost, ok := costs[curr.state()]; ok && cost < curr.cost() {
			continue
		}

		for _, next := range curr.moves() {
			if cost, ok := costs[next.state()]; ok && cost < next.cost() {
				continue
			}
			costs[next.state()] = next.cost()
			next := next
			heap.Push(&q, &next)
		}
	}
	return 0
}
