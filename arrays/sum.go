package arrays

// Sum returns the sum of the elements within a slice.
func Sum(numbers []int) int {
	sum := 0
	for _, number := range numbers {
		sum += number
	}
	return sum
}
