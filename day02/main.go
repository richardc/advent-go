package main

import (
	_ "embed"
	"fmt"
	"math"
	"strconv"
	"strings"

	"golang.org/x/exp/constraints"
)

func paperNeeded(parcel string) int {
	sides := Map(strings.Split(parcel, "x"), Atoi)
	l, w, h := sides[0], sides[1], sides[2]

	smallestSide := math.MaxInt
	smallestSide = Min(smallestSide, l*w)
	smallestSide = Min(smallestSide, w*h)
	smallestSide = Min(smallestSide, h*l)

	return 2*l*w + 2*w*h + 2*h*l + smallestSide
}

func Atoi(s string) int {
	val, _ := strconv.Atoi(s)
	return val
}

func lines(all string) []string {
	return strings.Split(strings.Trim(all, " \n"), "\n")
}

func Min[T constraints.Ordered](a, b T) T {
	if a < b {
		return a
	}
	return b
}

func Map[S []E, E any, R any](in S, f func(E) R) []R {
	result := make([]R, 0, len(in))
	for _, v := range in {
		result = append(result, f(v))
	}
	return result
}

func Sum[S []E, E constraints.Integer](slice S) E {
	var total E
	for _, v := range slice {
		total += v
	}
	return total
}

//go:embed input.txt
var input string

func main() {
	part1 := Sum(Map(lines(input), paperNeeded))
	fmt.Println("Part1: ", part1)
}
