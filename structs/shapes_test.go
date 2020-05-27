package structs

import "testing"

func TestPerimeter(t *testing.T) {
	// Slice of anonymous structs, perimeterTests
	perimeterTests := []struct {
		name         string
		shape        Shape
		hasPerimeter float64
	}{
		{name: "Rectangle", shape: Rectangle{10, 10}, hasPerimeter: 40.0},
		{name: "Circle", shape: Circle{5}, hasPerimeter: 31.415926535897932},
		{name: "Triangle", shape: Triangle{5, 12, 13}, hasPerimeter: 30.0},
	}

	for _, tt := range perimeterTests {
		// Using tt.name from the case to use it as the 't.Run' test name
		t.Run(tt.name, func(t *testing.T) {
			got := tt.shape.Perimeter()
			if got != tt.hasPerimeter {
				t.Errorf("%#v got %g want %g", tt.shape, got, tt.hasPerimeter)
			}
		})
	}
}

func TestArea(t *testing.T) {
	// Slice of anonymous structs, areaTests
	areaTests := []struct {
		name    string
		shape   Shape
		hasArea float64
	}{
		{name: "Rectangle", shape: Rectangle{12, 6}, hasArea: 72.0},
		{name: "Circle", shape: Circle{10}, hasArea: 314.1592653589793},
		{name: "Triangle", shape: Triangle{3, 4, 5}, hasArea: 6.0},
	}

	for _, tt := range areaTests {
		// Using tt.name from the case to use it as the 't.Run' test name
		t.Run(tt.name, func(t *testing.T) {
			got := tt.shape.Area()
			if got != tt.hasArea {
				t.Errorf("%#v got %g want %g", tt.shape, got, tt.hasArea)
			}
		})
	}
}
