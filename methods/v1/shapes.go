package main

import "math"

// Rectangle has the dimensions of a rectangle.
type Rectangle struct {
	width float64
	height float64
}

// Circle represents a circle
type Circle struct {
	radius float64
}

// Area returns the area of a rectangle.
func (r Rectangle) Area() float64{
	return (r.width * r.height)
}

// Perimeter returns the perimeter of a rectangle.
func Perimeter(rectangle Rectangle) float64{
	return (rectangle.width + rectangle.height) * 2
}

// Area returns the area of a cirle 
func (c Circle) Area() float64 {
	return math.Pi * c.radius * c.radius
}