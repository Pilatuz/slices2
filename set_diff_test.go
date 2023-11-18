package slices2_test

import (
	"testing"

	"github.com/Pilatuz/slices2"
)

// TestSetDiff unit tests for `SetDiff` function.
func TestSetDiff(tt *testing.T) {
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

		test := func(a, b []string, e1, e2 []string) {
			t.Helper()
			if a1, a2 := slices2.SetDiff(a, b); !equal(a1, e1) || !equal(a2, e2) {
				t.Errorf("SetDiff(`%#v`,`%#v`)=(`%#v`,`%#v`), expected (`%#v`,`%#v`)", a, b, a1, a2, e1, e2)
			}
			if a1, a2 := slices2.SetDiffBy(func(v string) string { return v }, a, b); !equal(a1, e1) || !equal(a2, e2) {
				t.Errorf("SetDiffBy(`%#v`,`%#v`)=(`%#v`,`%#v`), expected (`%#v`,`%#v`)", a, b, a1, a2, e1, e2)
			}
		}

		test(Nil, Nil, Nil, Nil)         // diff(nil, nil) => nil, nil
		test(Nil, Empty, Nil, Empty)     // diff(nil, [])  => nil, []
		test(Empty, Nil, Empty, Nil)     // diff([], nil)  => [], nil
		test(Empty, Empty, Empty, Empty) // diff([], [])   => [], []

		test(Nil, Foo, Nil, Foo)     // diff(nil, [foo]) => nil, [foo]
		test(Empty, Foo, Empty, Foo) // diff([], [foo])  => [], [foo]
		test(Foo, Nil, Foo, Nil)     // diff([foo], nil) => [foo], nil
		test(Foo, Empty, Foo, Empty) // diff([foo], [])  => [foo], []

		test(Bar, Foo, Bar, Foo) // diff([bar], [foo]) => [bar], [foo]
		test(Foo, Bar, Foo, Bar) // diff([foo], [bar]) => [foo], [bar]
		test(Baz, Foo, Baz, Foo) // diff([baz], [foo]) => [baz], [foo]
		test(Foo, Baz, Foo, Baz) // diff([foo], [baz]) => [foo], [baz]

		test(FooBar, Foo, Bar, Nil)       // diff([foo bar], [foo]) => [bar], nil
		test(Foo, FooBar, Nil, Bar)       // diff([foo], [foo bar]) => nil, [bar]
		test(FooBarBaz, Foo, BarBaz, Nil) // diff([foo bar baz], [foo]) => [bar baz], nil
		test(Foo, FooBarBaz, Nil, BarBaz) // diff([foo], [foo bar baz]) => nil, [bar baz]

		test(FooBar, BarBaz, Foo, Baz)    // diff([foo bar], [bar baz]) => [foo], [baz]
		test(BarBaz, FooBar, Baz, Foo)    // diff([bar baz], [foo bar]) => [baz], [foo]
		test(FooBarBaz, FooBar, Baz, Nil) // diff([foo bar baz], [foo bar]) => [baz], nil
		test(FooBar, FooBarBaz, Nil, Baz) // diff([foo bar], [foo bar baz]) => nil, [baz]

		test(FooBarBaz, FooBarBaz, Nil, Nil) // diff([foo bar baz], [foo bar baz]) => nil, nil
		test(FooBar, FooBar, Nil, Nil)       // diff([foo bar], [foo bar]) => nil, nil
		test(BarBaz, BarBaz, Nil, Nil)       // diff([bar baz], [bar baz]) => nil, nil
		test(Foo, Foo, Nil, Nil)             // diff([foo], [foo]) => nil, nil
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

		test := func(a, b []int, e1, e2 []int) {
			t.Helper()
			if a1, a2 := slices2.SetDiff(a, b); !equal(a1, e1) || !equal(a2, e2) {
				t.Errorf("SetDiff(`%#v`,`%#v`)=(`%#v`,`%#v`), expected (`%#v`,`%#v`)", a, b, a1, a2, e1, e2)
			}
			if a1, a2 := slices2.SetDiffBy(func(v int) int { return v }, a, b); !equal(a1, e1) || !equal(a2, e2) {
				t.Errorf("SetDiffBy(`%#v`,`%#v`)=(`%#v`,`%#v`), expected (`%#v`,`%#v`)", a, b, a1, a2, e1, e2)
			}
		}

		test(Nil, Nil, Nil, Nil)         // diff(nil, nil) => nil, nil
		test(Nil, Empty, Nil, Empty)     // diff(nil, [])  => nil, []
		test(Empty, Nil, Empty, Nil)     // diff([], nil)  => [], nil
		test(Empty, Empty, Empty, Empty) // diff([], [])   => [], []

		test(Nil, Foo, Nil, Foo)     // diff(nil, [foo]) => nil, [foo]
		test(Empty, Foo, Empty, Foo) // diff([], [foo])  => [], [foo]
		test(Foo, Nil, Foo, Nil)     // diff([foo], nil) => [foo], nil
		test(Foo, Empty, Foo, Empty) // diff([foo], [])  => [foo], []

		test(Bar, Foo, Bar, Foo) // diff([bar], [foo]) => [bar], [foo]
		test(Foo, Bar, Foo, Bar) // diff([foo], [bar]) => [foo], [bar]
		test(Baz, Foo, Baz, Foo) // diff([baz], [foo]) => [baz], [foo]
		test(Foo, Baz, Foo, Baz) // diff([foo], [baz]) => [foo], [baz]

		test(FooBar, Foo, Bar, Nil)       // diff([foo bar], [foo]) => [bar], nil
		test(Foo, FooBar, Nil, Bar)       // diff([foo], [foo bar]) => nil, [bar]
		test(FooBarBaz, Foo, BarBaz, Nil) // diff([foo bar baz], [foo]) => [bar baz], nil
		test(Foo, FooBarBaz, Nil, BarBaz) // diff([foo], [foo bar baz]) => nil, [bar baz]

		test(FooBar, BarBaz, Foo, Baz)    // diff([foo bar], [bar baz]) => [foo], [baz]
		test(BarBaz, FooBar, Baz, Foo)    // diff([bar baz], [foo bar]) => [baz], [foo]
		test(FooBarBaz, FooBar, Baz, Nil) // diff([foo bar baz], [foo bar]) => [baz], nil
		test(FooBar, FooBarBaz, Nil, Baz) // diff([foo bar], [foo bar baz]) => nil, [baz]

		test(FooBarBaz, FooBarBaz, Nil, Nil) // diff([foo bar baz], [foo bar baz]) => nil, nil
		test(FooBar, FooBar, Nil, Nil)       // diff([foo bar], [foo bar]) => nil, nil
		test(BarBaz, BarBaz, Nil, Nil)       // diff([bar baz], [bar baz]) => nil, nil
		test(Foo, Foo, Nil, Nil)             // diff([foo], [foo]) => nil, nil
	})
}
