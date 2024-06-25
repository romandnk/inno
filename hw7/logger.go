package main

import (
	"fmt"
	"os"
	"strings"
	"time"
)

type SlowLogger struct{}

func (l *SlowLogger) Info(msg string, values ...any) {
	now := time.Now().UTC().Format(time.RFC3339Nano)
	output := fmt.Sprintf("%s | %s | %s", now, LevelInfo, msg)
	for _, val := range values {
		output += fmt.Sprintf(" %v", val)
	}
	fmt.Fprintf(os.Stdin, output)
}

type FastLogger struct{}

func (l FastLogger) Info(msg string, values ...any) {
	now := time.Now().UTC().Format(time.RFC3339Nano)
	output := strings.Builder{}
	estimatedSize := len(now) + 3 + len(msg) + len(values)*10
	output.Grow(estimatedSize)

	// Предположим, LevelInfo - это строковая константа "INFO"
	output.WriteString(now)
	output.WriteString(" | INFO | ")
	output.WriteString(msg)

	for _, val := range values {
		output.WriteString(" ")
		output.WriteString(fmt.Sprint(val))
	}

	// Замена прямого вызова os.Stdin.WriteString на os.Stdout.Write
	// Если осознанно используется Stdin, необходимо уточнить для чего именно
	os.Stdin.Write([]byte(output.String()))
}
