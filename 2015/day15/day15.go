package day15

import (
	_ "embed"
	"strings"

	"github.com/richardc/advent-go/input"
	"github.com/richardc/advent-go/maths"
	"github.com/richardc/advent-go/runner"
	"github.com/richardc/advent-go/slices"
)

//go:embed input.txt
var puzzle string

func init() {
	runner.Register(runner.Solution{
		Day:   15,
		Input: func() any { return slices.Map(input.Lines(puzzle), newIngredient) },
		Part1: func(i any) any { max, _ := bestScore4(i.([]ingredient)); return max },
		Part2: func(i any) any { _, cal := bestScore4(i.([]ingredient)); return cal },
	})
}

type ingredient struct {
	name       string
	capacity   int
	durability int
	flavour    int
	texture    int
	calories   int
}

func newIngredient(s string) ingredient {
	// Butterscotch: capacity -1, durability -2, flavor 6, texture 3, calories 8
	toks := strings.FieldsFunc(s, func(r rune) bool { return r == ' ' || r == ',' })
	return ingredient{
		name:       toks[0],
		capacity:   input.MustAtoi(toks[2]),
		durability: input.MustAtoi(toks[4]),
		flavour:    input.MustAtoi(toks[6]),
		texture:    input.MustAtoi(toks[8]),
		calories:   input.MustAtoi(toks[10]),
	}
}

func cookieScore(recipe map[ingredient]int) int {
	capacity := 0
	durability := 0
	flavour := 0
	texture := 0
	for k, v := range recipe {
		capacity += k.capacity * v
		durability += k.durability * v
		flavour += k.flavour * v
		texture += k.texture * v
	}

	capacity = maths.Max(capacity, 0)
	durability = maths.Max(durability, 0)
	flavour = maths.Max(flavour, 0)
	texture = maths.Max(texture, 0)

	return capacity * durability * flavour * texture
}

func calories(recipe map[ingredient]int) int {
	calories := 0
	for k, v := range recipe {
		calories += k.calories * v
	}
	return calories
}

func bestScore4(ingredients []ingredient) (max, calory_max int) {
	const spoons int = 100
	for one := 0; one <= spoons; one++ {
		for two := 0; two <= spoons-one; two++ {
			for three := 0; three <= spoons-(one+two); three++ {
				four := spoons - (one + two + three)
				recipe := map[ingredient]int{
					ingredients[0]: one,
					ingredients[1]: two,
					ingredients[2]: three,
					ingredients[3]: four,
				}

				score := cookieScore(recipe)
				max = maths.Max(max, score)

				if calories(recipe) == 500 {
					calory_max = maths.Max(calory_max, score)
				}
			}
		}
	}

	return
}

func bestScore2(ingredients []ingredient) int {
	max := 0
	const spoons int = 100
	for one := 0; one <= spoons; one++ {
		two := spoons - one
		recipe := map[ingredient]int{
			ingredients[0]: one,
			ingredients[1]: two,
		}

		max = maths.Max(max, cookieScore(recipe))
	}

	return max
}
