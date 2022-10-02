package input

import "strings"

func Lines(s string) []string {
	return strings.Split(strings.Trim(s, "\n"), "\n")
}
