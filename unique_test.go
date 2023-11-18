package slices2_test

import (
	"fmt"
	"testing"

	"github.com/Pilatuz/slices2"
)

// ExampleUnique an example for `Unique` function.
func ExampleUnique() {
	s := []int{4, 1, 2, 1, 2, 3, 2, 1}
	fmt.Println(slices2.Unique(s))
	// Output:
	// [4 1 2 3]
}

// TestUnique unit tests for Unique* set of functions.
func TestUnique(tt *testing.T) {
	// string
	tt.Run("str", func(t *testing.T) {
		var Nil []string
		Empty := []string{}
		Foo := []string{"foo"}
		FooFoo := []string{"foo", "foo"}
		FooBar := []string{"foo", "bar"}
		BarFoo := []string{"bar", "foo"}
		FooBarBar := []string{"foo", "bar", "bar", "foo"}

		test := func(input, expected []string) {
			t.Helper()

			if actual := slices2.Unique(input); !equal(actual, expected) || same(actual, input) {
				t.Errorf("Unique(`%#v`)=`%#v`, expected `%#v`", input, actual, expected)
			}
			if actual := slices2.UniqueBy(func(v string) string { return v }, input); !equal(actual, expected) || same(actual, input) {
				t.Errorf("UniqueBy(`%#v`)=`%#v`, expected `%#v`", input, actual, expected)
			}

			input2 := slices2.Clone(input)
			if actual := slices2.UniqueInPlace(input2); !equal(actual, expected) {
				t.Errorf("UniqueInPlace(`%#v`)=`%#v`, expected `%#v`", input, actual, expected)
			} else if len(actual) > 0 && !same(actual, input2) {
				t.Errorf("UniqueInPlace(`%#v`)=`%#v`, expected `%#v`", input, actual, expected)
			}

			input3 := slices2.Clone(input)
			if actual := slices2.UniqueInPlaceBy(func(v string) string { return v }, input3); !equal(actual, expected) {
				t.Errorf("UniqueInPlaceBy(`%#v`)=`%#v`, expected `%#v`", input, actual, expected)
			} else if len(actual) > 0 && !same(actual, input3) {
				t.Errorf("UniqueInPlaceBy(`%#v`)=`%#v`, expected `%#v`", input, actual, expected)
			}
		}

		test(Nil, Nil)       // Unique(nil) => nil
		test(Empty, Empty)   // Unique([]) => []
		test(Foo, Foo)       // Unique([foo]) => [foo]
		test(FooFoo, Foo)    // Unique([foo foo]) => [foo]
		test(FooBar, FooBar) // Unique([foo bar]) => [foo bar] (keeps original elements order!)
		test(BarFoo, BarFoo) // Unique([bar foo]) => [bar foo] (keeps original elements order!)
		test(FooBarBar, FooBar)
	})

	// integer
	tt.Run("int", func(t *testing.T) {
		var Nil []int
		Empty := []int{}
		Foo := []int{123}
		FooFoo := []int{123, 123}
		FooBar := []int{123, 456}
		BarFoo := []int{456, 123}
		FooBarBar := []int{123, 456, 456, 123}

		test := func(input, expected []int) {
			t.Helper()

			if actual := slices2.Unique(input); !equal(actual, expected) || same(actual, input) {
				t.Errorf("Unique(`%#v`)=`%#v`, expected `%#v`", input, actual, expected)
			}
			if actual := slices2.UniqueBy(func(v int) int { return v }, input); !equal(actual, expected) || same(actual, input) {
				t.Errorf("UniqueBy(`%#v`)=`%#v`, expected `%#v`", input, actual, expected)
			}

			input2 := slices2.Clone(input)
			if actual := slices2.UniqueInPlace(input2); !equal(actual, expected) {
				t.Errorf("UniqueInPlace(`%#v`)=`%#v`, expected `%#v`", input, actual, expected)
			} else if len(actual) > 0 && !same(actual, input2) {
				t.Errorf("UniqueInPlace(`%#v`)=`%#v`, expected `%#v`", input, actual, expected)
			}

			input3 := slices2.Clone(input)
			if actual := slices2.UniqueInPlaceBy(func(v int) int { return v }, input3); !equal(actual, expected) {
				t.Errorf("UniqueInPlaceBy(`%#v`)=`%#v`, expected `%#v`", input, actual, expected)
			} else if len(actual) > 0 && !same(actual, input3) {
				t.Errorf("UniqueInPlaceBy(`%#v`)=`%#v`, expected `%#v`", input, actual, expected)
			}
		}

		test(Nil, Nil)       // Unique(nil) => nil
		test(Empty, Empty)   // Unique([]) => []
		test(Foo, Foo)       // Unique([foo]) => [foo]
		test(FooFoo, Foo)    // Unique([foo foo]) => [foo]
		test(FooBar, FooBar) // Unique([foo bar]) => [foo bar] (keeps original elements order!)
		test(BarFoo, BarFoo) // Unique([bar foo]) => [bar foo] (keeps original elements order!)
		test(FooBarBar, FooBar)
	})

	// struct
	tt.Run("struct", func(t *testing.T) {
		type User struct {
			ID   int
			Name string
		}

		var Nil []User
		Empty := []User{}
		Foo := []User{{1, "Foo"}}
		FooFoo := []User{{1, "Foo"}, {2, "Foo"}}
		FooBar := []User{{1, "Foo"}, {2, "Bar"}}

		test := func(input, expected []User) {
			t.Helper()

			if actual := slices2.Unique(input); !equal(actual, expected) || same(actual, input) {
				t.Errorf("Unique(`%#v`)=`%#v`, expected `%#v`", input, actual, expected)
			}

			input2 := slices2.Clone(input)
			if actual := slices2.UniqueInPlace(input2); !equal(actual, expected) {
				t.Errorf("UniqueInPlace(`%#v`)=`%#v`, expected `%#v`", input, actual, expected)
			} else if len(actual) > 0 && !same(actual, input2) {
				t.Errorf("UniqueInPlace(`%#v`)=`%#v`, expected `%#v`", input, actual, expected)
			}
		}

		testBy := func(input, expected []User) {
			t.Helper()

			byName := func(v User) string {
				return v.Name
			}

			if actual := slices2.UniqueBy(byName, input); !equal(actual, expected) || same(actual, input) {
				t.Errorf("UniqueBy(`%#v`)=`%#v`, expected `%#v`", input, actual, expected)
			}

			input2 := slices2.Clone(input)
			if actual := slices2.UniqueInPlaceBy(byName, input2); !equal(actual, expected) {
				t.Errorf("UniqueInPlaceBy(`%#v`)=`%#v`, expected `%#v`", input, actual, expected)
			} else if len(actual) > 0 && !same(actual, input2) {
				t.Errorf("UniqueInPlaceBy(`%#v`)=`%#v`, expected `%#v`", input, actual, expected)
			}
		}

		test(Nil, Nil)       // Unique(nil) => nil
		test(Empty, Empty)   // Unique([]) => []
		test(Foo, Foo)       // Unique([foo]) => [foo]
		test(FooFoo, FooFoo) // Unique([foo foo]) => [foo foo] (various ids)
		test(FooBar, FooBar) // Unique([foo bar]) => [foo bar] (keeps original elements order!)

		testBy(Nil, Nil)       // Unique(nil) => nil
		testBy(Empty, Empty)   // Unique([]) => []
		testBy(Foo, Foo)       // Unique([foo]) => [foo]
		testBy(FooFoo, Foo)    // Unique([foo foo]) => [foo] (same name)
		testBy(FooBar, FooBar) // Unique([foo bar]) => [foo bar] (keeps original elements order!)
	})
}
