package main

import "testing"

func TestSum(t *testing.T) {
	sumVal := Sum(2, 3)
	if sumVal != 5 {
		t.Error("Expected ", 5, " But got ", sumVal)
	}
}
