package main

// Rectangle has the dimensions of a rectangle.
type Rectangle struct {
	width float64
	height float64
}

// Area returns the area of a rectangle.
func Area (rectangle Rectangle) float64{
	return (rectangle.width * rectangle.height)
}

// Perimeter returns the perimeter of a rectangle.
func Perimeter(rectangle Rectangle) float64{
	return (rectangle.width + rectangle.height) * 2
}