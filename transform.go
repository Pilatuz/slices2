package slices2

import (
	"errors"
)

// Transform transforms each slice element from type E1 to type E2.
// Returns a new slice of transformed elements.
func Transform[E2 any, S1 ~[]E1, E1 any](s S1, convFn func(E1) E2) []E2 {
	if s == nil {
		return nil // nil transforms to nil
	}

	out := make([]E2, 0, len(s))
	for _, v := range s {
		out = append(out, convFn(v))
	}

	return out
}

// TransformEx transforms each slice element from type E1 to type E2 with error handling.
// Stops on the first error. Special error [ErrSkip] skips elements.
func TransformEx[E2 any, S1 ~[]E1, E1 any](s S1, convFn func(E1) (E2, error)) ([]E2, error) {
	if s == nil {
		return nil, nil // nil transforms to nil without error
	}

	out := make([]E2, 0, len(s))
	for _, v := range s {
		v2, err := convFn(v)
		if err != nil {
			if errors.Is(err, ErrSkip) {
				continue // skip this element
			}
			return nil, err // stop on first error
		}
		out = append(out, v2)
	}

	return out, nil // done
}

// ErrSkip is a special sentinel error that indicates the current element should be skipped.
// Used in transformation functions with error handling.
var ErrSkip = errors.New("skip")

// Deref dereferences pointers to values.
// Nil pointers are skipped with [ErrSkip] error.
func Deref[T any](p *T) (T, error) {
	if p == nil {
		var EMPTY T
		return EMPTY, ErrSkip
	}

	return *p, nil
}
