package main

import (
	"testing"
)

func Benchmark_Part3_Normal(b *testing.B) {
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		part3()
	}

	b.StopTimer()

	nsPerOp := float64(b.Elapsed().Nanoseconds()) / float64(b.N)

	b.ReportMetric(nsPerOp/1e6, "ms/op")
}

func Benchmark_Part3_Para(b *testing.B) {
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		part3Para()
	}

	b.StopTimer()

	nsPerOp := float64(b.Elapsed().Nanoseconds()) / float64(b.N)

	b.ReportMetric(nsPerOp/1e6, "ms/op")
}
