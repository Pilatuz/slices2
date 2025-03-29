package slices2_test

import (
	"fmt"
	"strconv"
	"strings"
	"testing"

	"github.com/Pilatuz/slices2"
)

// ExampleTransform an example for `Transform` function.
func ExampleTransform() {
	fromStr := func(s string) int {
		out, _ := strconv.Atoi(s)
		return out
	}
	ss := []string{"123", "456"}
	ii := slices2.Transform(ss, fromStr)
	fmt.Println(ii)
	// Output:
	// [123 456]
}

// TestTransform unit tests for `Transform` function.
func TestTransform(tt *testing.T) {
	// string => int
	tt.Run("str_int", func(t *testing.T) {
		var Nil []string
		Empty := []string{}
		FooBar := []string{"123", "456"}

		fn := func(s string) int {
			i, _ := strconv.Atoi(s)
			return i
		}

		test := func(input []string, expected []int) {
			t.Helper()

			if actual := slices2.Transform(input, fn); !equal(actual, expected) {
				t.Errorf("Transform(`%#v`)=`%#v`, expected `%#v`", input, actual, expected)
			}
		}

		test(Nil, nil)       // Transform(nil) => nil
		test(Empty, []int{}) // Transform([]) => []

		// Transform(["a", "b", ...]) gives slice [a, b, ...]
		test(FooBar[:1], []int{123})
		test(FooBar[1:], []int{456})
		test(FooBar, []int{123, 456})
	})

	// string => boolean
	tt.Run("str_bool", func(t *testing.T) {
		var Nil []string
		Empty := []string{}
		FooBar := []string{"true", "false"}

		fn := func(s string) bool {
			b, _ := strconv.ParseBool(s)
			return b
		}

		test := func(input []string, expected []bool) {
			t.Helper()

			if actual := slices2.Transform(input, fn); !equal(actual, expected) {
				t.Errorf("Transform(`%#v`)=`%#v`, expected `%#v`", input, actual, expected)
			}
		}

		test(Nil, nil)        // Transform(nil) => nil
		test(Empty, []bool{}) // Transform([]) => []

		// Transform(["a", "b", ...]) gives slice [a, b, ...]
		test(FooBar[:1], []bool{true})
		test(FooBar[1:], []bool{false})
		test(FooBar, []bool{true, false})
	})
}

// TestTransformEx unit tests for `TransformEx` function.
func TestTransformEx(tt *testing.T) {
	// string => int
	tt.Run("str_int", func(t *testing.T) {
		var Nil []string
		Empty := []string{}
		FooBar := []string{"123", "456"}

		test := func(input []string, expected []int) {
			t.Helper()

			if actual, err := slices2.TransformEx(input, strconv.Atoi); err != nil {
				t.Errorf("TransformEx(`%#v`)=`%#v`, failed with `%v`", input, actual, err)
			} else if !equal(actual, expected) {
				t.Errorf("TransformEx(`%#v`)=`%#v`, expected `%#v`", input, actual, expected)
			}
		}

		test(Nil, nil)       // Transform(nil) => nil
		test(Empty, []int{}) // Transform([]) => []

		// Transform(["a", "b", ...]) gives slice [a, b, ...]
		test(FooBar[:1], []int{123})
		test(FooBar[1:], []int{456})
		test(FooBar, []int{123, 456})
	})

	tt.Run("str_int_bad", func(t *testing.T) {
		Bad := []string{"123", "BAD", "456"}

		// stop on first error
		if actual, err := slices2.TransformEx(Bad, strconv.Atoi); err == nil {
			t.Errorf("TransformEx(`%#v`)=`%#v`, error expected", Bad, actual)
		} else if actual != nil {
			t.Errorf("TransformEx(`%#v`)=`%#v`, no result expected expected", Bad, actual)
		} else if !strings.Contains(err.Error(), "invalid syntax") {
			t.Errorf("TransformEx(`%#v`)=`%#v`, another error expected `%v`", Bad, actual, err)
		}

		skipBad := func(s string) (int, error) {
			i, err := strconv.Atoi(s)
			if err != nil {
				return 0, slices2.ErrSkip
			}
			return i, nil
		}

		// skip parsing errors
		expected := []int{123, 456}
		if actual, err := slices2.TransformEx(Bad, skipBad); err != nil {
			t.Errorf("TransformEx(`%#v`)=`%#v`, failed with `%v`", Bad, actual, err)
		} else if !equal(actual, expected) {
			t.Errorf("TransformEx(`%#v`)=`%#v`, expected `%#v`", Bad, actual, expected)
		}
	})

	// string => boolean
	tt.Run("str_bool", func(t *testing.T) {
		var Nil []string
		Empty := []string{}
		FooBar := []string{"true", "false"}

		test := func(input []string, expected []bool) {
			t.Helper()

			if actual, err := slices2.TransformEx(input, strconv.ParseBool); err != nil {
				t.Errorf("TransformEx(`%#v`)=`%#v`, failed with `%#v`", input, actual, err)
			} else if !equal(actual, expected) {
				t.Errorf("TransformEx(`%#v`)=`%#v`, expected `%#v`", input, actual, expected)
			}
		}

		test(Nil, nil)        // Transform(nil) => nil
		test(Empty, []bool{}) // Transform([]) => []

		// Transform(["a", "b", ...]) gives slice [a, b, ...]
		test(FooBar[:1], []bool{true})
		test(FooBar[1:], []bool{false})
		test(FooBar, []bool{true, false})
	})

	// *string => string
	tt.Run("str_ptr", func(t *testing.T) {
		foo, bar := "foo", "bar"

		// skip nils
		input := []*string{nil, &foo, nil, &bar, nil}
		expected := []string{foo, bar}
		if actual, err := slices2.TransformEx(input, slices2.Deref[string]); err != nil {
			t.Errorf("TransformEx(`%#v`)=`%#v`, failed with `%v`", input, actual, err)
		} else if !equal(actual, expected) {
			t.Errorf("TransformEx(`%#v`)=`%#v`, expected `%#v`", input, actual, expected)
		}
	})
}
