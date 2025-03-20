package main

import (
	"reflect"
	"testing"
)

func TestSum(t *testing.T) {
	nums := []int{1, 2, 3, 4, 5}
	got := Sum(nums)
	want := 15
	if got != want {
		t.Errorf("want: %d, got: %d, given: %v", want, got, nums)
	}
}

func TestSumAll(t *testing.T) {
	result := SumAll([]int{1, 2, 3, 4}, []int{5, 6, 7})
	expected := []int{10, 18}
	if !reflect.DeepEqual(result, expected) {
		t.Errorf("expected: %v, got: %v", expected, result)
	}
}
