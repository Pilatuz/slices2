// Package slices2 contains slice related helpers missing in standard "slices" package.
//
// To avoid conflicts with standard "slices" package the package named as slices2:
//
//	import (
//		"github.com/Pilatuz/slices2"
//	)
//
// Some helpers able to do "in-place" work, i.e. there is no memory allocation for result.
// Such functions usually have "InPlace" suffix in their names.
//
// # Filtering and Rejecting
//
// [Filter] and [Reject] are very similar but use inversed condition.
//   - [Filter] removes elements that DO NOT PASS condition
//   - [Reject] removes elements that DO PASS condition
//
// Both have options to modify input slice "in-place" (without memory allocation).
//
// # Uniqueness
//
// [Unique] and [UniqueBy] remove duplicates from a slice.
// Unlike standard slices.Compact, they work with unsorted input.
// Also available "InPlace" versions for memory-efficient operations.
//
// # Set Operations
//
// The package provides functions for set operations:
//   - [SetAnd] — intersection (elements present in both slices)
//   - [SetOr] — union (unique elements from both slices)
//   - [SetSub] — difference (elements in s1 missing in s2)
//   - [SetDiff] — symmetric difference (elements present in only one of the slices)
//
// All set operation functions have "By" variants for custom key extraction.
//
// # Sorting and Reversing
//
// [SortInPlace] and [SortFuncInPlace] sort a slice in ascending order.
// [Sorted] and [SortedFunc] return a new sorted slice.
// [Reversed] returns a new slice with elements in reverse order.
// [ReverseInPlace] reverses elements of a slice in place.
//
// # Transformation
//
// [Transform] transforms each element of a slice from one type to another.
// [TransformEx] transforms elements with error handling.
// Special error [ErrSkip] allows skipping elements.
// [Deref] dereferences pointers, skipping nil values.
//
// # Grouping
//
// [GroupBy] groups slice elements by a key.
//
// # Joining Slices
//
// [Join] joins multiple slices into one.
//
// # Cloning
//
// [Clone] makes a copy of a slice.
//
// # Set Type
//
// Type [Set] represents a set of unique elements based on map[K]struct{}.
// Methods:
//   - Push — add elements
//   - Pop — remove elements
//   - Has — check if element exists
//   - HasAll — check if all elements exist
//   - HasAny — check if any element exists
//   - All — get all elements as a slice
//
// # Go Versions
//
// The package supports Go 1.21+ and earlier versions.
// For Go 1.21+ it uses standard functions slices.Clone, slices.Reverse, slices.Sort.
package slices2
