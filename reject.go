package slices2

// Reject removes elements that DO PASS condition.
//
// Returns new slice with elements removed.
func Reject[S ~[]E, E any](condFn func(E) bool, s S) S {
	return reject(initNew[S, E], condFn, s)
}

// RejectInPlace removes elements that DO PASS condition.
//
// Returns original slice with elements removed in-place.
func RejectInPlace[S ~[]E, E any](condFn func(E) bool, s S) S {
	return reject(initSame[S, E], condFn, s)
}

// reject removes elements that DO PASS condition.
func reject[S ~[]E, E any](initFn func(S) S, condFn func(E) bool, s S) S {
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
