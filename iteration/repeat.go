package iteration

// Repeat repeats the string value passed in
func Repeat(value string, times int) (repeated string) {
	for i := 0; i < times; i++ {
		repeated += value
	}
	return
}
