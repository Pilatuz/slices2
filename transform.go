package slices2

// Transform transforms each element of slice.
func Transform[S ~[]E, E any, T any](fn func(E) T, s S) []T {
	if s == nil {
		return nil
	}

	out := make([]T, 0, len(s))
	for _, v := range s {
		out = append(out, fn(v))
	}

	return out
}
