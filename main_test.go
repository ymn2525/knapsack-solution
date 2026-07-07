package main

import (
	"testing"
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
