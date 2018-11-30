package main

import "testing"

func TestHello(t *testing.T) {

	assertCorrectMessage := func(t *testing.T, got, want string) {
		t.Helper()
		if got != want {
			t.Errorf("Got '%s', Want '%s'", got, want)
		}
	}

	t.Run("Saying hello to people", func(t *testing.T) {
		got := Hello("Hitesh", "")
		want := "Hello, Hitesh"
		assertCorrectMessage(t, got, want)
	})

	t.Run("Saying hello world when no values are passed", func(t *testing.T) {
		got := Hello("", "")
		want := "Hello, World"
		assertCorrectMessage(t, got, want)
	})

	t.Run("Saying Hello in French", func(t *testing.T) {
		got := Hello("Hitesh", "French")
		want := "Bonjour Hitesh"
		assertCorrectMessage(t, got, want)
	})

	t.Run("Saying Hello in Spanish", func(t *testing.T) {
		got := Hello("Hitesh", "Spanish")
		want := "Hola Hitesh"
		assertCorrectMessage(t, got, want)
	})

}
