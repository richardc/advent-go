package input

import (
	"strconv"
	"strings"
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
