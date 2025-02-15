package slices2_test

import (
	"fmt"
	"testing"
	"unicode/utf8"

	"github.com/Pilatuz/slices2"
)

// ExampleReject an example for `Reject` function.
func ExampleReject() {
	ss := []string{"foo", "barbaz"}
	cond := func(s string) bool {
		return utf8.RuneCountInString(s) > 3
	}
	fmt.Println(slices2.Reject(ss, cond))
	// Output:
	// [foo]
}

// TestReject unit tests for `Reject` function
func TestReject(tt *testing.T) {
	// string
	tt.Run("str", func(t *testing.T) {
		var Nil []string
		Empty := []string{}
		Foo := []string{"foo"}
		FooBar := []string{"foo", "bar"}
		BarFoo := []string{"bar", "foo"}
		FooBarBaz := []string{"foo", "bar", "baz"}

		test := func(input, expected []string) {
			t.Helper()

			isFoo := func(s string) bool {
				return s == "foo"
			}

			if actual := slices2.Reject(input, isFoo); !equal(actual, expected) || same(actual, input) {
				t.Errorf("Reject(`%#v`)=`%#v`, expected `%#v`", input, actual, expected)
			}

			input2 := slices2.Clone(input)
			if actual := slices2.RejectInPlace(input2, isFoo); !equal(actual, expected) {
				t.Errorf("RejectInPlace(`%#v`)=`%#v`, expected `%#v`", input, actual, expected)
			} else if len(actual) > 0 && !same(actual, input2) {
				t.Errorf("RejectInPlace(`%#v`)=`%#v`, expected `%#v`", input, actual, expected)
			}
		}

		test(Nil, Nil)                 // Reject(nil) => nil
		test(Empty, Empty)             // Reject([]) => []
		test(Foo, Empty)               // Reject([foo]) => []
		test(FooBar, FooBar[1:])       // Reject([foo bar]) => [bar]
		test(BarFoo, FooBar[1:])       // Reject([bar foo]) => [bar]
		test(FooBarBaz, FooBarBaz[1:]) // Reject([foo bar baz]) => [bar baz]
	})

	// integer
	tt.Run("int", func(t *testing.T) {
		var Nil []int
		Empty := []int{}
		Foo := []int{123}
		FooBar := []int{123, 456}
		BarFoo := []int{456, 123}
		FooBarBaz := []int{123, 456, 789}

		test := func(input, expected []int) {
			t.Helper()

			isFoo := func(s int) bool {
				return s == 123
			}

			if actual := slices2.Reject(input, isFoo); !equal(actual, expected) || same(actual, input) {
				t.Errorf("Reject(`%#v`)=`%#v`, expected `%#v`", input, actual, expected)
			}

			input2 := slices2.Clone(input)
			if actual := slices2.RejectInPlace(input2, isFoo); !equal(actual, expected) {
				t.Errorf("RejectInPlace(`%#v`)=`%#v`, expected `%#v`", input, actual, expected)
			} else if len(actual) > 0 && !same(actual, input2) {
				t.Errorf("RejectInPlace(`%#v`)=`%#v`, expected `%#v`", input, actual, expected)
			}
		}

		test(Nil, Nil)                 // Reject(nil) => nil
		test(Empty, Empty)             // Reject([]) => []
		test(Foo, Empty)               // Reject([foo]) => []
		test(FooBar, FooBar[1:])       // Reject([foo bar]) => [bar]
		test(BarFoo, FooBar[1:])       // Reject([bar foo]) => [bar]
		test(FooBarBaz, FooBarBaz[1:]) // Reject([foo bar baz]) => [bar baz]
	})
}
