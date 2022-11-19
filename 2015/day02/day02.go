package day02

import (
	_ "embed"
	"math"
	"strconv"
	"strings"

	"golang.org/x/exp/slices"

	mth "github.com/richardc/advent-go/maths"
	"github.com/richardc/advent-go/runner"
	slces "github.com/richardc/advent-go/slices"
)

//go:embed input.txt
var input string

func init() {
	runner.Register(runner.Solution{
		Day:   2,
		Input: func() any { return slces.Map(lines(input), newParcel) },
		Part1: func(parcels any) any { return slces.Sum(slces.Map(parcels.([]Parcel), paperNeeded)) },
		Part2: func(parcels any) any { return slces.Sum(slces.Map(parcels.([]Parcel), ribbonNeeded)) },
	})
}

type Parcel struct {
	l, w, h int
}

func newParcel(s string) Parcel {
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
