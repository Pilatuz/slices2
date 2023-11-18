package slices2

// SetOr returns the union between two (or more) sets.
// I.e. unique elements presented in at least one slice.
func SetOr[S ~[]E, E comparable](ss ...S) S {
	return SetOrBy(me[E], ss...)
}

// SetOrBy returns the union between two (or more) sets by custom key.
// I.e. unique elements presented in at least one slice.
func SetOrBy[S ~[]E, E any, K comparable](byFn func(E) K, ss ...S) S {
	switch len(ss) {
	case 0: // no slices
		return nil
	case 1: // only one slice
		return ss[0] // as is // Clone(ss[0])?
	}

	// all elements seen so far
	seen := make(map[K]struct{})

	var out S // capacity is unknown
	for _, s := range ss {
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
