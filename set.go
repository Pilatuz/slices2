package slices2

// Set is a collection of unique elements.
// Just a thin wrapper on map of struct{}.
type Set[K comparable] map[K]struct{}

// NewSet creates a new set with optional elements.
func NewSet[K comparable](kk ...K) Set[K] {
	out := make(Set[K], len(kk)) // with capacity
	out.Push(kk...)
	return out
}

// NewSetBy creates a new set with optional elements converted.
func NewSetBy[S ~[]E, E any, K comparable](s S, byFn func(E) K) Set[K] {
	out := make(Set[K], len(s)) // with capacity
	for _, v := range s {
		out[byFn(v)] = struct{}{} // see Push()
	}
	return out
}

// Push adds some elements to the set.
func (s Set[K]) Push(kk ...K) {
	for _, k := range kk {
		s[k] = struct{}{}
	}
}

// Pop removes some elements from the set.
func (s Set[K]) Pop(kk ...K) {
	for _, k := range kk {
		delete(s, k)
	}
}

// Has checks if an element is present in the set.
func (s Set[K]) Has(k K) bool {
	_, ok := s[k]
	return ok
}

// HasAll checks if ALL elements are present in the set.
func (s Set[K]) HasAll(kk ...K) bool {
	for _, k := range kk {
		if !s.Has(k) {
			return false
		}
	}

	return true
}

// HasAny checks if ANY element is present in the set.
func (s Set[K]) HasAny(kk ...K) bool {
	for _, k := range kk {
		if s.Has(k) {
			return true
		}
	}

	return false
}

// All gets all elements as a slice.
// Order is undefined!
func (s Set[K]) All() []K {
	if len(s) == 0 {
		return nil
	}

	out := make([]K, 0, len(s))
	for k := range s {
		out = append(out, k)
	}

	return out
}
