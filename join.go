package slices2

// Join joins multiple slices.
func Join[S ~[]E, E any](ss ...S) S {
	switch len(ss) {
	case 0: // no slices to join
		return nil
	case 1: // only one slice to join
		return ss[0] // as is
	}

	// output's slice capacity
	var n int
	for _, s := range ss {
		n += len(s)
	}

	if n == 0 { // all slices are empty
		return nil
	}

	// join the slices
	out := make(S, 0, n)
	for _, s := range ss {
		out = append(out, s...)
	}

	return out
}
