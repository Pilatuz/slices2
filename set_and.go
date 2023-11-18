package slices2

// SetAnd returns the intersection between two sets.
// I.e. elements presented in both slices.
func SetAnd[S ~[]E, E comparable](s1 S, s2 S) S {
	return SetAndBy(me[E], s1, s2)
}

// SetAndBy returns the intersection between two sets by custom key.
// I.e. elements presented in both slices.
func SetAndBy[S ~[]E, E any, K comparable](byFn func(E) K, s1 S, s2 S) S {
	if len(s1) == 0 || len(s2) == 0 {
		return nil
	}

	// all elements seen in s1
	seen := make(map[K]struct{}, len(s1))
	for _, v := range s1 {
		seen[byFn(v)] = struct{}{}
	}

	var out S // capacity is unknown
	for _, v := range s2 {
		if _, ok := seen[byFn(v)]; !ok {
			continue // skip it
		}
		out = append(out, v)
	}

	return out
}
