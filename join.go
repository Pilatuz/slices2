package slices2

// Join joins multiple slices into one.
// Returns a new slice containing elements from all input slices.
func Join[S ~[]E, E any](ss ...S) S {
	// take a look at slices.Concat

	switch len(ss) {
	case 0: // no slices to join
		return nil
	case 1: // only one slice
		return ss[0] // as is
	}

	// capacity of the resulting slice
	var n int
	for _, s := range ss {
		n += len(s)
	}

	if n == 0 { // all slices are empty
		return nil
	}

	// join slices
	out := make(S, 0, n)
	for _, s := range ss {
		out = append(out, s...)
	}

	return out
}
