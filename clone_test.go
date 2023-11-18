package slices2_test

import (
	"testing"

	"github.com/Pilatuz/slices2"
)

// TestClone unit tests for `clone` helper.
func TestClone(tt *testing.T) {
	// string
	tt.Run("str", func(t *testing.T) {
		var Nil []string
		Empty := []string{}
		Foo := []string{"foo"}
		FooBar := []string{"foo", "bar"}

		test := func(a []string) {
			t.Helper()

			if b := slices2.Clone(a); same(a, b) || !equal(a, b) {
				t.Errorf("clone(`%#v`)=`%#v`, expected `%#v`", a, b, a)
			}
		}

		test(Nil)
		test(Empty)
		test(Foo)
		test(FooBar)
	})

	// integer
	tt.Run("int", func(t *testing.T) {
		var Nil []int
		Empty := []int{}
		Foo := []int{123}
		FooBar := []int{123, 456}

		test := func(a []int) {
			t.Helper()

			if b := slices2.Clone(a); same(a, b) || !equal(a, b) {
				t.Errorf("clone(`%#v`)=`%#v`, expected `%#v`", a, b, a)
			}
		}

		test(Nil)
		test(Empty)
		test(Foo)
		test(FooBar)
	})
}
