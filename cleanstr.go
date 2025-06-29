package goreloaded

import "strings"

func CleanStr(s string) string {
	slc := strings.Fields(s)
	return strings.Join(slc, " ")
}
