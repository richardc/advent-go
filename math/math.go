package math

import "golang.org/x/exp/constraints"

func Min[T constraints.Ordered](a, b T) T {
	if a < b {
		return a
	}
	return b
}

func Max[T constraints.Ordered](a, b T) T {
	if a > b {
		return a
	}
	return b
}

func Signum[T constraints.Signed](n T) T {
	switch {
	case n < 0:
		return -1
	case n > 0:
		return 1
	}
	return 0
}
