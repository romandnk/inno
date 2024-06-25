package main

import "testing"

func BenchmarkFastLogger_Info(b *testing.B) {
	logger := FastLogger{}
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		logger.Info("test", "any value")
	}
}

func BenchmarkSlowLogger_Info(b *testing.B) {
	logger := SlowLogger{}
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		logger.Info("test", "any value")
	}
}
