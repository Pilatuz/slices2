//go:build go1.21

package slices2

import (
	"slices"
)

// Clone makes copy of a slice.
func Clone[S ~[]E, E any](s S) S {
	return slices.Clone(s)
}
