package iteration

import (
	"fmt"
	"testing"
)

func TestRepeat(t *testing.T) {

	assertRepeaterCorrect := func(t *testing.T, text string, times int, expected string) {
		t.Helper()
		repeated := Repeat(text, times)
		if repeated != expected {
			t.Errorf("expected %q but got %q", expected, repeated)
		}
	}

	t.Run("repeat 10 times", func(t *testing.T) {
		assertRepeaterCorrect(t, "a", 10, "aaaaaaaaaa")
	})

	t.Run("repeat 4 times", func(t *testing.T) {
		assertRepeaterCorrect(t, "a", 4, "aaaa")
	})

	t.Run("repeat 0 times", func(t *testing.T) {
		assertRepeaterCorrect(t, "a", 0, "")
	})
}

func BenchmarkRepeat(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Repeat("a", 10)
	}
}

func ExampleRepeat() {
	fmt.Println(Repeat("foo", 2))
	// Output: foofoo
}
