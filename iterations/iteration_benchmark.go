package iterations

import "testing"

func BenchmarkIterations(b testing.B) {
	var result string
	for i := 0; i < b.N; i++ {
		result = Repeat("a")
	}
	_ = result
}
