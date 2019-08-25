package structs

// Perimeter takes the float64 width & length of a reactangle
// and returns it's float64 perimeter
func Perimeter(width float64, length float64) float64 {
	return 2 * (width * length)
}

// Area takes the float64 width & length of a reactangle
// and returns it's float64 are
func Area(width float64, length float64) float64 {
	return width * length
}
