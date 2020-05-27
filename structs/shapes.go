package structs

import "math"

// A Shape is implemented by anything with perimeter or area.
type Shape interface {
	Perimeter() float64
	Area() float64
}

// A Rectangle represents a rectangle shape with a width and height.
type Rectangle struct {
	Width  float64
	Height float64
}

// Perimeter returns the perimeter of a Rectangle.
func (r Rectangle) Perimeter() float64 {
	return 2 * (r.Width + r.Height)
}

// Area returns the area of a Rectangle.
func (r Rectangle) Area() float64 {
	return r.Width * r.Height
}

// A Circle represents a circle shape with a radius.
type Circle struct {
	Radius float64
}

// Perimeter returns the perimeter of a Circle.
func (c Circle) Perimeter() float64 {
	return 2 * math.Pi * c.Radius
}

// Area returns the area of a Circle.
func (c Circle) Area() float64 {
	return math.Pi * c.Radius * c.Radius
}
