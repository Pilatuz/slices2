package slices2

// Reject removes elements that DO PASS condition.
//
// Returns new slice with elements removed.
func Reject[S ~[]E, E any](s S, condFn func(E) bool) S {
	return reject(s, initNew[S, E], condFn)
}

// RejectInPlace removes elements that DO PASS condition.
//
// Returns original slice with elements removed in-place.
func RejectInPlace[S ~[]E, E any](s S, condFn func(E) bool) S {
	return reject(s, initSame[S, E], condFn)
}

// reject removes elements that DO PASS condition.
func reject[S ~[]E, E any](s S, initFn func(S) S, condFn func(E) bool) S {
	if s == nil {
		return nil
	}

	out := initFn(s) // copy or share
	for _, v := range s {
		if condFn(v) {
			continue // skip it
		}
		out = append(out, v)
	}

	return out
}
