package hello

import "testing"

func TestHelloLanguage(t *testing.T) {
	t.Run("spanish greeting", func(t *testing.T) {
		got := Hello("Rethik", "Spanish")
		want := "Hola Rethik"
		AssertCorrectMessage(t, got, want)
	})
	t.Run("english greeting", func(t *testing.T) {
		got := Hello("Rethik", "English")
		want := "Hello Rethik"
		AssertCorrectMessage(t, got, want)
	})
	t.Run("french greeting", func(t *testing.T) {
		got := Hello("Rethik", "French")
		want := "Bonjour Rethik"
		AssertCorrectMessage(t, got, want)
	})
	t.Run("no greeting (must default to english)", func(t *testing.T) {
		got := Hello("Rethik", "")
		want := "Hello Rethik"
		AssertCorrectMessage(t, got, want)
	})
}
