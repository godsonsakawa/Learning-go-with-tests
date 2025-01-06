package main

func Sum(arr []int) int{
	sum := 0
	// using the blank identifier to ignore the index value
	// range iterates over an array
	// on each iteration range returns two values the index and value
	for _, number := range arr {
		sum += number
	}	
	return sum
}