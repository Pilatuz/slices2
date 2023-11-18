package slices2_test

import (
	"fmt"
	"sync"
	"testing"

	"github.com/Pilatuz/slices2"
)

// ExampleFromChannel an example for `FromChannel` function.
func ExampleFromChannel() {
	wg := &sync.WaitGroup{}
	errCh := make(chan error)

	wg.Add(3)
	go func() {
		defer wg.Done()
		errCh <- nil
	}()
	go func() {
		defer wg.Done()
		errCh <- nil
	}()
	go func() {
		defer wg.Done()
		errCh <- nil
	}()

	go func() {
		wg.Wait()
		close(errCh)
	}()

	res := slices2.FromChannel(errCh)
	fmt.Println(res)
	// Output:
	// [<nil> <nil> <nil>]
}

// TestFromChannel unit tests for `FromChannel` function.
func TestFromChannel(tt *testing.T) {
	// string
	tt.Run("str", func(t *testing.T) {
		// FromChannel(nil) gives nil
		if e, a := ([]string)(nil), slices2.FromChannel[string](nil); !equal(a, e) {
			t.Errorf("expected `%#v`, found `%#v`", e, a)
		}

		// FromChan
		ch := make(chan string)
		go func() {
			defer close(ch)
			ch <- "foo"
			ch <- "bar"
		}()
		if e, a := []string{"foo", "bar"}, slices2.FromChannel(ch); !equal(a, e) {
			t.Errorf("expected `%#v`, found `%#v`", e, a)
		}
	})

	// integer
	tt.Run("int", func(t *testing.T) {
		// FromChannel(nil) gives nil
		if e, a := ([]int)(nil), slices2.FromChannel[int](nil); !equal(a, e) {
			t.Errorf("expected `%#v`, found `%#v`", e, a)
		}

		// FromChannel
		ch := make(chan int)
		go func() {
			defer close(ch)
			ch <- 123
			ch <- 456
		}()
		if e, a := []int{123, 456}, slices2.FromChannel(ch); !equal(a, e) {
			t.Errorf("expected `%#v`, found `%#v`", e, a)
		}
	})

	// boolean
	tt.Run("bool", func(t *testing.T) {
		// FromChannel(nil) gives nil
		if e, a := ([]bool)(nil), slices2.FromChannel[bool](nil); !equal(a, e) {
			t.Errorf("expected `%#v`, found `%#v`", e, a)
		}

		// FromChannel
		ch := make(chan bool)
		go func() {
			defer close(ch)
			ch <- true
			ch <- false
		}()
		if e, a := []bool{true, false}, slices2.FromChannel(ch); !equal(a, e) {
			t.Errorf("expected `%#v`, found `%#v`", e, a)
		}
	})
}

// TestToChannel unit tests for `ToChannel` function.
func TestToChannel(tt *testing.T) {
	// string
	tt.Run("str", func(t *testing.T) {
		// ToChannel(nil) gives not nil, but empty channel
		var ss []string
		ch := slices2.ToChannel(ss, 0)
		if ch == nil {
			t.Errorf("expected not-nil channel")
		}
		v, ok := <-ch
		if ok || v != "" {
			t.Errorf("expected empty channel, found `%#v`", v)
		}

		ss = []string{"foo", "bar"}
		if e, a := ss, slices2.FromChannel(slices2.ToChannel(ss, 0)); !equal(a, e) {
			t.Errorf("expected `%#v`, found `%#v`", e, a)
		}
	})

	// integer
	tt.Run("int", func(t *testing.T) {
		// ToChannel(nil) gives not nil, but empty channel
		var ss []int
		ch := slices2.ToChannel(ss, 0)
		if ch == nil {
			t.Errorf("expected not-nil channel")
		}
		v, ok := <-ch
		if ok || v != 0 {
			t.Errorf("expected empty channel, found `%#v`", v)
		}

		ss = []int{123, 456}
		if e, a := ss, slices2.FromChannel(slices2.ToChannel(ss, 0)); !equal(a, e) {
			t.Errorf("expected `%#v`, found `%#v`", e, a)
		}
	})

	// boolean
	tt.Run("bool", func(t *testing.T) {
		// ToChannel(nil) gives not nil, but empty channel
		var ss []bool
		ch := slices2.ToChannel(ss, 0)
		if ch == nil {
			t.Errorf("expected not-nil channel")
		}
		v, ok := <-ch
		if ok || v != false {
			t.Errorf("expected empty channel, found `%#v`", v)
		}

		ss = []bool{true, false}
		if e, a := ss, slices2.FromChannel(slices2.ToChannel(ss, 0)); !equal(a, e) {
			t.Errorf("expected `%#v`, found `%#v`", e, a)
		}
	})
}
