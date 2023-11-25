package slices2_test

import (
	"reflect"
	"strings"
	"testing"

	"github.com/Pilatuz/slices2"
)

// TestGroupBy unit tests for `GroupBy` function.
func TestGroupBy(tt *testing.T) {
	// string
	tt.Run("str", func(t *testing.T) {
		var Nil []string
		Empty := []string{}

		isCapital := func(v string) bool {
			return strings.ToUpper(v) == v
		}

		test := func(x []string, e map[bool][]string) {
			t.Helper()
			if a := slices2.GroupBy(isCapital, x); !reflect.DeepEqual(a, e) {
				t.Errorf("GroupBy(`%#v`)=`%#v`, expected `%#v`", x, a, e)
			}
		}

		test(Nil, map[bool][]string{})
		test(Empty, map[bool][]string{})

		test([]string{"foo"},
			map[bool][]string{
				false: {"foo"},
			})
		test([]string{"foo", "FOO"},
			map[bool][]string{
				false: {"foo"},
				true:  {"FOO"},
			})
		test([]string{"foo", "FOO", "bar", "BAR", "baz"},
			map[bool][]string{
				false: {"foo", "bar", "baz"},
				true:  {"FOO", "BAR"},
			})
	})

	// integer
	tt.Run("int", func(t *testing.T) {
		var Nil []int
		Empty := []int{}

		isEven := func(v int) int {
			return v & 1
		}

		test := func(x []int, e map[int][]int) {
			t.Helper()
			if a := slices2.GroupBy(isEven, x); !reflect.DeepEqual(a, e) {
				t.Errorf("GroupBy(`%#v`)=`%#v`, expected `%#v`", x, a, e)
			}
		}

		test(Nil, map[int][]int{})
		test(Empty, map[int][]int{})

		test([]int{1},
			map[int][]int{
				1: {1},
			})
		test([]int{1, 2},
			map[int][]int{
				0: {2},
				1: {1},
			})
		test([]int{1, 2, 3, 4, 5},
			map[int][]int{
				0: {2, 4},
				1: {1, 3, 5},
			})
	})
}
