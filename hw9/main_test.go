package main

import "testing"

func BenchmarkMinEl(b *testing.B) {
	a := make([]int, 1000)
	for i := range 1000 {
		a[i] = i
	}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		MinEl(a)
	}
}

func BenchmarkMinEl2(b *testing.B) {
	a := make([]int, 1000)
	for i := range 1000 {
		a[i] = i
	}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		MinEl2(a)
	}
}

func BenchmarkMinElImproved(b *testing.B) {
	a := make([]int, 1000)
	for i := range 1000 {
		a[i] = i
	}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		MinElImproved(a)
	}
}
