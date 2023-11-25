//go:build go1.21

package slices2

import (
	"slices"
)

// ReverseInPlace reverses elements of a slice (in place).
func ReverseInPlace[S ~[]E, E any](s S) {
	slices.Reverse(s)
}
