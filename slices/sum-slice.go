package slices

// Sum takes a slice of integers and returns it's sum
func Sum(numbers []int) (sum int) {
	for _, number := range numbers {
		sum += number
	}

	return
}

// SumAll takes an multiple slices of integers
// and returns a slice with the sum of the numbers in each slice
func SumAll(numbers ...[]int) []int {
	lengthOfNumbers := len(numbers)
	sumOfArrays := make([]int, lengthOfNumbers)

	for i, array := range numbers {
		sumOfArrays[i] = Sum(array)
	}

	return sumOfArrays
}
