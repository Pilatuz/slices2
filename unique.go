package slices2

// Unique gets only unique elements.
//
// Returns new slice with duplicates removed.
func Unique[S ~[]E, E comparable](s S) S {
	if len(s) <= 1 {
		return Clone(s) // a copy of
	}

	return uniqueBy(initNew[S, E], me[E], s)
}

// UniqueBy gets only unique elements by custom key.
//
// Returns new slice with duplicates removed.
func UniqueBy[S ~[]E, K comparable, E any](byFn func(E) K, s S) S {
	if len(s) <= 1 {
		return Clone(s) // a copy of
	}

	return uniqueBy(initNew[S, E], byFn, s)
}

// UniqueInPlace gets only unique elements.
//
// Returns original slice with duplicates removed in-place.
func UniqueInPlace[S ~[]E, E comparable](s S) S {
	if len(s) <= 1 {
		return s // as is
	}

	return uniqueBy(initSame[S, E], me[E], s)
}

// UniqueInPlaceBy gets only unique elements by custom key.
//
// Returns original slice with duplicates removed in-place.
func UniqueInPlaceBy[S ~[]E, K comparable, E any](byFn func(E) K, s S) S {
	if len(s) <= 1 {
		return s // as is
	}

	return uniqueBy(initSame[S, E], byFn, s)
}

// uniqueBy gets only unique elements by custom key.
func uniqueBy[S ~[]E, K comparable, E any](initFn func(S) S, byFn func(E) K, s S) S {
	seen := make(map[K]struct{}, len(s))

	out := initFn(s) // copy or share
	for _, v := range s {
		key := byFn(v)
		if _, ok := seen[key]; ok {
			continue // ignore, already seen
		}

		seen[key] = struct{}{}
		out = append(out, v)
	}

	return out
}
