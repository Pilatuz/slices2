package slices2

// SetOr returns the union between two sets.
// I.e. unique elements presented in at least one slice.
func SetOr[S ~[]E, E comparable](s1 S, s2 S) S {
	return SetOrBy(s1, s2, me[E])
}

// SetOrBy returns the union between two sets by custom key.
// I.e. unique elements presented in at least one slice.
func SetOrBy[S ~[]E, E any, K comparable](s1 S, s2 S, byFn func(E) K) S {
	// all elements seen so far
	seen := make(map[K]struct{})

	var out S // capacity is unknown
	for _, s := range []S{s1, s2} {
		for _, v := range s {
			key := byFn(v)
			if _, ok := seen[key]; ok {
				continue // skip it
			}
			seen[key] = struct{}{}
			out = append(out, v)
		}
	}

	return out
}
