package slices2

// Set represents a collection of unique elements.
// It's a thin wrapper over map[K]struct{} for convenient set operations.
//
// The Set type supports all basic set operations:
// add (Push), remove (Pop), check presence (Has, HasAll, HasAny),
// and get all elements (All).
type Set[K comparable] map[K]struct{}

// NewSet creates a new set with optional elements.
// Duplicates are automatically removed.
func NewSet[K comparable](kk ...K) Set[K] {
	out := make(Set[K], len(kk)) // with capacity
	out.Push(kk...)
	return out
}

// NewSetBy creates a new set from a slice with element transformation via key.
// The byFn function extracts a key from each element.
func NewSetBy[S ~[]E, E any, K comparable](s S, byFn func(E) K) Set[K] {
	out := make(Set[K], len(s)) // with capacity
	for _, v := range s {
		out[byFn(v)] = struct{}{} // see Push()
	}
	return out
}

// Push adds elements to the set.
// Duplicates are ignored.
func (s Set[K]) Push(kk ...K) {
	for _, k := range kk {
		s[k] = struct{}{}
	}
}

// Pop removes elements from the set.
// Removing non-existent elements is ignored.
func (s Set[K]) Pop(kk ...K) {
	for _, k := range kk {
		delete(s, k)
	}
}

// Has checks if an element is present in the set.
// Returns true if the element exists.
func (s Set[K]) Has(k K) bool {
	_, ok := s[k]
	return ok
}

// HasAll checks if ALL specified elements are present in the set.
// Returns true if all elements exist.
// Empty argument list returns true.
func (s Set[K]) HasAll(kk ...K) bool {
	for _, k := range kk {
		if !s.Has(k) {
			return false
		}
	}

	return true
}

// HasAny checks if ANY of the specified elements is present in the set.
// Returns true if at least one element exists in the set.
// Empty argument list returns false.
func (s Set[K]) HasAny(kk ...K) bool {
	for _, k := range kk {
		if s.Has(k) {
			return true
		}
	}

	return false
}

// All returns all elements of the set as a slice.
// Element order is undefined (depends on map implementation).
// Returns nil for empty sets.
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
