package slices2

import (
	"errors"
)

// Transform transforms each element of slice E1 to E2.
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

// TransformEx transforms each element of slice E1 to E2 and check errors.
// Stops on first error. Special ErrSkip skips elements.
func TransformEx[E2 any, S1 ~[]E1, E1 any](s S1, convFn func(E1) (E2, error)) ([]E2, error) {
	if s == nil {
		return nil, nil // nil transforms to nil with no error
	}

	out := make([]E2, 0, len(s))
	for _, v := range s {
		v2, err := convFn(v)
		if err != nil {
			if errors.Is(err, ErrSkip) {
				continue // skip it
			}
			return nil, err // stop on first error
		}
		out = append(out, v2)
	}

	return out, nil // done
}

// ErrSkip special sentinel error indicating that current element should be skipped.
var ErrSkip = errors.New("skip")
