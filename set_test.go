package slices2_test

import (
	"fmt"
	"reflect"
	"testing"

	"github.com/Pilatuz/slices2"
)

// ExampleSet an example for Set type and its methods.
func ExampleSet() {
	// Create a new set
	s := slices2.NewSet("foo", "bar", "baz")
	fmt.Println("Set:", s)

	// Check if element exists
	fmt.Println("Has(foo):", s.Has("foo"))
	fmt.Println("Has(qux):", s.Has("qux"))

	// Add elements
	s.Push("qux")
	fmt.Println("After Push(qux):", s)

	// Remove elements
	s.Pop("bar")
	fmt.Println("After Pop(bar):", s)

	// Check all elements (order is not guaranteed, so sort for predictable output)
	fmt.Println("All():", slices2.Sorted(s.All()))

	// Check HasAll
	fmt.Println("HasAll(foo, qux):", s.HasAll("foo", "qux"))
	fmt.Println("HasAll(foo, bar):", s.HasAll("foo", "bar"))

	// Check HasAny
	fmt.Println("HasAny(foo, bar):", s.HasAny("foo", "bar"))
	fmt.Println("HasAny(qux, quux):", s.HasAny("qux", "quux"))
	// Output:
	// Set: map[bar:{} baz:{} foo:{}]
	// Has(foo): true
	// Has(qux): false
	// After Push(qux): map[bar:{} baz:{} foo:{} qux:{}]
	// After Pop(bar): map[baz:{} foo:{} qux:{}]
	// All(): [baz foo qux]
	// HasAll(foo, qux): true
	// HasAll(foo, bar): false
	// HasAny(foo, bar): true
	// HasAny(qux, quux): true
}

// ExampleNewSetBy an example for NewSetBy function.
func ExampleNewSetBy() {
	type Person struct {
		Name string
		Age  int
	}
	people := []Person{{"Alice", 30}, {"Bob", 25}, {"Charlie", 30}}
	// Group by age
	ages := slices2.NewSetBy(people, func(p Person) int { return p.Age })
	fmt.Println(ages)
	// Output:
	// map[25:{} 30:{}]
}

