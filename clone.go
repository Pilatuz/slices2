package slices2

// Clone makes copy of slice.
func Clone[S ~[]E, E any](s S) S {
	if s == nil {
		return nil
	}

	return append(S{}, s...)
}
