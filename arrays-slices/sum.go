package arraysslices

// Sum adds all the numbers in the given array
func Sum(numbers [5]int) (sum int) {
	for _, number := range numbers {
		sum += number
	}
	return
}
