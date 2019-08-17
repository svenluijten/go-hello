package arraysslices

import (
	"fmt"
	"reflect"
	"testing"
)

func TestSum(t *testing.T) {
	t.Run("collection of numbers", func(t *testing.T) {
		numbers := []int{1, 2, 3, 4, 5}

		actual := Sum(numbers)
		expected := 15

		if expected != actual {
			t.Errorf("got %d, want %d. Given %v", actual, expected, numbers)
		}
	})
}

func TestSumAll(t *testing.T) {
	actual := SumAll([]int{1, 2}, []int{0, 9})
	expected := []int{3, 9}

	if !reflect.DeepEqual(actual, expected) {
		t.Errorf("got %v want %v", actual, expected)
	}
}

func TestSumAllTails(t *testing.T) {
	actual := SumAllTails([]int{1, 2}, []int{0, 9})
	expected := []int{2, 9}

	if !reflect.DeepEqual(actual, expected) {
		t.Errorf("got %v, expected %v", actual, expected)
	}
}

func ExampleSum() {
	result := Sum([]int{1, 2, 3, 4, 5})
	fmt.Println(result)
	// Output: 15
}

func ExampleSumAll() {
	result := SumAll([]int{1, 2, 3}, []int{4, 5, 6})
	fmt.Println(result)
	// Output: [6 15]
}

func ExampleSumAllTails() {
	result := SumAllTails([]int{1, 2, 3}, []int{4, 5, 6})
	fmt.Println(result)
	// Output: [5 11]
}
