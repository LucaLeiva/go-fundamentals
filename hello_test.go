package main

import "testing"

func TestHello(t *testing.T) {
	got := Hello("Luca")
	want := "Hola Luca"

	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}
