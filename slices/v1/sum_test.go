package main

import "testing"

func TestSum(t *testing.T){

	t.Run("collection of 5 numbers", func(t *testing.T) {
		// A slice does not encode the size of a collection
		// Testing a slice of different sizes for redundancy
		numbers := []int{1, 2, 3, 4, 5,}

		got := Sum(numbers)
		want := 15

		if got != want {
			t.Errorf("got %d want %d given, %v", got, want, numbers)
		}
	})

	t.Run("Collection of any size", func(t *testing.T) {
		numbers := []int{1, 2, 3}

		got := Sum(numbers)
		want := 6

		if got != want {
			t.Errorf("got %d want %d given, %v", got, want, numbers)
		}
	})
}