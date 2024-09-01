package integers

import "testing"

func AssertCorrectMessage(t testing.TB, got, want int) {
	t.Helper()
	if got != want {
		t.Errorf("Error: wanted %q but got %q", want, got)

	}

}
