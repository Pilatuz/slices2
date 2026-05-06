package slices2

// Reversed returns a new slice with elements in reverse order.
// As opposite to [ReverseInPlace], this function always allocates a new array.
func Reversed[S ~[]E, E any](s S) S {
	s = Clone(s)
	ReverseInPlace(s)
	return s
}