// TestSet unit tests for set.
func TestSet(t *testing.T) {
	t.Run("NewSet_empty", func(t *testing.T) {
		a := slices2.NewSet[int]()
		expected := slices2.Set[int]{}
		if actual := a; !reflect.DeepEqual(actual, expected) {
			t.Errorf("NewSet()=`%#v`, expected `%#v`", actual, expected)
		}
	})

	t.Run("NewSet_int", func(t *testing.T) {
		a := slices2.NewSet(1, 2, 3, 2, 1)
		expected := slices2.Set[int]{1: {}, 2: {}, 3: {}}
		if actual := a; !reflect.DeepEqual(actual, expected) {
			t.Errorf("NewSet(1, 2, 3, 2, 1)=`%#v`, expected `%#v`", actual, expected)
		}
	})

	t.Run("NewSet_strings", func(t *testing.T) {
		a := slices2.NewSet("foo", "bar", "foo")
		expected := slices2.Set[string]{"foo": {}, "bar": {}}
		if actual := a; !reflect.DeepEqual(actual, expected) {
			t.Errorf("NewSet(`foo`, `bar`, `foo`)=`%#v`, expected `%#v`", actual, expected)
		}
	})

	t.Run("Push_1by1", func(t *testing.T) {
		a := make(slices2.Set[string], 10) // bigger capacity!
		a.Push("foo")
		a.Push("bar")
		a.Push("foo") // duplicate
		expected := slices2.Set[string]{"foo": {}, "bar": {}}
		if actual := a; !reflect.DeepEqual(actual, expected) {
			t.Errorf("Push(`foo`, `bar`, `foo`)=`%#v`, expected `%#v`", actual, expected)
		}
	})

	t.Run("Push_many", func(t *testing.T) {
		a := make(slices2.Set[int], 10) // bigger capacity
		a.Push(1, 2, 3, 2, 1)
		expected := slices2.Set[int]{1: {}, 2: {}, 3: {}}
		if actual := a; !reflect.DeepEqual(actual, expected) {
			t.Errorf("Push(1, 2, 3, 2, 1)=`%#v`, expected `%#v`", actual, expected)
		}
	})

	t.Run("Push_none", func(t *testing.T) {
		a := make(slices2.Set[string])
		a.Push() // no args
		expected := slices2.Set[string]{}
		if actual := a; !reflect.DeepEqual(actual, expected) {
			t.Errorf("Push()=`%#v`, expected `%#v`", actual, expected)
		}
	})

	t.Run("Pop", func(t *testing.T) {
		a := slices2.NewSet("foo", "bar", "baz")
		a.Pop("foo")
		expected := slices2.Set[string]{"bar": {}, "baz": {}}
		if actual := a; !reflect.DeepEqual(actual, expected) {
			t.Errorf("Pop(`foo`)=`%#v`, expected `%#v`", actual, expected)
		}
	})

	t.Run("Pop_many", func(t *testing.T) {
		a := slices2.NewSet(1, 2, 3, 4)
		a.Pop(1, 2)
		expected := slices2.Set[int]{3: {}, 4: {}}
		if actual := a; !reflect.DeepEqual(actual, expected) {
			t.Errorf("Pop(1, 2)=`%#v`, expected `%#v`", actual, expected)
		}
	})

	t.Run("Pop_missing", func(t *testing.T) {
		a := slices2.NewSet("foo", "bar")
		a.Pop("baz") // not in set
		expected := slices2.Set[string]{"foo": {}, "bar": {}}
		if actual := a; !reflect.DeepEqual(actual, expected) {
			t.Errorf("Pop(`baz`)=`%#v`, expected `%#v`", actual, expected)
		}
	})

	t.Run("Pop_missing2", func(t *testing.T) {
		var a slices2.Set[string]
		a.Pop("baz") // not in set
		expected := slices2.Set[string](nil)
		if actual := a; !reflect.DeepEqual(actual, expected) {
			t.Errorf("Pop(`baz`)=`%#v`, expected `%#v`", actual, expected)
		}
	})

	t.Run("Pop_none", func(t *testing.T) {
		a := slices2.NewSet("foo", "bar")
		a.Pop() // no args
		expected := slices2.Set[string]{"foo": {}, "bar": {}}
		if actual := a; !reflect.DeepEqual(actual, expected) {
			t.Errorf("Pop()=`%#v`, expected `%#v`", actual, expected)
		}
	})

	t.Run("Pop_all", func(t *testing.T) {
		a := slices2.NewSet("foo", "bar")
		a.Pop("foo", "bar")
		expected := slices2.Set[string]{}
		if actual := a; !reflect.DeepEqual(actual, expected) {
			t.Errorf("Pop(`foo`, `bar`)=`%#v`, expected `%#v`", actual, expected)
		}
	})

	t.Run("Has", func(t *testing.T) {
		a := slices2.NewSet("foo", "bar")
		if !a.Has("foo") {
			t.Errorf("Has(`foo`)=false, expected true")
		}
		if !a.Has("bar") {
			t.Errorf("Has(`bar`)=false, expected true")
		}
		if a.Has("baz") {
			t.Errorf("Has(`baz`)=true, expected false")
		}
	})

	t.Run("Has_on_empty", func(t *testing.T) {
		a := slices2.NewSet[string]()
		if a.Has("foo") {
			t.Errorf("Has(`foo`)=true, expected false")
		}
	})

	t.Run("Has_on_nil", func(t *testing.T) {
		var a slices2.Set[string]
		if a.Has("foo") {
			t.Errorf("Has(`foo`)=true, expected false")
		}
	})

	t.Run("HasAll", func(t *testing.T) {
		a := slices2.NewSet("foo", "bar", "baz")
		if !a.HasAll("foo", "bar") {
			t.Errorf("HasAll(`foo`, `bar`)=false, expected true")
		}
		if !a.HasAll("foo", "bar", "baz") {
			t.Errorf("HasAll(`foo`, `bar`, `baz`)=false, expected true")
		}
		if a.HasAll("foo", "qux") {
			t.Errorf("HasAll(`foo`, `qux`)=true, expected false")
		}
		if !a.HasAll() {
			t.Errorf("HasAll()=false, expected true")
		}
	})

	t.Run("HasAll_on_empty", func(t *testing.T) {
		a := slices2.NewSet[string]()
		if !a.HasAll() {
			t.Errorf("HasAll()=false, expected true")
		}
		if a.HasAll("foo") {
			t.Errorf("HasAll(`foo`)=true, expected false")
		}
	})

	t.Run("HasAll_on_nil", func(t *testing.T) {
		var a slices2.Set[string]
		if !a.HasAll() {
			t.Errorf("HasAll()=false, expected true")
		}
		if a.HasAll("foo") {
			t.Errorf("HasAll(`foo`)=true, expected false")
		}
	})

	t.Run("HasAny", func(t *testing.T) {
		a := slices2.NewSet("foo", "bar", "baz")
		if !a.HasAny("foo", "qux") {
			t.Errorf("HasAny(`foo`, `qux`)=false, expected true (foo in set)")
		}
		if !a.HasAny("qux", "bar") {
			t.Errorf("HasAny(`qux`, `bar`)=false, expected true (bar in set)")
		}
		if !a.HasAny("foo", "bar") {
			t.Errorf("HasAny(`foo`, `bar`)=false, expected true (all elements present)")
		}
		if !a.HasAny("foo", "bar", "baz") {
			t.Errorf("HasAny(`foo`, `bar`, `baz`)=false, expected true (all elements present)")
		}
		if a.HasAny("qux", "quux") {
			t.Errorf("HasAny(`qux`, `quux`)=true, expected false (no elements present)")
		}
		if a.HasAny() {
			t.Errorf("HasAny()=true, expected false")
		}
	})

	t.Run("HasAny_on_empty", func(t *testing.T) {
		a := slices2.NewSet[string]()
		if a.HasAny("foo") {
			t.Errorf("HasAny(`foo`)=true, expected false (element not in empty set)")
		}
		if a.HasAny("foo", "bar") {
			t.Errorf("HasAny(`foo`, `bar`)=true, expected false (element not in empty set)")
		}
		if a.HasAny() {
			t.Errorf("HasAny()=true, expected false")
		}
	})

	t.Run("HasAny_on_nil", func(t *testing.T) {
		var a slices2.Set[string]
		if a.HasAny("foo") {
			t.Errorf("HasAny(`foo`)=true, expected false (element not in empty set)")
		}
		if a.HasAny("foo", "bar") {
			t.Errorf("HasAny(`foo`, `bar`)=true, expected false (element not in empty set)")
		}
		if a.HasAny() {
			t.Errorf("HasAny()=true, expected false")
		}
	})

	t.Run("All", func(t *testing.T) {
		a := slices2.NewSet("foo", "bar", "baz")
		expected := []string{"bar", "baz", "foo"}
		if actual := slices2.Sorted(a.All()); !equal(actual, expected) {
			t.Errorf("All()=`%#v`, expected `%#v`", actual, expected)
		}
	})

	t.Run("All_on_empty", func(t *testing.T) {
		a := slices2.NewSet[string]()
		expected := []string(nil)
		if actual := slices2.Sorted(a.All()); !equal(actual, expected) {
			t.Errorf("All()=`%#v`, expected `%#v`", actual, expected)
		}
	})

	t.Run("All_on_nil", func(t *testing.T) {
		var a slices2.Set[string]
		expected := []string(nil)
		if actual := slices2.Sorted(a.All()); !equal(actual, expected) {
			t.Errorf("All()=`%#v`, expected `%#v`", actual, expected)
		}
	})

	t.Run("All_one", func(t *testing.T) {
		a := slices2.NewSet("foo")
		expected := []string{"foo"}
		if actual := a.All(); !equal(actual, expected) {
			t.Errorf("All()=`%#v`, expected `%#v`", actual, expected)
		}
	})

	t.Run("All_int", func(t *testing.T) {
		a := slices2.NewSet(3, 1, 2)
		expected := []int{1, 2, 3}
		if actual := slices2.Sorted(a.All()); !equal(actual, expected) {
			t.Errorf("All()=`%#v`, expected `%#v`", actual, expected)
		}
	})
}
