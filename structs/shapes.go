package structs

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
func Area(rectangle Rectangle) float64 {
	return rectangle.Width * rectangle.Length
}
