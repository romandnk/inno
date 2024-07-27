package slow

import (
	"os"
	"testing"
)

func BenchmarkLogger_Info(b *testing.B) {
	s := NewLogger()
	s.SetOutput(os.Stdin)
	for i := 0; i < b.N; i++ {
		s.Info("hello world", "data1", "data2", "data3", "data4", "data5")
	}
}
