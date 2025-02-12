package main

import "testing"

func TestHello(t *testing.T) {
	t.Run("saying hello to people", func(t *testing.T) {
		got := Hello("samara", "english")
		want := "hello samara"
		assertCorrectMessage(t, got, want)
	})
	t.Run("say 'hello world' when an empty string is supplied", func(t *testing.T) {
		got := Hello("", "english")
		want := "hello world"
		assertCorrectMessage(t, got, want)
	})
	t.Run("in spanish", func(t *testing.T) {
		got := Hello("elodie", "spanish")
		want := "hola elodie"
		assertCorrectMessage(t, got, want)
	})
	t.Run("in french", func(t *testing.T) {
		got := Hello("mary", "french")
		want := "bonjour mary"
		assertCorrectMessage(t, got, want)
	})
}

func assertCorrectMessage(t testing.TB, got, want string) {
	t.Helper()
	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}
