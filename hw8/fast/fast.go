package fast

import (
	"fmt"
	"io"
	"os"
	"strings"
	"sync"
	"time"
)

const separator = " | "

// 20 is the length of UTC RFC3339 time string
const growNum int = 3*len(separator) + 20

type level string

const (
	info level = "INFO"
)

var pool = &sync.Pool{
	New: func() any {
		return new(strings.Builder)
	},
}

type Logger struct {
	output io.Writer
}

func NewLogger() *Logger {
	return &Logger{
		output: os.Stdout,
	}
}

func (l *Logger) SetOutput(output io.Writer) {
	l.output = output
}

func (l Logger) Info(msg string, args ...any) {
	now := time.Now().UTC()
	builder := strings.Builder{}
	builder.Grow(len(info) + growNum + len(msg))
	builder.WriteString(string(info))
	builder.WriteString(separator)
	builder.WriteString(now.Format(time.RFC3339))
	builder.WriteString(separator)
	builder.WriteString(msg)
	builder.WriteString(separator)
	for i, arg := range args {
		data := fmt.Sprintf("%v", arg)
		builder.Grow(len(data) + 1)
		builder.WriteString(data)
		if i != len(args)-1 {
			builder.WriteString(" ")
		} else {
			builder.WriteString("\n")
		}
	}

	io.Copy(l.output, strings.NewReader(builder.String()))
}

// slower
//BenchmarkLogger_Info-12          2010916               590.7 ns/op           274 B/op         10 allocs/op
//func (l Logger) Info(msg string, args ...any) {
//	builder := pool.Get().(*strings.Builder)
//	defer pool.Put(builder)
//
//	builder.Reset()
//
//	now := time.Now().UTC()
//	builder.Grow(len(info) + growNum + len(msg))
//	builder.WriteString(string(info))
//	builder.WriteString(separator)
//	builder.WriteString(now.Format(time.RFC3339))
//	builder.WriteString(separator)
//	builder.WriteString(msg)
//	builder.WriteString(separator)
//	for i, arg := range args {
//		builder.WriteString(fmt.Sprintf("%v", arg))
//		if i != len(args)-1 {
//			builder.WriteString(" ")
//		}
//	}
//	builder.WriteString("\n")
//
//	io.Copy(l.output, strings.NewReader(builder.String()))
//}

// slower
//BenchmarkLogger_Info-12          1926086               623.7 ns/op           338 B/op         10 allocs/op
//func (l Logger) Info(msg string, args ...any) {
//	now := time.Now().UTC()
//	buffer := bytes.Buffer{}
//	buffer.Grow(len(info) + growNum + len(msg))
//	buffer.WriteString(string(info))
//	buffer.WriteString(separator)
//	buffer.WriteString(now.Format(time.RFC3339))
//	buffer.WriteString(separator)
//	buffer.WriteString(msg)
//	buffer.WriteString(separator)
//
//	for _, arg := range args {
//		buffer.WriteString(fmt.Sprintf("%v", arg))
//		buffer.WriteString(" ")
//	}
//	buffer.Truncate(buffer.Len() - 1) // delete ' '
//	buffer.WriteString("\n")
//
//	l.output.Write(buffer.Bytes())
//}
