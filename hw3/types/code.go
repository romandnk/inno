package types

import "fmt"

type Code struct{}

func (c Code) Format(str string) string {
	return fmt.Sprintf("`%s`", str)
}
