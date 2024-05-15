package main

import "testing"

func TestSum(t *testing.T) {

	type test struct {
		data   []int
		answer int
	}

	tests := []test{
		{data: []int{3, 3, 4}, answer: 10},
		{data: []int{4, 5, 6}, answer: 15},
		{data: []int{-1, 0, 6, 3}, answer: 8},
		{data: []int{-2, -2}, answer: -4},
	}

	for _, t1 := range tests {
		calculatedAns := Sum(t1.data...)

		if calculatedAns != t1.answer {
			t.Error("Expected ", t1.answer, " Got ", calculatedAns)
		}
	}
}
