package slices2_test

import (
	"fmt"
	"testing"

	"github.com/Pilatuz/slices2"
)

// ExampleSetOr an example for `SetOr` function.
func ExampleSetOr() {
	a := []string{"foo", "bar"}
	b := []string{"bar", "baz"}

	c := slices2.SetOr(a, b)
	fmt.Println(a, "|", b, "=", c)
	// Output:
	// [foo bar] | [bar baz] = [foo bar baz]
}

// TestSetOr unit tests for `SetOr` function.
func TestSetOr(tt *testing.T) {
	// string
	tt.Run("str", func(t *testing.T) {
		var Nil []string
		Empty := []string{}
		Foo := []string{"foo"}
		Bar := []string{"bar"}
		FooBar := []string{"foo", "bar"}
		BarFoo := []string{"bar", "foo"}
		FooBarBaz := []string{"foo", "bar", "baz"}

		test := func(expected []string, input ...[]string) {
			t.Helper()
			if actual := slices2.SetOr(input...); !equal(actual, expected) {
				t.Errorf("SetOr(`%#v`)=`%#v`, expected `%#v`", input, actual, expected)
			}
			if actual := slices2.SetOrBy(func(v string) string { return v }, input...); !equal(actual, expected) {
				t.Errorf("SetOrBy(`%#v`)=`%#v`, expected `%#v`", input, actual, expected)
			}
		}

		test(Nil, Nil, Nil)     // nil || nil => nil
		test(Nil, Nil, Empty)   // nil || [] => nil
		test(Nil, Empty, Nil)   // [] || nil => nil
		test(Nil, Empty, Empty) // [] || [] => nil
		test(Foo, Nil, Foo)     // nil || [foo] => [foo]
		test(Foo, Empty, Foo)   // [] && [foo] => [foo]
		test(Foo, Foo, Nil)     // [foo] && nil => [foo]
		test(Foo, Foo, Empty)   // [foo] && [] => [foo]

		test(FooBar, Foo, Bar) // [foo] || [bar] => [foo bar]
		test(BarFoo, Bar, Foo) // [foo] || [bar] => [bar foo]

		test(FooBar, FooBar, Foo)       // [foo bar] || [foo] => [foo bar]
		test(FooBar, Foo, FooBar)       // [foo] || [foo bar] => [foo bar]
		test(FooBarBaz, FooBarBaz, Foo) // [foo bar baz] || [foo] => [foo bar baz]
		test(FooBarBaz, Foo, FooBarBaz) // [foo] || [foo bar baz] => [foo bar baz]

		test(FooBarBaz, FooBar, FooBarBaz)    // [foo bar] || [foo bar baz] => [foo bar baz]
		test(FooBarBaz, FooBarBaz, FooBar)    // [foo bar baz] || [foo bar] => [foo bar baz]
		test(FooBarBaz, FooBarBaz, FooBarBaz) // [foo bar baz] || [foo bar baz] => [foo bar baz]
		test(FooBar, FooBar, FooBar)          // [foo bar] || [foo bar] => [foo bar]
		test(Foo, Foo, Foo)                   // [foo] || [foo] => [foo]

		test(Nil)
		test(Empty, Empty)
		test(Foo, Foo, Foo, Empty, Nil)
		test(FooBar, Nil, Empty, Foo, Empty, Bar, Empty, Nil, FooBar)
	})

	// integer
	tt.Run("int", func(t *testing.T) {
		var Nil []int
		Empty := []int{}
		Foo := []int{123}
		Bar := []int{456}
		FooBar := []int{123, 456}
		BarFoo := []int{456, 123}
		FooBarBaz := []int{123, 456, 789}

		test := func(expected []int, input ...[]int) {
			t.Helper()
			if actual := slices2.SetOr(input...); !equal(actual, expected) {
				t.Errorf("SetOr(`%#v`)=`%#v`, expected `%#v`", input, actual, expected)
			}
			if actual := slices2.SetOrBy(func(v int) int { return v }, input...); !equal(actual, expected) {
				t.Errorf("SetOrBy(`%#v`)=`%#v`, expected `%#v`", input, actual, expected)
			}
		}

		test(Nil, Nil, Nil)     // nil || nil => nil
		test(Nil, Nil, Empty)   // nil || [] => nil
		test(Nil, Empty, Nil)   // [] || nil => nil
		test(Nil, Empty, Empty) // [] || [] => nil
		test(Foo, Nil, Foo)     // nil || [foo] => [foo]
		test(Foo, Empty, Foo)   // [] && [foo] => [foo]
		test(Foo, Foo, Nil)     // [foo] && nil => [foo]
		test(Foo, Foo, Empty)   // [foo] && [] => [foo]

		test(FooBar, Foo, Bar) // [foo] || [bar] => [foo bar]
		test(BarFoo, Bar, Foo) // [foo] || [bar] => [bar foo]

		test(FooBar, FooBar, Foo)       // [foo bar] || [foo] => [foo bar]
		test(FooBar, Foo, FooBar)       // [foo] || [foo bar] => [foo bar]
		test(FooBarBaz, FooBarBaz, Foo) // [foo bar baz] || [foo] => [foo bar baz]
		test(FooBarBaz, Foo, FooBarBaz) // [foo] || [foo bar baz] => [foo bar baz]

		test(FooBarBaz, FooBar, FooBarBaz)    // [foo bar] || [foo bar baz] => [foo bar baz]
		test(FooBarBaz, FooBarBaz, FooBar)    // [foo bar baz] || [foo bar] => [foo bar baz]
		test(FooBarBaz, FooBarBaz, FooBarBaz) // [foo bar baz] || [foo bar baz] => [foo bar baz]
		test(FooBar, FooBar, FooBar)          // [foo bar] || [foo bar] => [foo bar]
		test(Foo, Foo, Foo)                   // [foo] || [foo] => [foo]

		test(Nil)
		test(Empty, Empty)
		test(Foo, Foo, Foo, Empty, Nil)
		test(FooBar, Nil, Empty, Foo, Empty, Bar, Empty, Nil, FooBar)
	})
}
