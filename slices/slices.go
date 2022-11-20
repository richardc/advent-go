package slices

import (
	"golang.org/x/exp/constraints"
)

func Identity[T any](v T) T {
	return v
}

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

type Group[K any, E any] struct {
	Key   K
	Group []E
}

func GroupBy[E comparable](s []E) []Group[E, E] {
	return GroupByFunc(s, func(i E) E { return i })
}

func GroupByFunc[E any, K comparable](s []E, f func(E) K) []Group[K, E] {
	groups := []Group[K, E]{}
	current := Group[K, E]{f(s[0]), []E{s[0]}}

	for _, v := range s[1:] {
		key := f(v)
		if current.Key != key {
			groups = append(groups, current)
			current = Group[K, E]{key, []E{v}}
		} else {
			current.Group = append(current.Group, v)
		}
	}

	return append(groups, current)
}

func Sum[S []E, E constraints.Integer | constraints.Float](slice S) E {
	var total E
	for _, v := range slice {
		total += v
	}
	return total
}

func Product[E constraints.Integer | constraints.Float](s []E) E {
	if len(s) < 1 {
		return 0
	}
	total := E(1)
	for _, v := range s {
		total *= v
	}
	return total
}

func MinBy[E any, O constraints.Ordered](s []E, f func(E) O) E {
	min := s[0]
	minMapped := f(s[0])
	for _, v := range s[1:] {
		mapped := f(v)
		if mapped < minMapped {
			minMapped = mapped
			min = v
		}
	}
	return min
}

func MaxBy[E any, O constraints.Ordered](s []E, f func(E) O) E {
	max := s[0]
	maxMapped := f(s[0])
	for _, v := range s[1:] {
		mapped := f(v)
		if mapped > maxMapped {
			maxMapped = mapped
			max = v
		}
	}
	return max
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

func CombinationsFunc[E any](s []E, n int, f func([]E)) {
	var comb func([]E, int)
	comb = func(s []E, k int) {
		if n == k {
			f(s[:n])
		} else {
			for i := k; i < len(s); i++ {
				s[k], s[i] = s[i], s[k]
				comb(s, k+1)
				s[k], s[i] = s[i], s[k]
			}
		}
	}

	comb(s, 0)
}

func Combinations[E any](s []E, n int) [][]E {
	result := [][]E{}
	CombinationsFunc(s, n, func(c []E) {
		result = append(result, append([]E{}, c...))
	})
	return result
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

func Powerset[E any](set []E) [][]E {
	powerSet := [][]E{{}}

	for _, es := range set {
		var u [][]E
		for _, er := range powerSet {
			u = append(u, append([]E{es}, er...))
		}
		powerSet = append(powerSet, u...)
	}

	return powerSet
}

func Reverse[E any](s []E) {
	for i, j := 0, len(s)-1; i < j; i, j = i+1, j-1 {
		s[i], s[j] = s[j], s[i]
	}
}

func Counts[E comparable](s []E) map[E]int {
	counts := map[E]int{}
	for _, v := range s {
		counts[v]++
	}
	return counts
}

func Fold[E any, A any](s []E, init A, f func(A, E) A) A {
	acc := init
	for _, v := range s {
		acc = f(acc, v)
	}
	return acc
}

func Range[E constraints.Integer](start, end E) []E {
	r := make([]E, end-start)
	for i := start; i < end; i++ {
		r[i] = i
	}
	return r
}

func Any[E any](s []E, f func(E) bool) bool {
	for _, v := range s {
		if f(v) {
			return true
		}
	}
	return false
}

func All[E any](s []E, f func(E) bool) bool {
	for _, v := range s {
		if !f(v) {
			return false
		}
	}
	return true
}
