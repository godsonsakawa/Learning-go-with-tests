package iteration

import "testing"

func TestRepeat(t *testing.T) {
	got := Repeat("a")
	want := "aaaaa"

	if got != want {
		t.Errorf("Expected %q but got %q", want, got)
	}
}

// Using Benchmarks to measure performance of the repeat function.
func BenchmarkRepeat(b *testing.B) {
	// The loop runs b.N times set by the GO testing framework
	// to determine how many iterations to run to get a stable measurement.
	for i := 0; i < b.N; i++ {
		Repeat("a")
	}
}