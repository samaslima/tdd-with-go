package main

import "testing"

func TestHello(t *testing.T) {
	got := Hello("Samara")
	want := "hello Samara"

	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}
