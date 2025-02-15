package slices2

// Transform transforms each element of slice E1 to E2.
func Transform[E2 any, S1 ~[]E1, E1 any](s S1, convFn func(E1) E2) []E2 {
	if s == nil {
		return nil
	}

	out := make([]E2, 0, len(s))
	for _, v := range s {
		out = append(out, convFn(v))
	}

	return out
}
