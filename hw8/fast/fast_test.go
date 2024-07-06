package fast

import "testing"

func BenchmarkLogger_Info(b *testing.B) {
	f := NewLogger()
	for i := 0; i < b.N; i++ {
		f.Info("hello world", "data1", "data2", "data3", "data4", "data5")
	}
}
