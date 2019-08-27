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
		{shape: Rectangle{Width: 10.0, Length: 20.0}, want: 200.0},
		{shape: Circle{Radius: 10}, want: 314.1592653589793},
		{shape: Triangle{Base: 12, Height: 6}, want: 36.0},
	}

	for _, tt := range areTests {
		got := tt.shape.Area()

		if got != tt.want {
			t.Errorf("Got %.2f, wanr %.2f", got, tt.want)
		}
	}
}
