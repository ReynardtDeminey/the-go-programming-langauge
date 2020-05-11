package main

import "testing"

// Benchmark the printArgs function
func BenchmarkPrintArgs(b *testing.B) {
	for i := 0; i < b.N; i++ {
		printArgs()
	}
}

// Benchmark the PrintArgsRange function
func BenchmarkPrintArgsRange(b *testing.B) {
	for i := 0; i < b.N; i++ {
		printArgsRange()
	}
}

// Benchmark the printArgsJoin function
func BenchmarkPrintArgsJoin(b *testing.B) {
	for i := 0; i < b.N; i++ {
		printArgsJoin()
	}
}
