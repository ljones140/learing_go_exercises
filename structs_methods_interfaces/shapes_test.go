package main

import (
	"testing"
)

func TestPerimeter(t *testing.T) {
	got := Perimeter(
		Rectangle{10.0, 10.0},
	)
	want := 40.0

	if got != want {
		t.Errorf("got %.2f want %.2f", got, want)
	}
}

func TestArea(t *testing.T) {

	checkArea := func(t *testing.T, shape Shape, want float64) {
		t.Helper()
		got := shape.Area()

		if got != want {
			t.Errorf("got %g want %g", got, want)
		}
	}
	t.Run("rectangles", func(t *testing.T) {
		checkArea(t, Rectangle{12.0, 6.0}, 72.0)
	})

	t.Run("cirlces", func(t *testing.T) {
		checkArea(t, Circle{10}, 314.1592653589793)
	})
}
