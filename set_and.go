package slices2

// SetAnd returns the intersection between two sets.
// I.e. elements presented in both slices.
// Order is preserved from s2.
func SetAnd[S ~[]E, E comparable](s1 S, s2 S) S {
	return SetAndBy(s1, s2, me[E])
}

// SetAndBy returns the intersection between two sets by custom key.
// I.e. elements presented in both slices.
// The byFn function extracts a comparison key from each element.
// Order is preserved from s2.
func SetAndBy[S ~[]E, E any, K comparable](s1 S, s2 S, byFn func(E) K) S {
	if len(s1) == 0 || len(s2) == 0 {
		return nil
	}

	// all elements seen in s1
	seen := NewSetBy(s1, byFn)

	var out S // capacity is unknown
	for _, v := range s2 {
		if !seen.Has(byFn(v)) {
			continue // skip it
		}
		out = append(out, v)
	}

	return out
}
