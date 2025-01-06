package main

import (
	"testing"
	"reflect"
)

func TestSumAll(t *testing.T){

	checkSums := func(t testing.TB, got, want []int) {
		t.Helper()
		// Go does not let you use equality operators with slices
		// use !reflect.DeepEqual to chck if two variables are the same
		// rather than iterating through each got and want slice and check their values
		// Please note reflect.DeepEqual is not typesafe
		if !reflect.DeepEqual(got, want) {
			t.Errorf("got %v want %v", got, want)
		}
	}

	t.Run("make the sums of some slices", func(t *testing.T){
		// The tail is all elements except the first one
		got := SumAllTrails([]int{1, 2}, []int{0, 9})
		want := []int{2, 9}
		checkSums(t, got, want)
	})

	t.Run("safely sum empty slices", func(t *testing.T){
		
		got := SumAllTrails([]int{}, []int{0, 9, 1})
		want := []int{0, 10}
		checkSums(t, got, want)
	})
}