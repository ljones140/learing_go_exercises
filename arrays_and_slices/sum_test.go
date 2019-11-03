package main

import "reflect"
import "testing"

func TestSum(t *testing.T) {
	t.Run("collection of and size", func(t *testing.T) {
		numbers := []int{1, 2, 3, 4}

		got := Sum(numbers)
		want := 10

		if got != want {
			t.Errorf("got %d want %d given, %v", got, want, numbers)
		}
	})
}

func TestSumAllTails(t *testing.T) {

	assertSlicesMatch := func(t *testing.T, got, want []int) {
		t.Helper()
		if !reflect.DeepEqual(got, want) {
			t.Errorf("got %d want %d", got, want)
		}
	}

	t.Run("make the sum of some slices", func(t *testing.T) {
		got := SumAllTails([]int{1, 2, 1}, []int{0, 9})
		want := []int{3, 9}
		assertSlicesMatch(t, got, want)
	})

	t.Run("defaults a slice total to zero if given an empty slice", func(t *testing.T) {
		got := SumAllTails([]int{}, []int{0, 9})
		want := []int{0, 9}
		assertSlicesMatch(t, got, want)
	})
}
