//go:build go1.21

package slices2

import (
	"cmp"
	"slices"
)

// SortInPlace sorts an input slice s of any ordered type in ascending order (in place).
// This sort is not guaranteed to be stable.
func SortInPlace[S ~[]E, E cmp.Ordered](s S) {
	slices.Sort(s)
}

// SortFuncInPlace sorts input slice s of any type in ascending order using the cmp function.
// This sort is not guaranteed to be stable.
func SortFuncInPlace[S ~[]E, E any](s S, cmp func(E, E) int) {
	slices.SortFunc(s, cmp)
}

// Sorted returns new slice with elements of input slice s sorted.
// As opposite to [SortInPlace] this function always allocates new array.
func Sorted[S ~[]E, E cmp.Ordered](s S) S {
	s = Clone(s)
	SortInPlace(s)
	return s
}

// SortedFunc returns new slice with elements of input slice s sorted using the cmp function.
// As opposite to [SortFuncInPlace] this function always allocates new array.
func SortedFunc[S ~[]E, E any](s S, cmp func(E, E) int) S {
	s = Clone(s)
	SortFuncInPlace(s, cmp)
	return s
}
