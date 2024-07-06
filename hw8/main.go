package main

import (
	"fmt"
	"inno/hw8/fast"
	"inno/hw8/slow"
)

// Написать 2 простейших реализации логгер:
//
// - одного метода Info достаточно
// - ставим дату перед сообщением
//
// Один должен быть максимально не эффективным(смотрим бенчмарки),
// другой эффективный

type Logger interface {
	Info(msg string, args ...any)
}

// fast
// BenchmarkLogger_Info-12          2074530               575.5 ns/op           274 B/op         10 allocs/op
// slow
// BenchmarkLogger_Info-12          1605933               755.1 ns/op           256 B/op         14 allocs/op

func main() {
	var s Logger = slow.NewLogger()
	var f Logger = fast.NewLogger()
	s.Info("slow", fmt.Sprintf("msg 1"), fmt.Sprintf("msg 2"))
	f.Info("fast", fmt.Sprintf("msg 1"), fmt.Sprintf("msg 2"))
}
