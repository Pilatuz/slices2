// package slices2 contains slice related helpers missing in standard "slices" package.
//
// To avoid conflicts with standard "slices" package the package name is as slices2:
//
//	import (
//		"github.com/Pilatuz/slices2"
//	)
//
// Some helpers able to do "in-place" work, i.e. there is no memory allocation for result.
// Such functions usually have "InPlace" suffix in their names.
//
// [Filter] and [Reject] are very similar but use inversed condition.
// [Filter] removes elements that DO NOT PASS condition
// while [Reject] removes elements that DO PASS condition.
// Both have option to modify input slice "in-place".
//
// A few helpers allow to do operations on sets:
// [SetAnd] and [SetOr] calculates conjunction and disjunction.
// [SetDiff] finds difference: i.e. elements added/removed in another set.
//
// [Unique] can work with unsorted input (opposite to standard `slices.Compact`)
// and returns result in new slice.
package slices2
