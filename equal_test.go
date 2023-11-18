package slices2_test

import (
	"testing"

	"github.com/Pilatuz/slices2"
)

// equal checks if two slices are equal.
func equal[S ~[]E, E comparable](a, b S) bool {
	if a == nil {
		return b == nil
	} else if b == nil {
		return false
	}

	if len(a) != len(b) {
		return false
	}

	for i := range a {
		if a[i] != b[i] {
			return false
		}
	}

	return true
}

// same checks if two slices are the same.
// both nil or empty slices are NOT the same!
func same[S ~[]E, E any](a, b S) bool {
	if a == nil || b == nil {
		return false // two nils are NOT same!
	}

	// check address of the latest element:
	A, B := cap(a), cap(b)
	return A > 0 && B > 0 &&
		&a[:A][A-1] == &b[:B][B-1]
}

// TestEqual unit tests for `equal` helper.
func TestEqual(tt *testing.T) {
	// string
	tt.Run("str", func(t *testing.T) {
		var Nil []string
		Empty := []string{}
		Foo := []string{"foo"}
		Bar := []string{"bar"}
		FooBar := []string{"foo", "bar"}
		BarFoo := []string{"bar", "foo"}

		test := func(a, b []string, expected bool) {
			t.Helper()

			if actual := equal(a, b); actual != expected {
				t.Errorf("equal(`%#v`, `%#v`)=%t, expected %t", a, b, actual, expected)
			}
		}

		test(Nil, Nil, true)     // nil == nil
		test(Nil, Empty, false)  // nil != []
		test(Empty, Nil, false)  // [] != nil
		test(Empty, Empty, true) // [] == []

		test(Empty, Foo, false)             // [] != [foo]
		test(Foo, Empty, false)             // [foo] != []
		test(Foo, Foo, true)                // [foo] == [foo]
		test(Foo, slices2.Clone(Foo), true) // [foo] == [foo]
		test(slices2.Clone(Foo), Foo, true) // [foo] == [foo]

		test(Foo, Bar, false)       // [foo] != [bar]
		test(Foo, FooBar, false)    // [foo] != [foo bar]
		test(FooBar, BarFoo, false) // [foo bar] != [bar foo]
	})

	// integer
	tt.Run("int", func(t *testing.T) {
		var Nil []int
		Empty := []int{}
		Foo := []int{123}
		Bar := []int{456}
		FooBar := []int{123, 456}
		BarFoo := []int{456, 123}

		test := func(a, b []int, expected bool) {
			t.Helper()

			if actual := equal(a, b); actual != expected {
				t.Errorf("equal(`%#v`, `%#v`)=%t, expected %t", a, b, actual, expected)
			}
		}

		test(Nil, Nil, true)     // nil == nil
		test(Nil, Empty, false)  // nil != []
		test(Empty, Nil, false)  // [] != nil
		test(Empty, Empty, true) // [] == []

		test(Empty, Foo, false)             // [] != [foo]
		test(Foo, Empty, false)             // [foo] != []
		test(Foo, Foo, true)                // [foo] == [foo]
		test(Foo, slices2.Clone(Foo), true) // [foo] == [foo]
		test(slices2.Clone(Foo), Foo, true) // [foo] == [foo]

		test(Foo, Bar, false)       // [foo] != [bar]
		test(Foo, FooBar, false)    // [foo] != [foo bar]
		test(FooBar, BarFoo, false) // [foo bar] != [bar foo]
	})
}

// TestSame unit tests for `same` helper.
func TestSame(tt *testing.T) {
	// string
	tt.Run("str", func(t *testing.T) {
		var Nil []string
		Empty := []string{}
		Foo := []string{"foo"}
		Bar := []string{"bar"}
		FooBar := []string{"foo", "bar"}
		BarFoo := []string{"bar", "foo"}

		test := func(a, b []string, expected bool) {
			t.Helper()

			if actual := same(a, b); actual != expected {
				t.Errorf("same(`%#v`, `%#v`)=%t, expected %t", a, b, actual, expected)
			}
		}

		test(Nil, Nil, false)     // nil != nil
		test(Nil, Empty, false)   // nil != []
		test(Empty, Nil, false)   // [] != nil
		test(Empty, Empty, false) // [] != []

		test(Empty, Foo, false)              // [] != [foo]
		test(Foo, Empty, false)              // [foo] != []
		test(Foo, Foo, true)                 // [foo] == [foo]
		test(Foo, slices2.Clone(Foo), false) // [foo] != [foo]*
		test(slices2.Clone(Foo), Foo, false) // [foo]* != [foo]

		test(Foo, Bar, false)       // [foo] != [bar]
		test(Foo, FooBar, false)    // [foo] != [foo bar]
		test(FooBar, BarFoo, false) // [foo bar] != [bar foo]
	})

	// integer
	tt.Run("int", func(t *testing.T) {
		var Nil []int
		Empty := []int{}
		Foo := []int{123}
		Bar := []int{456}
		FooBar := []int{123, 456}
		BarFoo := []int{456, 123}

		test := func(a, b []int, expected bool) {
			t.Helper()

			if actual := same(a, b); actual != expected {
				t.Errorf("same(`%#v`, `%#v`)=%t, expected %t", a, b, actual, expected)
			}
		}

		test(Nil, Nil, false)     // nil != nil
		test(Nil, Empty, false)   // nil != []
		test(Empty, Nil, false)   // [] != nil
		test(Empty, Empty, false) // [] != []

		test(Empty, Foo, false)              // [] != [foo]
		test(Foo, Empty, false)              // [foo] != []
		test(Foo, Foo, true)                 // [foo] == [foo]
		test(Foo, slices2.Clone(Foo), false) // [foo] == [foo]*
		test(slices2.Clone(Foo), Foo, false) // [foo]* == [foo]

		test(Foo, Bar, false)       // [foo] != [bar]
		test(Foo, FooBar, false)    // [foo] != [foo bar]
		test(FooBar, BarFoo, false) // [foo bar] != [bar foo]
	})
}
