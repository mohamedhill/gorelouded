package textproc

import (
	"fmt"
	"strconv"
	"strings"
)

// processTags modifies a slice of strings based on tag instructions
func processTags(words []string) []string {
	for i := 0; i < len(words); i++ {
		switch words[i] {
		case "(cap)", "(up)", "(low)":
			if i > 0 {
				switch words[i] {
				case "(cap)":
					words[i-1] = Capitalize(words[i-1])
				case "(up)":
					words[i-1] = strings.ToUpper(words[i-1])
				case "(low)":
					words[i-1] = strings.ToLower(words[i-1])
				}
			}
			words[i] = ""
			words = Cleanslice(words)
			i--
		case "(hex)", "(bin)":
			if i > 0 {
				var base int
				if words[i] == "(hex)" {
					base = 16
				} else {
					base = 2
				}
				num, err := strconv.ParseInt(words[i-1], base, 64)
				if err == nil {
					words[i-1] = strconv.Itoa(int(num))
					words[i] = ""
					words = Cleanslice(words)
					i--
				}
			}
		case "(cap,", "(up,", "(low,":
			if i+1 < len(words) && strings.HasSuffix(words[i+1], ")") {
				countStr := strings.TrimSuffix(words[i+1], ")")
				count, err := strconv.Atoi(countStr)
				if err != nil {
					fmt.Println("Invalid number in tag:", words[i+1])
					continue
				}
				for j := 1; j <= count && i-j >= 0; j++ {
					switch words[i] {
					case "(cap,":
						words[i-j] = Capitalize(words[i-j])
					case "(up,":
						words[i-j] = strings.ToUpper(words[i-j])
					case "(low,":
						words[i-j] = strings.ToLower(words[i-j])
					}
				}
				words[i], words[i+1] = "", ""
				words = Cleanslice(words)
				i--
			}
		}
	}
	return words
}
