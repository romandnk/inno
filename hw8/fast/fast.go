package fast

import (
	"fmt"
	"io"
	"os"
	"strings"
	"time"
)

const separator = " | "

// 20 is the length of UTC RFC3339 time string
const growNum int = 3*len(separator) + 20

type level string

const (
	info level = "INFO"
)

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
