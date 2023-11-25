package slices2

// Reversed returns new slice with elements of input slice s reversed.
// As opposite to [ReverseInPlace] this function always allocates new array.
func Reversed[S ~[]E, E any](s S) S {
	s = Clone(s)
	ReverseInPlace(s)
	return s
}
