package slices2

// Unique removes duplicates from a slice.
// Returns new slice with unique elements.
// Order is preserved (first occurrence remains).
func Unique[S ~[]E, E comparable](s S) S {
	if len(s) <= 1 {
		return Clone(s)
	}

	return uniqueBy(s, initNew[S, E], me[E])
}

// UniqueBy removes duplicates from a slice by custom key.
// Returns new slice with unique elements.
// The byFn function extracts a comparison key from each element.
// Order is preserved (first occurrence remains).
func UniqueBy[S ~[]E, K comparable, E any](s S, byFn func(E) K) S {
	if len(s) <= 1 {
		return Clone(s) // a copy of
	}

	return uniqueBy(s, initNew[S, E], byFn)
}

// UniqueInPlace removes duplicates from a slice.
// Returns original slice with duplicates removed in-place (no memory allocation).
// Order is preserved (first occurrence remains).
func UniqueInPlace[S ~[]E, E comparable](s S) S {
	if len(s) <= 1 {
		return s // as is
	}

	return uniqueBy(s, initSame[S, E], me[E])
}

// UniqueInPlaceBy removes duplicates from a slice by custom key.
// Returns original slice with duplicates removed in-place (no memory allocation).
// The byFn function extracts a comparison key from each element.
// Order is preserved (first occurrence remains).
func UniqueInPlaceBy[S ~[]E, K comparable, E any](s S, byFn func(E) K) S {
	if len(s) <= 1 {
		return s // as is
	}

	return uniqueBy(s, initSame[S, E], byFn)
}

// uniqueBy removes duplicates from a slice by custom key.
func uniqueBy[S ~[]E, K comparable, E any](s S, initFn func(S) S, byFn func(E) K) S {
	seen := make(Set[K], len(s))

	out := initFn(s) // copy or share
	for _, v := range s {
		key := byFn(v)
		if seen.Has(key) {
			continue // skip duplicate
		}

		out = append(out, v)
		seen.Push(key)
	}

	return out
}
