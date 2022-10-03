package runner

import (
	"flag"
	"fmt"
	"log"
	"os"
	"runtime/pprof"

	"golang.org/x/exp/slices"
)

type Solver func(any) any

type Solution struct {
	Day   int
	Input func() any
	Part1 Solver
	Part2 Solver
}

func (s Solution) Run() {
	fmt.Println("Day ", s.Day)
	var input any
	if s.Input != nil {
		input = s.Input()
	}
	if s.Part1 != nil {
		fmt.Println("    Part 1", s.Part1(input))
	}
	if s.Part2 != nil {
		fmt.Println("    Part 2", s.Part2(input))
	}
}

var solutions []Solution

func Register(s Solution) {
	solutions = append(solutions, s)
}

func Run() {
	cpuprofile := flag.String("cpuprofile", "", "PATH")
	runall := flag.Bool("all", false, "Runs all options")
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

	slices.SortFunc(solutions, func(s1, s2 Solution) bool { return s1.Day < s2.Day })

	if runall != nil && *runall {
		for _, s := range solutions {
			s.Run()
		}
	} else {
		solutions[len(solutions)-1].Run()
	}
}
