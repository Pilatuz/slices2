package slices2

// GroupBy splits an input slice s by some key.
func GroupBy[S ~[]E, E any, K comparable](s S, byFn func(E) K) map[K]S {
	out := make(map[K]S)
	for _, v := range s {
		k := byFn(v)
		out[k] = append(out[k], v)
	}

	return out
}
