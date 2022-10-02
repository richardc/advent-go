package main

import (
	_ "embed"
	"fmt"
	"math"
	"strconv"
	"strings"

	"golang.org/x/exp/slices"

	mth "advent2015/pkg/math"
	slces "advent2015/pkg/slices"
)

type Parcel struct {
	l, w, h int
}

func parcel(s string) Parcel {
	sides := slces.Map(strings.Split(s, "x"), Atoi)
	return Parcel{sides[0], sides[1], sides[2]}
}

func paperNeeded(p Parcel) int {
	smallestSide := math.MaxInt
	smallestSide = mth.Min(smallestSide, p.l*p.w)
	smallestSide = mth.Min(smallestSide, p.w*p.h)
	smallestSide = mth.Min(smallestSide, p.h*p.l)

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

//go:embed input.txt
var input string

func main() {
	parcels := slces.Map(lines(input), parcel)
	part1 := slces.Sum(slces.Map(parcels, paperNeeded))
	part2 := slces.Sum(slces.Map(parcels, ribbonNeeded))
	fmt.Println("Part1: ", part1)
	fmt.Println("Part2: ", part2)
}
