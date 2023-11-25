//go:build go1.21

package slices2_test

import (
	"cmp"
	"fmt"
	"strings"
	"testing"

	"github.com/Pilatuz/slices2"
)

// ExampleSorted an example for `Sorted` function.
func ExampleSorted() {
	s := []int{3, 1, 2}
	r := slices2.Sorted(s)
	fmt.Println(s)
	fmt.Println(r)
	// Output:
	// [3 1 2]
	// [1 2 3]
}

// TestSort unit tests for `Sort` function.
func TestSort(tt *testing.T) {
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
			if slices2.SortInPlace(y); !equal(y, e) {
				t.Errorf("SortInPlace(`%#v`)=`%#v`, expected `%#v`", x, y, e)
			}
			if a := slices2.Sorted(x); same(a, x) || !equal(a, e) {
				t.Errorf("Sorted(`%#v`)=`%#v`, expected `%#v`", x, a, e)
			}
			if a := slices2.SortedFunc(x, strings.Compare); same(a, x) || !equal(a, e) {
				t.Errorf("SortedFunc(`%#v`)=`%#v`, expected `%#v`", x, a, e)
			}
		}

		test(Nil, Nil)     // sort(nil) => nil
		test(Empty, Empty) // sort([]) => []

		test(Foo, Foo) // sort([foo]) => [foo]
		test(FooBar, []string{"bar", "foo"})
		test(slices2.Reversed(FooBar), []string{"bar", "foo"})
		test(FooBarBaz, []string{"bar", "baz", "foo"})
		test(slices2.Reversed(FooBarBaz), []string{"bar", "baz", "foo"})
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
			if slices2.SortInPlace(y); !equal(y, e) {
				t.Errorf("SortInPlace(`%#v`)=`%#v`, expected `%#v`", x, y, e)
			}
			if a := slices2.Sorted(x); same(a, x) || !equal(a, e) {
				t.Errorf("Sorted(`%#v`)=`%#v`, expected `%#v`", x, a, e)
			}
			if a := slices2.SortedFunc(x, cmp.Compare[int]); same(a, x) || !equal(a, e) {
				t.Errorf("SortedFunc(`%#v`)=`%#v`, expected `%#v`", x, a, e)
			}
		}

		test(Nil, Nil)     // sort(nil) => nil
		test(Empty, Empty) // sort([]) => []

		test(Foo, Foo) // sort([foo]) => [foo]
		test(FooBar, []int{123, 456})
		test(slices2.Reversed(FooBar), []int{123, 456})
		test(FooBarBaz, []int{123, 456, 789})
		test(slices2.Reversed(FooBarBaz), []int{123, 456, 789})
	})
}
