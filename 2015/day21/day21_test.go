package day21

import (
	"testing"
)

func TestWin(t *testing.T) {
	player := actor{health: 8, damage: 5, armor: 5}
	boss := actor{health: 12, damage: 7, armor: 2}

	expect := true
	got := win(player, boss)
	if got != expect {
		t.Errorf("got %v, want %v", got, expect)
	}
}
