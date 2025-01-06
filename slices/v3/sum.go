package main

// Sum calculates the total from a slice of numbers.
func Sum(numbers []int) int {
	sum := 0
	for _, number := range numbers {
		sum += number
	}
	return sum
}

// Writing a variadic function in Go
// SumAll calculates the respective sums of every slice passed in.
func SumAll(numbersToSum ...[]int) []int {

	// initialize an empty slice
	var sums []int
	for _, numbers := range numbersToSum {
		// append takes a slice and a new value and 
		// returns a slice with all items in it as we work through the varargs
		sums = append(sums, Sum(numbers))
	}
	return sums
}