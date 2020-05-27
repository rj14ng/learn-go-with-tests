package structs

import "math"

// A Shape is implemented by anything with perimeter or area.
type Shape interface {
	Perimeter() float64
	Area() float64
}

// A Rectangle represents a rectangle shape, defined with a width and height.
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

// A Circle represents a circle shape, defined with a radius.
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

// A Triangle represents a triangle shape, defined with its 3 sides.
type Triangle struct {
	SideA, SideB, SideC float64
}

// Perimeter returns the perimeter of a Triangle.
func (t Triangle) Perimeter() float64 {
	return t.SideA + t.SideB + t.SideC
}

// Area returns the area of a Triangle, using Heron's formula.
func (t Triangle) Area() float64 {
	p := t.Perimeter() / 2.0 // Semiperimeter
	return math.Sqrt(p * (p - t.SideA) * (p - t.SideB) * (p - t.SideC))
}
