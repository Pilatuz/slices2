package slices2_test

import (
	"fmt"
	"testing"

	"github.com/Pilatuz/slices2"
)

// ExampleSetSub an example for `SetSub` function.
func ExampleSetSub() {
	a := []string{"foo", "bar"}
	b := []string{"bar", "baz"}

	c := slices2.SetSub(a, b)
	fmt.Println(a, "-", b, "=", c)
	// Output:
	// [foo bar] - [bar baz] = [foo]
}

// TestSetSub unit tests for `SetSub` function.
func TestSetSub(tt *testing.T) {
	// string
	tt.Run("str", func(t *testing.T) {
		var Nil []string
		Empty := []string{}
		Foo := []string{"foo"}
		Bar := []string{"bar"}
		Baz := []string{"baz"}
		FooBar := []string{"foo", "bar"}
		BarBaz := []string{"bar", "baz"}
		FooBarBaz := []string{"foo", "bar", "baz"}

		test := func(a, b []string, e1 []string) {
			t.Helper()
			if a1 := slices2.SetSub(a, b); !equal(a1, e1) {
				t.Errorf("SetSub(`%#v`,`%#v`)=`%#v`, expected `%#v`", a, b, a1, e1)
			}
			if a1 := slices2.SetSubBy(a, b, func(v string) string { return v }); !equal(a1, e1) {
				t.Errorf("SetSubBy(`%#v`,`%#v`)=`%#v`, expected `%#v`", a, b, a1, e1)
			}
		}

		test(Nil, Nil, Nil)       // nil - nil => nil
		test(Nil, Empty, Nil)     // nil - [] => nil
		test(Empty, Nil, Empty)   // [] - nil => []
		test(Empty, Empty, Empty) // [] - [] => []
		test(Nil, Foo, Nil)       // nil - [foo] => nil
		test(Empty, Foo, Nil)     // [] - [foo] => nil
		test(Foo, Nil, Foo)       // [foo] - nil => [foo]
		test(Foo, Empty, Foo)     // [foo] - [] => [foo]

		test(Foo, Bar, Foo) // [foo] - [bar] => [foo]
		test(Bar, Foo, Bar) // [bar] - [foo] => [bar]

		test(FooBar, Foo, Bar)       // [foo bar] - [foo] => [bar]
		test(FooBar, Bar, Foo)       // [foo bar] - [bar] => [foo]
		test(Foo, FooBar, Nil)       // [foo] - [foo bar] => nil
		test(FooBarBaz, Foo, BarBaz) // [foo bar baz] - [foo] => [bar baz]
		test(Foo, FooBarBaz, nil)    // [foo] - [foo bar baz] => nil

		test(FooBar, FooBarBaz, Nil)    // [foo bar] - [foo bar baz] => nil
		test(FooBarBaz, FooBar, Baz)    // [foo bar baz] - [foo bar] => [baz]
		test(FooBarBaz, FooBarBaz, Nil) // [foo bar baz] - [foo bar baz] => nil
		test(FooBar, FooBar, Nil)       // [foo bar] && [foo bar] => nil
		test(Foo, Foo, Nil)             // [foo] && [foo] => nil
	})

	// integer
	tt.Run("int", func(t *testing.T) {
		var Nil []int
		Empty := []int{}
		Foo := []int{123}
		Bar := []int{456}
		Baz := []int{789}
		FooBar := []int{123, 456}
		BarBaz := []int{456, 789}
		FooBarBaz := []int{123, 456, 789}

		test := func(a, b []int, e1 []int) {
			t.Helper()
			if a1 := slices2.SetSub(a, b); !equal(a1, e1) {
				t.Errorf("SetSub(`%#v`,`%#v`)=`%#v`, expected `%#v`", a, b, a1, e1)
			}
			if a1 := slices2.SetSubBy(a, b, func(v int) int { return v }); !equal(a1, e1) {
				t.Errorf("SetSubBy(`%#v`,`%#v`)=`%#v`, expected `%#v`", a, b, a1, e1)
			}
		}

		test(Nil, Nil, Nil)       // nil - nil => nil
		test(Nil, Empty, Nil)     // nil - [] => nil
		test(Empty, Nil, Empty)   // [] - nil => []
		test(Empty, Empty, Empty) // [] - [] => []
		test(Nil, Foo, Nil)       // nil - [foo] => nil
		test(Empty, Foo, Nil)     // [] - [foo] => nil
		test(Foo, Nil, Foo)       // [foo] - nil => [foo]
		test(Foo, Empty, Foo)     // [foo] - [] => [foo]

		test(Foo, Bar, Foo) // [foo] - [bar] => [foo]
		test(Bar, Foo, Bar) // [bar] - [foo] => [bar]

		test(FooBar, Foo, Bar)       // [foo bar] - [foo] => [bar]
		test(FooBar, Bar, Foo)       // [foo bar] - [bar] => [foo]
		test(Foo, FooBar, Nil)       // [foo] - [foo bar] => nil
		test(FooBarBaz, Foo, BarBaz) // [foo bar baz] - [foo] => [bar baz]
		test(Foo, FooBarBaz, nil)    // [foo] - [foo bar baz] => nil

		test(FooBar, FooBarBaz, Nil)    // [foo bar] - [foo bar baz] => nil
		test(FooBarBaz, FooBar, Baz)    // [foo bar baz] - [foo bar] => [baz]
		test(FooBarBaz, FooBarBaz, Nil) // [foo bar baz] - [foo bar baz] => nil
		test(FooBar, FooBar, Nil)       // [foo bar] && [foo bar] => nil
		test(Foo, Foo, Nil)             // [foo] && [foo] => nil
	})
}
