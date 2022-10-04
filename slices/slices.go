package slices

import (
	"golang.org/x/exp/constraints"
)

func Filter[S ~[]E, E any](s S, p func(E) bool) S {
	result := make(S, 0)
	for _, v := range s {
		if p(v) {
			result = append(result, v)
		}
	}
	return result
}

func Map[S []E, E any, R any](in S, f func(E) R) []R {
	result := make([]R, 0, len(in))
	for _, v := range in {
		result = append(result, f(v))
	}
	return result
}

func Sum[S []E, E constraints.Integer | constraints.Float](slice S) E {
	var total E
	for _, v := range slice {
		total += v
	}
	return total
}

func Min[E constraints.Ordered](s []E) E {
	min := s[0]
	for _, v := range s {
		if v < min {
			min = v
		}
	}
	return min
}

func Max[E constraints.Ordered](s []E) E {
	max := s[0]
	for _, v := range s {
		if v > max {
			max = v
		}
	}
	return max
}

func MinMax[E constraints.Ordered](s []E) (E, E) {
	min := s[0]
	max := s[0]
	for _, v := range s {
		if v < min {
			min = v
		}
		if v > max {
			max = v
		}
	}
	return min, max
}

func Unique[S ~[]E, E comparable](s S) S {
	new := make(S, 0)
	seen := make(map[E]struct{})
	for _, v := range s {
		if _, ok := seen[v]; !ok {
			new = append(new, v)
			seen[v] = struct{}{}
		}
	}
	return new
}

func Permutations[E any](s []E) [][]E {
	result := [][]E{}

	var perm func([]E, int)
	perm = func(s []E, k int) {
		if k == len(s) {
			result = append(result, append([]E{}, s...))
		} else {
			for i := k; i < len(s); i++ {
				s[k], s[i] = s[i], s[k]
				perm(s, k+1)
				s[k], s[i] = s[i], s[k]
			}
		}
	}

	perm(s, 0)
	return result
}
