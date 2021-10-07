package main

import (
	"testing"
)

func TestMain(t *testing.T) {
	t.Run("default input 1", func(t *testing.T) {
		n := 3
		l := 2
		r := 2
		got := Solve(n, l, r)
		want := 2
		if got != want {
			t.Errorf("got %d want %d given %d %d %d", got, want, n, l, r)
		}
	})

	t.Run("default input 2", func(t *testing.T) {
		n := 4
		l := 1
		r := 2
		got := Solve(n, l, r)
		want := 2
		if got != want {
			t.Errorf("got %d want %d given %d %d %d", got, want, n, l, r)
		}
	})

	t.Run("input 1", func(t *testing.T) {
		n := 4
		l := 1
		r := 4
		got := Solve(n, l, r)
		want := 1
		if got != want {
			t.Errorf("got %d want %d given %d %d %d", got, want, n, l, r)
		}
	})

	t.Run("input 2", func(t *testing.T) {
		n := 5
		l := 1
		r := 5
		got := Solve(n, l, r)
		want := 1
		if got != want {
			t.Errorf("got %d want %d given %d %d %d", got, want, n, l, r)
		}
	})

	t.Run("input 3", func(t *testing.T) {
		n := 5
		l := 1
		r := 4
		got := Solve(n, l, r)
		want := 6
		// 1|5 1 4 3 2|4
		// 1|5 2 4 3 1|4
		// 1|5 3 4 2 1|4
		// 1|5 4 1 3 2|4
		// 1|5 4 2 3 1|4
		// 1|5 4 3 1 2|4
		if got != want {
			t.Errorf("got %d want %d given %d %d %d", got, want, n, l, r)
		}
	})

	t.Run("input 4", func(t *testing.T) {
		n := 5
		l := 2
		r := 3
		got := Solve(n, l, r)
		want := 18
		if got != want {
			t.Errorf("got %d want %d given %d %d %d", got, want, n, l, r)
		}
	})

	t.Run("input 5", func(t *testing.T) {
		n := 1
		l := 1
		r := 1
		got := Solve(n, l, r)
		want := 1
		if got != want {
			t.Errorf("got %d want %d given %d %d %d", got, want, n, l, r)
		}
	})

	t.Run("input 6", func(t *testing.T) {
		n := 2
		l := 1
		r := 2
		got := Solve(n, l, r)
		want := 1
		if got != want {
			t.Errorf("got %d want %d given %d %d %d", got, want, n, l, r)
		}
	})
}
