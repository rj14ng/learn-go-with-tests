package structs

// A Rectangle represents a rectangle shape with a width and height.
type Rectangle struct {
	Width  float64
	Height float64
}

// Perimeter returns the perimeter of a Rectangle.
func Perimeter(rectangle Rectangle) float64 {
	return 2 * (rectangle.Width + rectangle.Height)
}

// Area returns the area of a Rectangle.
func Area(rectangle Rectangle) float64 {
	return rectangle.Width * rectangle.Height
}
