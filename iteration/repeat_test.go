package iteration

import (
	"fmt"
	"testing"
)

func TestRepeat(t *testing.T) {
	assertCorrectOutput := func(t *testing.T, expected, actual string) {
		if expected != actual {
			t.Errorf("expected %q but got %q", expected, actual)
		}
	}

	t.Run("repeating a string 5 times", func(t *testing.T) {
		repeated := Repeat("a", 5)
		assertCorrectOutput(t, "aaaaa", repeated)
	})

	t.Run("repeating a string 10 times", func(t *testing.T) {
		repeated := Repeat("b", 10)
		assertCorrectOutput(t, "bbbbbbbbbb", repeated)
	})
}

func ExampleRepeat() {
	repeated := Repeat("a", 5)
	fmt.Println(repeated)
	// Output: aaaaa
}

func BenchmarkRepeat(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Repeat("a", 5)
	}
}
