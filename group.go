package slices2

// GroupBy groups slice elements by a key.
// Returns a map where keys are results of byFn function
// and values are slices of elements with that key.
func GroupBy[S ~[]E, E any, K comparable](s S, byFn func(E) K) map[K]S {
	out := make(map[K]S)
	for _, v := range s {
		k := byFn(v)
		out[k] = append(out[k], v)
	}

	return out
}
