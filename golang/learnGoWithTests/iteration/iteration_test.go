package iteration

import (
	"fmt"
	"testing"
)

func ExampleRepeat() {
	repeated := Repeat("k", 6)
	fmt.Println(repeated)
	// Output: kkkkkk
}

func TestRepeat(t *testing.T) {
	t.Run("Repeat string many times", func(t *testing.T) {
		repeated := Repeat("a", 5)
		expected := "aaaaa"

		if repeated != expected {
			t.Errorf("expected: %q, got: %q", expected, repeated)
		}
	})

	t.Run("Repeat string once", func(t *testing.T) {
		repeated := Repeat("a", 1)
		expected := "a"

		if repeated != expected {
			t.Errorf("expected: %q, got: %q", expected, repeated)
		}
	})
}

func BenchmarkRepeat(b *testing.B) {
	for b.Loop() {
		Repeat("a", 5)
	}
}
