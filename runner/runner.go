package runner

import (
	"flag"
	"fmt"
	"log"
	"os"
	"runtime/pprof"
	"time"

	slcs "github.com/richardc/advent-go/slices"
	"golang.org/x/exp/slices"
)

type Solver func(any) any

type Solution struct {
	Year  int
	Day   int
	Input func() any
	Part1 Solver
	Part2 Solver
}

func (s Solution) Run() {
	fmt.Println("Day", s.Day)
	var input any
	if s.Input != nil {
		start := time.Now()
		input = s.Input()
		duration := time.Since(start)
		fmt.Println("    Input ", duration)
	}
	if s.Part1 != nil {
		start := time.Now()
		result := s.Part1(input)
		duration := time.Since(start)
		fmt.Println("    Part 1", result)
		fmt.Println("          ", duration)
	}
	if s.Part2 != nil {
		start := time.Now()
		result := s.Part2(input)
		duration := time.Since(start)
		fmt.Println("    Part 2", result)
		fmt.Println("          ", duration)
	}
	fmt.Println()
}

var solutions []Solution

func Register(s Solution) {
	solutions = append(solutions, s)
}

func Run() {
	cpuprofile := flag.String("cpuprofile", "", "PATH")
	runall := flag.Bool("all", false, "Runs all days solutions")
	day := flag.Int("day", 0, "Day to run")
	year := flag.Int("year", 0, "Year to run")
	flag.Parse()

	if cpuprofile != nil && *cpuprofile != "" {
		f, err := os.Create(*cpuprofile)
		if err != nil {
			log.Fatal(err)
		}
		err = pprof.StartCPUProfile(f)
		if err != nil {
			log.Fatal(err)
		}
		defer pprof.StopCPUProfile()
	}

	slices.SortFunc(solutions, func(s1, s2 Solution) bool { return s1.Year < s2.Year && s1.Day < s2.Day })

	filtered := false
	if year != nil && *year != 0 {
		solutions = slcs.Filter(solutions, func(s Solution) bool { return s.Year == *year })
		filtered = true
	}

	if day != nil && *day != 0 {
		solutions = slcs.Filter(solutions, func(s Solution) bool { return s.Day == *day })
		filtered = true
	}

	if (runall != nil && *runall) || filtered {
		for _, s := range solutions {
			s.Run()
		}
	} else {
		solutions[len(solutions)-1].Run()
	}
}
