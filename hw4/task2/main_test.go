package main

import "testing"

// result on my PC
//
// goos: darwin
// goarch: arm64
// pkg: inno/hw4/task2
// BenchmarkArrSeparationByPrimeAndCompositeNumbersWithoutWorkerPool-12                  58          20645888 ns/op         7964293 B/op      31522 allocs/op
// BenchmarkArrSeparationByPrimeAndCompositeNumbersWithWorkerPool-12               a     193          6243485 ns/op         7964320 B/op      31526 allocs/op

func BenchmarkArrSeparationByPrimeAndCompositeNumbersWithoutWorkerPool(b *testing.B) {
	arr := make([]int, 10000)
	for i := range 10000 {
		arr[i] = i
	}
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		_, _ = ArrSeparationByPrimeAndCompositeNumbersWithoutWorkerPool(arr)
	}
}

func BenchmarkArrSeparationByPrimeAndCompositeNumbersWithWorkerPool(b *testing.B) {
	arr := make([]int, 10000)
	for i := range 10000 {
		arr[i] = i
	}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_, _ = ArrSeparationByPrimeAndCompositeNumbersWithWorkerPool(arr)
	}
}
