package structs

import "testing"

func TestPerimeter(t *testing.T) {

	rectangle := Rectangle{10.0, 20.0}

	got := Perimeter(rectangle)
	hasArea := 400.0

	if got != hasArea {
		t.Errorf("Got %.2f, hasArea %.2f", got, hasArea)
	}
}

func TestArea(t *testing.T) {

	areTests := []struct {
		shape   Shape
		hasArea float64
	}{
		{shape: Rectangle{Width: 10.0, Length: 20.0}, hasArea: 200.0},
		{shape: Circle{Radius: 10}, hasArea: 314.1592653589793},
		{shape: Triangle{Base: 12, Height: 6}, hasArea: 36.0},
	}

	for _, tt := range areTests {
		got := tt.shape.Area()

		if got != tt.hasArea {
			t.Errorf("Got %.2f, wanr %.2f", got, tt.hasArea)
		}
	}
}
