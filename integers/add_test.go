package integers

import (
	"fmt"
	"testing"
)

func TestIntegers(t *testing.T) {
	t.Run("check addition", func(t *testing.T) {
		got := Add(3, 4)
		want := 7
		AssertCorrectMessage(t, got, want)

	})
}

func ExampleAdd() {
	sum := Add(1, 5)
	fmt.Println(sum)
	// Output: 6
}
