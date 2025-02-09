package main

import "testing"

func TestHello(t *testing.T) {
	t.Run("saying hello to people", func(t *testing.T) {
		got := Hello("Levi", "")
		want := "Hello, Levi!"

		assertCorrectMessage(t, got, want)
	})

	t.Run("say 'Hello, world!' when an empty string is supplied", func(t *testing.T) {
		got := Hello("", "")
		want := "Hello, world!"

		assertCorrectMessage(t, got, want)
	})

	t.Run("in Portuguese", func(t *testing.T) {
		got := Hello("Levi", "Portuguese")
		want := "Olá, Levi!"

		assertCorrectMessage(t, got, want)
	})

	t.Run("in Portuguese, with empty string", func(t *testing.T) {
		got := Hello("", "Portuguese")
		want := "Olá, mundo!"

		assertCorrectMessage(t, got, want)
	})

	t.Run("in French", func(t *testing.T) {
		got := Hello("Levi", "French")
		want := "Bonjour, Levi!"

		assertCorrectMessage(t, got, want)
	})

	t.Run("in French, with empty string", func(t *testing.T) {
		got := Hello("", "French")
		want := "Bonjour, monde!"

		assertCorrectMessage(t, got, want)
	})
}

func assertCorrectMessage(t testing.TB, got, want string) {
	t.Helper()
	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}
