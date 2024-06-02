package types

import (
	"fmt"
)

type Bold struct{}

func (b Bold) Format(str string) string {
	bold := "\033[1m"
	reset := "\033[0m"
	return fmt.Sprintf("%s%s%s", bold, str, reset)
}
