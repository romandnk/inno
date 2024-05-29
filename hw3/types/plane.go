package types

import "strings"

type Plain struct{}

func (p Plain) Format(str string) string {
	str = strings.ReplaceAll(str, "\u001B[3m", "")
	str = strings.ReplaceAll(str, "\u001B[1m", "")
	str = strings.ReplaceAll(str, "`", "") // тут он правда может убрать нужные скобки)
	str = strings.ReplaceAll(str, "\u001B[0m", "")
	return str
}
