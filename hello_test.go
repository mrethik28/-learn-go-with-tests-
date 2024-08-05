package main

import (
	"testing"
)

func TestHello(t *testing.T) {
	t.Run("say hello to people using name", func(t *testing.T) {
		got := Hello("Rethik")
		want := "Hello Rethik"
		assertCorrectMessage(t, got, want)
	})

	t.Run("Saying hello to world if name is not given", func(t *testing.T) {
		got := Hello("")
		want := "Hello Worldd"
		assertCorrectMessage(t, got, want)

	})

}

func assertCorrectMessage(t testing.TB, got, want string) {
	t.Helper()
	if got != want {
		t.Errorf("Error: wanted %q but got %q", want, got)

	}

}
