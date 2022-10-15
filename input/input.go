package input

import (
	"strconv"
	"strings"

	"github.com/richardc/advent-go/slices"
)

func MustAtoi(s string) int {
	v, err := strconv.Atoi(s)
	if err != nil {
		panic(err)
	}
	return v
}

func Lines(s string) []string {
	return strings.Split(strings.Trim(s, "\n"), "\n")
}

func Sheet[E any](s string, f func(string) E) [][]E {
	return slices.Map(Lines(s), func(l string) []E { return slices.Map(strings.Fields(l), f) })
}
