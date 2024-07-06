package slow

import (
	"fmt"
	"os"
	"time"
)

type Logger struct{}

func NewLogger() *Logger {
	return &Logger{}
}

func (l *Logger) Info(msg string, args ...any) {
	now := time.Now().UTC()
	strArgs := ""
	for _, arg := range args {
		strArgs += fmt.Sprintf("%v ", arg)
	}
	strArgs = strArgs[0 : len(strArgs)-1]
	// использую Stdin, чтобы при бенчмарках не печаталось ничего в консоль)
	fmt.Fprintf(os.Stdin, "INFO | %s | %s | %s\n", now.Format(time.RFC3339), msg, strArgs)
}
