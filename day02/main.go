package main

import (
	_ "embed"
	"fmt"
	"math"
	"strconv"
	"strings"

	"golang.org/x/exp/constraints"
	"golang.org/x/exp/slices"
)

type Parcel struct {
	l, w, h int
}

func parcel(s string) Parcel {
	sides := Map(strings.Split(s, "x"), Atoi)
	return Parcel{sides[0], sides[1], sides[2]}
}

func paperNeeded(p Parcel) int {
	smallestSide := math.MaxInt
	smallestSide = Min(smallestSide, p.l*p.w)
	smallestSide = Min(smallestSide, p.w*p.h)
	smallestSide = Min(smallestSide, p.h*p.l)

	return 2*p.l*p.w + 2*p.w*p.h + 2*p.h*p.l + smallestSide
}

func ribbonNeeded(p Parcel) int {
	volume := p.l * p.w * p.h
	edges := []int{p.l, p.w, p.h}
	slices.Sort(edges)
	return volume + edges[0]*2 + edges[1]*2
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
	parcels := Map(lines(input), parcel)
	part1 := Sum(Map(parcels, paperNeeded))
	part2 := Sum(Map(parcels, ribbonNeeded))
	fmt.Println("Part1: ", part1)
	fmt.Println("Part2: ", part2)
}
