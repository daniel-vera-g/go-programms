package main

import (
	"testing"
)

func TestHello(t *testing.T) {

	assertCorrectMessage := func(t *testing.T, got, want string) {

		t.Helper()

		if got != want {
			t.Errorf("Got %q want %q", got, want)
		}
	}

	t.Run("Saying Hello to people", func(t *testing.T) {

		got := Hello("Daniel")
		want := "Hello Daniel"

		assertCorrectMessage(t, got, want)

	})

	t.Run("Empty string defaults to 'World'", func(t *testing.T) {
		got := Hello("")
		want := "Hello World"

		assertCorrectMessage(t, got, want)
	})
}
