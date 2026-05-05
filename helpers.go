package slices2

// me returns element as a key (trivial transformation).
// Used by default in *By functions for comparable types.
// Likely to be inlined by compiler.
func me[K comparable](v K) K {
	return v
}

// initNew initializes a new slice for copying.
// Creates a slice with zero length and capacity of the original.
// Likely to be inlined by compiler.
func initNew[S ~[]E, E any](from S) S {
	return make(S, 0, len(from)) // new slice
}

// initSame initializes the same slice (reset).
// Returns a slice with zero length sharing memory with the original.
// Used for InPlace operations without memory allocation.
// Likely to be inlined by compiler.
func initSame[S ~[]E, E any](from S) S {
	return from[0:0] // share memory
}
