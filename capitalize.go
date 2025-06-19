package goreloaded

import "strings"

func Capitalize(s string) string {
	capit := ""
	for i := 0; i < len(s); i++ {
		if i == 0 {
			capit += strings.ToUpper(string(s[i]))
		} else {
			capit += strings.ToLower(string(s[i]))
		}
	}
	return capit
}
