package main

import (
	"testing"
	"reflect"
)

func TestSumAll(t *testing.T){
	// The tail is all elements except the first one
	got := SumAllTrails([]int{1, 2}, []int{0, 9})
	want := []int{2, 9}

	// Go does not let you use equality operators with slices
	// use !reflect.DeepEqual to chck if two variables are the same
	// rather than iterating through each got and want slice and check their values
	// Please note reflect.DeepEqual is not typesafe
	if !reflect.DeepEqual(got, want) {
		t.Errorf("got %v want %v", got, want)
	}
}