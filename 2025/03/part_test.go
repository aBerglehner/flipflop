package main

import (
	"testing"
)

func Benchmark_Para_Ch_Part2(b *testing.B) {
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		part2ParaCh()
	}

	b.StopTimer()

	nsPerOp := float64(b.Elapsed().Nanoseconds()) / float64(b.N)

	b.ReportMetric(nsPerOp/1e6, "ms/op")
}

func Benchmark_Para_Mutex_Part2(b *testing.B) {
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		part2ParaMutex()
	}

	b.StopTimer()

	nsPerOp := float64(b.Elapsed().Nanoseconds()) / float64(b.N)

	b.ReportMetric(nsPerOp/1e6, "ms/op")
}

func Benchmark_Normal_Part2(b *testing.B) {
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		part2()
	}

	b.StopTimer()

	nsPerOp := float64(b.Elapsed().Nanoseconds()) / float64(b.N)

	b.ReportMetric(nsPerOp/1e6, "ms/op")
}
