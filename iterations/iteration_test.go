package iterations

import "testing"

func TestIteration(t *testing.T) {
	t.Run("checking", func(t *testing.T) {
		got := Repeat("a")
		want := "aaaaa"
		AssertCorrectMessage(t, got, want)
	})
}

func AssertCorrectMessage(t testing.TB, got, want string) {
	t.Helper()
	if got != want {
		t.Errorf("Error: wanted %q but got %q", want, got)

	}

}
