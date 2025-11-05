package utils

import "strings"

func MultiLine(lines ...string) string {
	return strings.Join(lines, "\n")
}
