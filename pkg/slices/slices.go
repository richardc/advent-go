package slices

import (
	"golang.org/x/exp/constraints"
)

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
