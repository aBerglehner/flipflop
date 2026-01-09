package main

import (
	"testing"
)

func BenchmarkNormal(b *testing.B) {
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		part1()
	}

	b.StopTimer()

	nsPerOp := float64(b.Elapsed().Nanoseconds()) / float64(b.N)

	b.ReportMetric(nsPerOp/1e6, "ms/op")
}
