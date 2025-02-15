package slices2

// SetDiff returns the difference between two sets.
// Where out1 - elements presented in s1 but missing in s2.
// And out2 - elements presented in s2 but missing in s1.
func SetDiff[S ~[]E, E comparable](s1 S, s2 S) (out1 S, out2 S) {
	return SetDiffBy(s1, s2, me[E])
}

// SetDiffBy returns the difference between two slices by custom key.
// Where out1 - elements presented in s1 but missing in s2.
// And out2 - elements presented in s2 but missing in s1.
func SetDiffBy[S ~[]E, E any, K comparable](s1 S, s2 S, byFn func(E) K) (out1 S, out2 S) {
	if len(s1) == 0 || len(s2) == 0 {
		return s1, s2 // Clone(s1), Clone(s2)?
	}

	// elements seen in s1
	seen1 := make(map[K]struct{}, len(s1))
	for _, v := range s1 {
		seen1[byFn(v)] = struct{}{}
	}

	// elements seen in s2
	seen2 := make(map[K]struct{}, len(s2))
	for _, v := range s2 {
		seen2[byFn(v)] = struct{}{}

		// presented in s2, missing in s1
		if _, ok := seen1[byFn(v)]; ok {
			continue
		}
		out2 = append(out2, v)
	}

	// presented in s1, missing in s2
	for _, v := range s1 {
		if _, ok := seen2[byFn(v)]; ok {
			continue // skip it
		}
		out1 = append(out1, v)
	}

	return
}
