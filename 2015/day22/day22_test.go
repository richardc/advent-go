package day22

import (
	"fmt"
	"testing"
)

func TestCheapestWin(t *testing.T) {
	var tests = []struct {
		bossHealth int
		expected   int
	}{
		{13, 173 + 53},
		{14, 229 + 113 + 73 + 173 + 53},
	}

	for _, testcase := range tests {
		game := game{
			b: boss{
				health: testcase.bossHealth,
				damage: 8,
			},
			w: wizard{
				health: 10,
				mana:   250,
			},
		}
		testname := fmt.Sprintf("boss: %d", testcase.bossHealth)
		t.Run(testname, func(t *testing.T) {
			ans := cheapestWin(game)
			if ans != testcase.expected {
				t.Errorf("got %v, want %v", ans, testcase.expected)
			}
		})
	}
}
