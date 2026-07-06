package main

import (
	"testing"
	"github.com/ymn2525/knapsack-solution/main"
)

func BenchmarkDP(b *testing.B) {
	for i := 0; i < b.N; i++ {
		DynamicProgramming()
	}
}

func BenchmarkBruteForce(b *testing.B) {
	for i := 0; i < b.N; i++ {
		BruteForce()
	}
}
