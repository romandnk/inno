package types

import "fmt"

type Italics struct{}

func (i Italics) Format(str string) string {
	italic := "\033[3m"
	reset := "\033[0m"
	return fmt.Sprintf("%s%s%s", italic, str, reset)
}
