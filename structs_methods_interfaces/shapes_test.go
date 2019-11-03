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

	areaTests := []struct {
		shape Shape
		want  float64
	}{
		{Rectangle{12.0, 6.0}, 72.0},
		{Circle{10}, 314.1592653589793},
	}

	for _, testData := range areaTests {
		got := testData.shape.Area()
		if got != testData.want {
			t.Errorf("got %g want %g", got, testData.want)
		}
	}
}
