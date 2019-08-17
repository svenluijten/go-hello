package arraysslices

import (
	"fmt"
	"testing"
)

func TestSum(t *testing.T) {
	numbers := [5]int{1, 2, 3, 4, 5}

	got := Sum(numbers)
	want := 15

	if want != got {
		t.Errorf("got %d, want %d. Given %v", got, want, numbers)
	}
}

func ExampleSum() {
	result := Sum([5]int{1, 2, 3, 4, 5})
	fmt.Println(result)
	// Output: 15
}
