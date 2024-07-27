package slow

import (
	"fmt"
	"io"
	"os"
	"time"
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

func (l *Logger) Info(msg string, args ...any) {
	now := time.Now().UTC()
	strArgs := ""
	for _, arg := range args {
		strArgs += fmt.Sprintf("%v ", arg)
	}
	strArgs = strArgs[0 : len(strArgs)-1]

	fmt.Fprintf(l.output, "INFO | %s | %s | %s\n", now.Format(time.RFC3339), msg, strArgs)
}
