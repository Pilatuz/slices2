package slices2

// SetSub returns set where all s2 elements removed from s1.
// I.e. all elements presented in s1 and missing in s2.
func SetSub[S ~[]E, E comparable](s1 S, s2 S) S {
	return SetSubBy(s1, s2, me[E])
}

// SetSubBy returns set where all s2 elements removed from s1 by custom key.
// I.e. all elements presented in s1 and missing in s2.
//
// The byFn function extracts a comparison key from each element.
func SetSubBy[S ~[]E, E any, K comparable](s1 S, s2 S, byFn func(E) K) S {
	if len(s2) == 0 {
		return s1 // Clone(s1)?
	}

	// all elements seen in s2
	seen := NewSetBy(s2, byFn)

	var out S // capacity is unknown
	for _, v := range s1 {
		if seen.Has(byFn(v)) {
			continue // skip it
		}
		out = append(out, v)
	}

	return out
}
