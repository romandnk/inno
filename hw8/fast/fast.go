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

type Logger struct{}

func NewLogger() *Logger {
	return &Logger{}
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
	// использую Stderr, чтобы при бенчмарках не печаталось ничего в консоль)
	io.Copy(os.Stdin, strings.NewReader(builder.String()))
}
