package main

import "strings"

// formatStrings trims all leading and trailing white spaces and convert it to lower case
func formatString(str string) string {
	return strings.ToLower(strings.TrimSpace(str))
}
