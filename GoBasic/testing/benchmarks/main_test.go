package main

import (
	"testing"
)

func BenchmarkSum(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Sum(4, 1, 2, -5, 9)
	}
}
