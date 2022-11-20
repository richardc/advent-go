package maths

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

func Abs[T constraints.Signed](n T) T {
	if n < 0 {
		return -n
	}
	return n
}

func AbsDiff[T constraints.Signed](a, b T) T {
	if a < b {
		return Abs(b - a)
	}
	return Abs(a - b)
}

func GCD[T constraints.Integer](a, b T) T {
	for b != 0 {
		a, b = b, a%b
	}
	return Max(a, -a)
}
