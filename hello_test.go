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

	// General Hello
	t.Run("Saying Hello to people", func(t *testing.T) {

		got := Hello("Daniel", "")
		want := "Hello Daniel"

		assertCorrectMessage(t, got, want)

	})

	// Empty defaults
	t.Run("Empty string defaults to 'World'", func(t *testing.T) {
		got := Hello("", "")
		want := "Hello World"

		assertCorrectMessage(t, got, want)
	})

	// Spanish
	t.Run("in Spanish", func(t *testing.T) {
		got := Hello("Lukas", "Spanish")
		want := "Hola Lukas"

		assertCorrectMessage(t, got, want)
	})

	// Deutsch
	t.Run("in German", func(t *testing.T) {
		got := Hello("Mark", "Deutsch")
		want := "Hallo Mark"

		assertCorrectMessage(t, got, want)
	})

}
