package slices2_test

import (
	"fmt"
	"testing"

	"github.com/Pilatuz/slices2"
)

// ExampleSetAnd an example for `SetAnd` function.
func ExampleSetAnd() {
	a := []string{"foo", "bar"}
	b := []string{"bar", "baz"}

	c := slices2.SetAnd(a, b)
	fmt.Println(a, "&", b, "=", c)
	// Output:
	// [foo bar] & [bar baz] = [bar]
}

// TestSetAnd unit tests for `SetAnd` function.
func TestSetAnd(tt *testing.T) {
	// string
	tt.Run("str", func(t *testing.T) {
		var Nil []string
		Empty := []string{}
		Foo := []string{"foo"}
		Bar := []string{"bar"}
		FooBar := []string{"foo", "bar"}
		FooBarBaz := []string{"foo", "bar", "baz"}

		test := func(a, b []string, e1 []string) {
			t.Helper()
			if a1 := slices2.SetAnd(a, b); !equal(a1, e1) {
				t.Errorf("SetAnd(`%#v`,`%#v`)=`%#v`, expected `%#v`", a, b, a1, e1)
			}
			if a1 := slices2.SetAndBy(func(v string) string { return v }, a, b); !equal(a1, e1) {
				t.Errorf("SetAndBy(`%#v`,`%#v`)=`%#v`, expected `%#v`", a, b, a1, e1)
			}
		}

		test(Nil, Nil, Nil)     // nil && nil => nil
		test(Nil, Empty, Nil)   // nil && [] => nil
		test(Empty, Nil, Nil)   // [] && nil => nil
		test(Empty, Empty, Nil) // [] && [] => nil
		test(Nil, Foo, Nil)     // nil && [foo] => nil
		test(Empty, Foo, Nil)   // [] && [foo] => nil
		test(Foo, Nil, Nil)     // [foo] && nil => nil
		test(Foo, Empty, Nil)   // [foo] && [] => nil

		test(Foo, Bar, Nil) // [foo] && [bar] => nil
		test(Bar, Foo, Nil) // [bar] && [foo] => nil

		test(FooBar, Foo, Foo)    // [foo bar] && [foo] => [foo]
		test(Foo, FooBar, Foo)    // [foo] && [foo bar] => [foo]
		test(FooBarBaz, Foo, Foo) // [foo bar baz] && [foo] => [foo]
		test(Foo, FooBarBaz, Foo) // [foo] && [foo bar baz] => [foo]

		test(FooBar, FooBarBaz, FooBar)       // [foo bar] && [foo bar baz] => [foo bar]
		test(FooBarBaz, FooBar, FooBar)       // [foo bar baz] && [foo bar] => [foo bar]
		test(FooBarBaz, FooBarBaz, FooBarBaz) // [foo bar baz] && [foo bar baz] => [foo bar baz]
		test(FooBar, FooBar, FooBar)          // [foo bar] && [foo bar] => [foo bar]
		test(Foo, Foo, Foo)                   // [foo] && [foo] => [foo]
	})

	// integer
	tt.Run("int", func(t *testing.T) {
		var Nil []int
		Empty := []int{}
		Foo := []int{123}
		Bar := []int{456}
		FooBar := []int{123, 456}
		FooBarBaz := []int{123, 456, 789}

		test := func(a, b []int, e1 []int) {
			t.Helper()
			if a1 := slices2.SetAnd(a, b); !equal(a1, e1) {
				t.Errorf("SetAnd(`%#v`,`%#v`)=`%#v`, expected `%#v`", a, b, a1, e1)
			}
			if a1 := slices2.SetAndBy(func(v int) int { return v }, a, b); !equal(a1, e1) {
				t.Errorf("SetAndBy(`%#v`,`%#v`)=`%#v`, expected `%#v`", a, b, a1, e1)
			}
		}

		test(Nil, Nil, Nil)     // nil && nil => nil
		test(Nil, Empty, Nil)   // nil && [] => nil
		test(Empty, Nil, Nil)   // [] && nil => nil
		test(Empty, Empty, Nil) // [] && [] => nil
		test(Nil, Foo, Nil)     // nil && [foo] => nil
		test(Empty, Foo, Nil)   // [] && [foo] => nil
		test(Foo, Nil, Nil)     // [foo] && nil => nil
		test(Foo, Empty, Nil)   // [foo] && [] => nil

		test(Foo, Bar, Nil) // [foo] && [bar] => nil
		test(Bar, Foo, Nil) // [bar] && [foo] => nil

		test(FooBar, Foo, Foo)    // [foo bar] && [foo] => [foo]
		test(Foo, FooBar, Foo)    // [foo] && [foo bar] => [foo]
		test(FooBarBaz, Foo, Foo) // [foo bar baz] && [foo] => [foo]
		test(Foo, FooBarBaz, Foo) // [foo] && [foo bar baz] => [foo]

		test(FooBar, FooBarBaz, FooBar)       // [foo bar] && [foo bar baz] => [foo bar]
		test(FooBarBaz, FooBar, FooBar)       // [foo bar baz] && [foo bar] => [foo bar]
		test(FooBarBaz, FooBarBaz, FooBarBaz) // [foo bar baz] && [foo bar baz] => [foo bar baz]
		test(FooBar, FooBar, FooBar)          // [foo bar] && [foo bar] => [foo bar]
		test(Foo, Foo, Foo)                   // [foo] && [foo] => [foo]
	})
}
