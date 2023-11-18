package slices2_test

import (
	"fmt"
	"strconv"
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
	ii := slices2.Transform(fromStr, ss)
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

			if actual := slices2.Transform(fn, input); !equal(actual, expected) {
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

			if actual := slices2.Transform(fn, input); !equal(actual, expected) {
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
