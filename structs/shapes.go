package structs

import (
	"math"
)

// Rectangle with width and length
type Rectangle struct {
	Width  float64
	Length float64
}

// Perimeter takes the float64 width & length of a reactangle
// and returns it's float64 perimeter
func Perimeter(rectangle Rectangle) float64 {
	return 2 * (rectangle.Width * rectangle.Length)
}

// Area takes the float64 width & length of a reactangle
// and returns it's float64 are
func (rectangle Rectangle) Area() float64 {
	return rectangle.Width * rectangle.Length
}

// Circle with radius
type Circle struct {
	Radius float64
}

// Area takes the radius of a circle
// and returns it's area
func (c Circle) Area() float64 {

	return c.Radius * c.Radius * math.Pi

}
