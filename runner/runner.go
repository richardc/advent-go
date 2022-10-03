package runner

import (
	"fmt"
	"log"
	"os"

	"github.com/mitchellh/cli"
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
	c := cli.NewCLI("advent", "0.0.0")
	c.Args = os.Args[1:]
	c.Commands = map[string]cli.CommandFactory{
		"all":    allCommandFactory,
		"latest": latestCommandFactory,
		"":       latestCommandFactory,
	}

	exitStatus, err := c.Run()
	if err != nil {
		log.Println(err)
	}

	os.Exit(exitStatus)
}
func allCommandFactory() (cli.Command, error) {
	return &allCommand{}, nil
}

type allCommand struct{}

func (c *allCommand) Synopsis() string { return "Runs all days solutions" }
func (c *allCommand) Help() string     { return "" }
func (c *allCommand) Run([]string) int {
	slices.SortFunc(solutions, func(s1, s2 Solution) bool { return s1.Day < s2.Day })

	for _, s := range solutions {
		s.Run()
	}

	return 0
}

func latestCommandFactory() (cli.Command, error) {
	return &latestCommand{}, nil
}

type latestCommand struct{}

func (c *latestCommand) Synopsis() string { return "Runs latest days solution" }
func (c *latestCommand) Help() string     { return "" }
func (c *latestCommand) Run([]string) int {
	slices.SortFunc(solutions, func(s1, s2 Solution) bool { return s1.Day < s2.Day })

	solutions[len(solutions)-1].Run()

	return 0
}
