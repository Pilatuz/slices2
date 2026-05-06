package slices2

// Filter removes elements that DO NOT PASS condition.
// Returns new slice with elements that pass the condition.
//
// The condFn function returns true for elements that should be kept.
func Filter[S ~[]E, E any](s S, condFn func(E) bool) S {
	return filter(s, initNew[S, E], condFn)
}

// FilterInPlace removes elements that DO NOT PASS condition.
// Returns original slice with elements removed in-place (no memory allocation).
//
// The condFn function returns true for elements that should be kept.
func FilterInPlace[S ~[]E, E any](s S, condFn func(E) bool) S {
	return filter(s, initSame[S, E], condFn)
}

// filter removes elements that DO NOT PASS condition.
func filter[S ~[]E, E any](s S, initFn func(S) S, condFn func(E) bool) S {
	if s == nil {
		return nil
	}

	out := initFn(s) // copy or share
	for _, v := range s {
		if !condFn(v) {
			continue // skip it
		}
		out = append(out, v)
	}

	return out
}
