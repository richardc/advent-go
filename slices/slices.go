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
