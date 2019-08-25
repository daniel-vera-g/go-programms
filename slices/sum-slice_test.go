package slices

import (
	"reflect"
	"testing"
)

func TestSum(t *testing.T) {

	t.Run("Collection of any size", func(t *testing.T) {

		numbers := []int{1, 2, 3, 4}

		got := Sum(numbers)
		want := 10

		if got != want {
			t.Errorf("Got %d, want %d. %v", got, want, numbers)
		}
	})
}

func TestSumAll(t *testing.T) {

	numbers := []int{1, 2, 3, 4}
	numbers2 := []int{1, 2, 3}

	got := SumAll(numbers, numbers2)
	want := []int{10, 6}

	if !reflect.DeepEqual(got, want) {
		t.Errorf("Got %v, want %v", got, want)
	}
}
