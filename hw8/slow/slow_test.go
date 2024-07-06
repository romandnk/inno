package slow

import "testing"

func BenchmarkLogger_Info(b *testing.B) {
	s := NewLogger()
	for i := 0; i < b.N; i++ {
		s.Info("hello world", "data1", "data2", "data3", "data4", "data5")
	}
}
