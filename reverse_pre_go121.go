//go:build !go1.21

package slices2

// ReverseInPlace reverses elements of a slice (in place).
func ReverseInPlace[S ~[]E, E any](s S) {
	for i, j := 0, len(s)-1; i < j; i, j = i+1, j-1 {
		s[i], s[j] = s[j], s[i]
	}
}
