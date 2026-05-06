//go:build go1.21

package slices2

import (
	"slices"
)

// Clone makes a copy of a slice.
// Returns nil for nil slice.
func Clone[S ~[]E, E any](s S) S {
	return slices.Clone(s)
}
