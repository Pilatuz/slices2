package slices2_test

import (
	"fmt"
	"testing"

	"github.com/Pilatuz/slices2"
)

// ExampleReversed an example for `Reversed` function.
func ExampleReversed() {
	s := []int{1, 2, 3}
	r := slices2.Reversed(s)
	fmt.Println(s)
	fmt.Println(r)
	// Output:
	// [1 2 3]
	// [3 2 1]
}

// TestReverse unit tests for `Reverse` function.
func TestReverse(tt *testing.T) {
	// string
	tt.Run("str", func(t *testing.T) {
		var Nil []string
		Empty := []string{}
		Foo := []string{"foo"}
		FooBar := []string{"foo", "bar"}
		FooBarBaz := []string{"foo", "bar", "baz"}

		test := func(x []string, e []string) {
			t.Helper()
			y := slices2.Clone(x)
			if slices2.ReverseInPlace(y); !equal(y, e) {
				t.Errorf("ReverseInPlace(`%#v`)=`%#v`, expected `%#v`", x, y, e)
			}
			if a := slices2.Reversed(x); same(a, x) || !equal(a, e) {
				t.Errorf("Reversed(`%#v`)=`%#v`, expected `%#v`", x, a, e)
			}
		}

		test(Nil, Nil)     // reverse(nil) => nil
		test(Empty, Empty) // reverse([]) => []

		test(Foo, Foo) // reverse([foo]) => [foo]
		test(FooBar, []string{"bar", "foo"})
		test(FooBarBaz, []string{"baz", "bar", "foo"})
	})

	// integer
	tt.Run("int", func(t *testing.T) {
		var Nil []int
		Empty := []int{}
		Foo := []int{123}
		FooBar := []int{123, 456}
		FooBarBaz := []int{123, 456, 789}

		test := func(x []int, e []int) {
			t.Helper()
			y := slices2.Clone(x)
			if slices2.ReverseInPlace(y); !equal(y, e) {
				t.Errorf("ReverseInPlace(`%#v`)=`%#v`, expected `%#v`", x, y, e)
			}
			if a := slices2.Reversed(x); same(a, x) || !equal(a, e) {
				t.Errorf("Reversed(`%#v`)=`%#v`, expected `%#v`", x, a, e)
			}
		}

		test(Nil, Nil)     // reverse(nil) => nil
		test(Empty, Empty) // reverse([]) => []

		test(Foo, Foo) // reverse([foo]) => [foo]
		test(FooBar, []int{456, 123})
		test(FooBarBaz, []int{789, 456, 123})
	})
}
