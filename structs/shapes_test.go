package structs

import "testing"

func TestPerimeter(t *testing.T) {

	rectangle := Rectangle{10.0, 20.0}

	got := Perimeter(rectangle)
	want := 400.0

	if got != want {
		t.Errorf("Got %.2f, want %.2f", got, want)
	}
}

func TestArea(t *testing.T) {

	areTests := []struct {
		shape Shape
		want  float64
	}{
		{Rectangle{10.0, 20.0}, 200.0},
		{Circle{10}, 314.1592653589793},
	}

	for _, tt := range areTests {
		got := tt.shape.Area()

		if got != tt.want {
			t.Errorf("Got %.2f, wanr %.2f", got, tt.want)
		}
	}
}
