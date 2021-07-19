package main

import (
	"testing"
)

func TestMain(t *testing.T) {
	t.Run("1", func(t *testing.T) {
		arr := [][]int{
			{1, 2, 3, 4, 5},
			{5, 4, 3, 2, 1},
			{2, 3, 4, 5, 6},
			{6, 5, 4, 3, 2},
			{1, 2, 1, 2, 1},
		}
		got := Solve(arr, 5, 5)
		want := 19

		if got != want {
			t.Errorf("got %d want %d\ngiven, %v", got, want, arr)
		}
	})

	t.Run("2", func(t *testing.T) {
		arr := [][]int{
			{1, 2, 3, 4, 5},
			{1, 2, 3, 4, 5},
			{1, 2, 3, 4, 5},
			{1, 2, 3, 4, 5},
			{1, 2, 3, 4, 5},
		}
		got := Solve(arr, 5, 5)
		want := 20

		if got != want {
			t.Errorf("got %d want %d\ngiven, %v", got, want, arr)
		}
	})

	t.Run("3", func(t *testing.T) {
		arr := [][]int{
			{1, 2, 1, 2, 1, 2, 1, 2, 1, 2},
			{2, 1, 2, 1, 2, 1, 2, 1, 2, 1},
			{1, 2, 1, 2, 1, 2, 1, 2, 1, 2},
			{2, 1, 2, 1, 2, 1, 2, 1, 2, 1},
		}
		got := Solve(arr, 4, 10)
		want := 7

		if got != want {
			t.Errorf("got %d want %d\ngiven, %v", got, want, arr)
		}
	})

	t.Run("4", func(t *testing.T) {
		arr := [][]int{
			{0, 0, 0, 0},
			{0, 0, 0, 0},
			{0, 0, 0, 0},
			{1, 2, 3, 4},
		}
		got := Solve(arr, 4, 4)
		want := 10

		if got != want {
			t.Errorf("got %d want %d\ngiven, %v", got, want, arr)
		}
	})

	t.Run("5", func(t *testing.T) {
		arr := [][]int{
			{0, 0, 0, 1},
			{0, 0, 0, 2},
			{0, 0, 0, 3},
			{0, 0, 0, 4},
		}
		got := Solve(arr, 4, 4)
		want := 10

		if got != want {
			t.Errorf("got %d want %d\ngiven, %v", got, want, arr)
		}
	})

	t.Run("6", func(t *testing.T) {
		arr := [][]int{
			{0, 0, 0, 0},
			{0, 0, 0, 0},
			{0, 0, 1, 2},
			{0, 0, 3, 4},
		}
		got := Solve(arr, 4, 4)
		want := 10

		if got != want {
			t.Errorf("got %d want %d\ngiven, %v", got, want, arr)
		}
	})

	t.Run("7", func(t *testing.T) {
		arr := [][]int{
			{0, 0, 0, 0},
			{0, 0, 1, 0},
			{0, 0, 2, 0},
			{0, 0, 3, 4},
		}
		got := Solve(arr, 4, 4)
		want := 10

		if got != want {
			t.Errorf("got %d want %d\ngiven, %v", got, want, arr)
		}
	})

	t.Run("8", func(t *testing.T) {
		arr := [][]int{
			{0, 0, 0, 0},
			{0, 0, 0, 0},
			{0, 1, 2, 3},
			{0, 4, 0, 0},
		}
		got := Solve(arr, 4, 4)
		want := 10

		if got != want {
			t.Errorf("got %d want %d\ngiven, %v", got, want, arr)
		}
	})

	t.Run("9", func(t *testing.T) {
		arr := [][]int{
			{0, 0, 0, 0},
			{0, 0, 1, 2},
			{0, 0, 0, 3},
			{0, 0, 0, 4},
		}
		got := Solve(arr, 4, 4)
		want := 10

		if got != want {
			t.Errorf("got %d want %d\ngiven, %v", got, want, arr)
		}
	})

	t.Run("10", func(t *testing.T) {
		arr := [][]int{
			{0, 0, 0, 0},
			{0, 0, 0, 0},
			{0, 0, 0, 1},
			{0, 4, 3, 2},
		}
		got := Solve(arr, 4, 4)
		want := 10

		if got != want {
			t.Errorf("got %d want %d\ngiven, %v", got, want, arr)
		}
	})

	t.Run("11", func(t *testing.T) {
		arr := [][]int{
			{0, 0, 0, 0},
			{0, 0, 0, 1},
			{0, 0, 0, 2},
			{0, 0, 4, 3},
		}
		got := Solve(arr, 4, 4)
		want := 10

		if got != want {
			t.Errorf("got %d want %d\ngiven, %v", got, want, arr)
		}
	})

	t.Run("12", func(t *testing.T) {
		arr := [][]int{
			{0, 0, 0, 0},
			{0, 0, 0, 0},
			{0, 1, 0, 0},
			{0, 2, 3, 4},
		}
		got := Solve(arr, 4, 4)
		want := 10

		if got != want {
			t.Errorf("got %d want %d\ngiven, %v", got, want, arr)
		}
	})

	t.Run("13", func(t *testing.T) {
		arr := [][]int{
			{0, 0, 0, 0},
			{0, 0, 2, 1},
			{0, 0, 3, 0},
			{0, 0, 4, 0},
		}
		got := Solve(arr, 4, 4)
		want := 10

		if got != want {
			t.Errorf("got %d want %d\ngiven, %v", got, want, arr)
		}
	})

	t.Run("14", func(t *testing.T) {
		arr := [][]int{
			{0, 0, 0, 0},
			{0, 0, 0, 0},
			{0, 1, 2, 3},
			{0, 0, 0, 4},
		}
		got := Solve(arr, 4, 4)
		want := 10

		if got != want {
			t.Errorf("got %d want %d\ngiven, %v", got, want, arr)
		}
	})

	t.Run("15", func(t *testing.T) {
		arr := [][]int{
			{0, 0, 0, 0},
			{0, 0, 0, 1},
			{0, 0, 2, 3},
			{0, 0, 4, 0},
		}
		got := Solve(arr, 4, 4)
		want := 10

		if got != want {
			t.Errorf("got %d want %d\ngiven, %v", got, want, arr)
		}
	})

	t.Run("16", func(t *testing.T) {
		arr := [][]int{
			{0, 0, 0, 0},
			{0, 0, 1, 0},
			{0, 0, 2, 3},
			{0, 0, 0, 4},
		}
		got := Solve(arr, 4, 4)
		want := 10

		if got != want {
			t.Errorf("got %d want %d\ngiven, %v", got, want, arr)
		}
	})

	t.Run("17", func(t *testing.T) {
		arr := [][]int{
			{0, 0, 0, 0},
			{0, 0, 0, 0},
			{0, 1, 2, 0},
			{0, 0, 3, 4},
		}
		got := Solve(arr, 4, 4)
		want := 10

		if got != want {
			t.Errorf("got %d want %d\ngiven, %v", got, want, arr)
		}
	})

	t.Run("18", func(t *testing.T) {
		arr := [][]int{
			{0, 0, 0, 0},
			{0, 0, 0, 0},
			{0, 0, 3, 4},
			{0, 1, 2, 0},
		}
		got := Solve(arr, 4, 4)
		want := 10

		if got != want {
			t.Errorf("got %d want %d\ngiven, %v", got, want, arr)
		}
	})

	t.Run("19", func(t *testing.T) {
		arr := [][]int{
			{0, 0, 0, 0},
			{0, 0, 0, 1},
			{0, 0, 2, 3},
			{0, 0, 0, 4},
		}
		got := Solve(arr, 4, 4)
		want := 10

		if got != want {
			t.Errorf("got %d want %d\ngiven, %v", got, want, arr)
		}
	})

	t.Run("20", func(t *testing.T) {
		arr := [][]int{
			{0, 0, 0, 0},
			{0, 0, 1, 0},
			{0, 0, 2, 3},
			{0, 0, 4, 0},
		}
		got := Solve(arr, 4, 4)
		want := 10

		if got != want {
			t.Errorf("got %d want %d\ngiven, %v", got, want, arr)
		}
	})

	t.Run("21", func(t *testing.T) {
		arr := [][]int{
			{0, 0, 0, 0},
			{0, 0, 0, 0},
			{0, 0, 1, 0},
			{0, 2, 3, 4},
		}
		got := Solve(arr, 4, 4)
		want := 10

		if got != want {
			t.Errorf("got %d want %d\ngiven, %v", got, want, arr)
		}
	})

	t.Run("22", func(t *testing.T) {
		arr := [][]int{
			{0, 0, 0, 0},
			{0, 0, 0, 0},
			{0, 1, 2, 3},
			{0, 0, 4, 0},
		}
		got := Solve(arr, 4, 4)
		want := 10

		if got != want {
			t.Errorf("got %d want %d\ngiven, %v", got, want, arr)
		}
	})
}
