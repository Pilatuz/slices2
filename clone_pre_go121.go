//go:build !go1.21

package slices2

// Clone makes a copy of a slice.
// Returns nil for nil slice.
func Clone[S ~[]E, E any](s S) S {
	if s == nil {
		return nil
	}

	return append(S{}, s...)
}
