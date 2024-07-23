package fast

import (
	"os"
	"testing"
)

func BenchmarkLogger_Info(b *testing.B) {
	f := NewLogger()
	f.SetOutput(os.Stdin)
	for i := 0; i < b.N; i++ {
		f.Info("hello world", "data1", "data2", "data3", "data4", "data5")
	}
}
