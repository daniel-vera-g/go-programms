package structs

import "testing"

func TestPerimeter(t *testing.T) {

	got := Perimeter(10.0, 20.0)
	want := 400.0

	if got != want {
		t.Errorf("Got %.2f, want %.2f", got, want)
	}
}

func TestArea(t *testing.T) {

	got := Area(10.0, 20.0)
	want := 200.0

	if got != want {
		t.Errorf("Got %.2f, want %.2f", got, want)
	}
}
