package slices2_test

import (
	"fmt"
	"testing"

	"github.com/Pilatuz/slices2"
)

// ExampleJoin an example for `Join` function.
func ExampleJoin() {
	s := slices2.Join(
		[]string{"foo", "bar"},
		[]string{"baz"})
	fmt.Println(s)
	// Output:
	// [foo bar baz]
}

// TestJoin unit tests for `Join` function.
func TestJoin(tt *testing.T) {
	// string
	tt.Run("str", func(t *testing.T) {
		var Nil []string
		Empty := []string{}
		Foo := []string{"foo"}
		Bar := []string{"bar"}
		FooBar := []string{"foo", "bar"}

		test := func(expected []string, aa ...[]string) {
			t.Helper()

			if actual := slices2.Join(aa...); !equal(actual, expected) {
				t.Errorf("Join(`%#v`)=`%#v`, expected `%#v`", aa, actual, expected)
			}
		}

		test(Nil)                    // Join() => nil
		test(Nil, Nil)               // Join(nil) => nil
		test(Empty, Empty)           // Join([]) => nil
		test(Nil, Empty, Nil, Empty) // Join([], nil, []) => nil

		test(Foo, Nil, Foo, Empty)         // Join(nil, [foo], []) => [foo]
		test(Bar, Nil, Bar, Empty)         // Join(nil, [bar], []) => [bar]
		test(FooBar, Nil, Foo, Bar, Empty) // Join(nil, [foo] [bar], []) => [foo bar]
	})

	// integer
	tt.Run("int", func(t *testing.T) {
		var Nil []int
		Empty := []int{}
		Foo := []int{123}
		Bar := []int{456}
		FooBar := []int{123, 456}

		test := func(expected []int, aa ...[]int) {
			t.Helper()

			if actual := slices2.Join(aa...); !equal(actual, expected) {
				t.Errorf("Join(`%#v`)=`%#v`, expected `%#v`", aa, actual, expected)
			}
		}

		test(Nil)                    // Join() => nil
		test(Nil, Nil)               // Join(nil) => nil
		test(Empty, Empty)           // Join([]) => nil
		test(Nil, Empty, Nil, Empty) // Join([], nil, []) => nil

		test(Foo, Nil, Foo, Empty)         // Join(nil, [foo], []) => [foo]
		test(Bar, Nil, Bar, Empty)         // Join(nil, [bar], []) => [bar]
		test(FooBar, Nil, Foo, Bar, Empty) // Join(nil, [foo] [bar], []) => [foo bar]
	})
}
