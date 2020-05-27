package structs

import "testing"

func TestPerimeter(t *testing.T) {
	// Slice of anonymous structs, perimeterTests
	perimeterTests := []struct {
		shape Shape
		want  float64
	}{
		{Rectangle{10, 10}, 40.0},
		{Circle{5}, 31.415926535897932},
	}

	for _, tt := range perimeterTests {
		got := tt.shape.Perimeter()
		if got != tt.want {
			t.Errorf("got %g want %g", got, tt.want)
		}
	}
}

func TestArea(t *testing.T) {
	// Slice of anonymous structs, areaTests
	areaTests := []struct {
		shape Shape
		want  float64
	}{
		{Rectangle{12, 6}, 72.0},
		{Circle{10}, 314.1592653589793},
	}

	for _, tt := range areaTests {
		got := tt.shape.Area()
		if got != tt.want {
			t.Errorf("got %g want %g", got, tt.want)
		}
	}
}
