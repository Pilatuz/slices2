package slices2

// me uses element itself as a trivial key.
// Likely this function should be inlined by compiler.
func me[K comparable](v K) K {
	return v
}

// initNew initializes a copy of slice.
// Likely this function should be inlined by compiler.
func initNew[S ~[]E, E any](from S) S {
	return make(S, 0, len(from)) // new
}

// initSame uses the same slice (reset).
// Likely this function should be inlined by compiler.
func initSame[S ~[]E, E any](from S) S {
	return from[0:0] // share memory
}
