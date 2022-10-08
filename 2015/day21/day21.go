package day21

import (
	_ "embed"
	"strings"

	"github.com/richardc/advent-go/input"
	"github.com/richardc/advent-go/runner"
	"github.com/richardc/advent-go/slices"
)

//go:embed input.txt
var puzzle string

func init() {
	runner.Register(runner.Solution{
		Day:   21,
		Input: func() any { return newBoss(input.Lines(puzzle)) },
		Part1: func(i any) any { return cheapestWinningLoadout(i.(actor)) },
		Part2: func(i any) any { return expensiveLosingLoadout(i.(actor)) },
	})
}

type buff struct {
	cost   int
	damage int
	armor  int
}

var weapons = []buff{
	{8, 4, 0},
	{10, 5, 0},
	{25, 6, 0},
	{40, 7, 0},
	{74, 8, 0},
}

var armor = []buff{
	{13, 0, 1},
	{31, 0, 2},
	{53, 0, 3},
	{75, 0, 4},
	{102, 0, 5},
}

var rings = []buff{
	{25, 1, 0},
	{50, 2, 0},
	{100, 3, 0},
	{20, 0, 1},
	{40, 0, 2},
	{80, 0, 3},
}

type actor struct {
	health int
	damage int
	armor  int
}

func newBoss(s []string) actor {
	boss := actor{}
	for _, line := range s {
		field, value, _ := strings.Cut(line, ": ")
		switch field {
		case "Hit Points":
			boss.health = input.MustAtoi(value)
		case "Damage":
			boss.damage = input.MustAtoi(value)
		case "Armor":
			boss.armor = input.MustAtoi(value)
		}
	}
	return boss
}

func win(player, boss actor) bool {
	// fmt.Println(player, boss)
	for {
		// fmt.Println(player, boss)
		boss.health -= player.damage - boss.armor
		if boss.health < 1 {
			return true
		}
		player.health -= boss.damage - player.armor
		if player.health < 1 {
			return false
		}
	}
}

func (a actor) with(b buff) actor {
	buffed := a
	buffed.armor += b.armor
	buffed.damage += b.damage
	return buffed
}

type loadout struct {
	player actor
	cost   int
}

func (l loadout) with(b buff) loadout {
	buffed := l
	buffed.player = l.player.with(b)
	buffed.cost += b.cost
	return buffed
}

// https://rosettacode.org/wiki/Cartesian_product_of_two_or_more_lists#Go
func Product[E any](slices ...[]E) [][]E {
	c := 1
	for _, a := range slices {
		c *= len(a)
	}
	if c == 0 {
		return nil
	}
	p := make([][]E, c)
	b := make([]E, c*len(slices))
	n := make([]int, len(slices))
	s := 0
	for i := range p {
		e := s + len(slices)
		pi := b[s:e]
		p[i] = pi
		s = e
		for j, n := range n {
			pi[j] = slices[j][n]
		}
		for j := len(n) - 1; j >= 0; j-- {
			n[j]++
			if n[j] < len(slices[j]) {
				break
			}
			n[j] = 0
		}
	}
	return p
}

func loadouts() []loadout {
	base := loadout{actor{100, 0, 0}, 0}
	buffs := Product(weapons, append(armor, buff{}), append(rings, buff{}), append(rings, buff{}))

	dedup_rings := slices.Filter(buffs, func(doll []buff) bool {
		return doll[2] == buff{} || doll[3] == buff{} || doll[2] != doll[3]
	})

	loadouts := slices.Map(dedup_rings, func(b []buff) loadout {
		l := base
		for _, buff := range b {
			l = l.with(buff)
		}
		return l
	})

	return slices.Unique(loadouts)
}

func cheapestWinningLoadout(boss actor) int {
	options := loadouts()
	winning := slices.Filter(options, func(l loadout) bool { return win(l.player, boss) })
	costs := slices.Map(winning, func(l loadout) int { return l.cost })
	return slices.Min(costs)
}

func expensiveLosingLoadout(boss actor) int {
	options := loadouts()
	losing := slices.Filter(options, func(l loadout) bool { return !win(l.player, boss) })
	costs := slices.Map(losing, func(l loadout) int { return l.cost })
	return slices.Max(costs)
}
