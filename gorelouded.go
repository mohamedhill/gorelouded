package goreloaded

import (
	"fmt"
	"os"
	"strconv"
	"strings"
	"unicode"
)

func Runponc(s rune) bool {
	if s == '.' || s == '?' || s == '!' || s == ';' || s == ':' || s == ',' {
		return true
	}
	return false
}

func normalizePunctuation(input string) string {
	runes := []rune(input)
	var result []rune

	for i := 0; i < len(runes); {
		r := runes[i]

		if Runponc(r) {

			for len(result) > 0 && result[len(result)-1] == ' ' {
				result = result[:len(result)-1]
			}

			start := i
			for i+1 < len(runes) && Runponc(runes[i+1]) {
				i++
			}
			result = append(result, runes[start:i+1]...)
			i++

			if i < len(runes) {
				result = append(result, ' ')
			}
		} else {
			result = append(result, r)
			i++
		}
	}
	return string(result)
}

func processTags(zrox []string) []string {
	for i := 0; i < len(zrox); i++ {
		if i == 0 {
			switch zrox[i] {
			case "(up)", "(cap)", "(low)", "(hex)", "(bin)":
				zrox[i] = ""
				zrox = Cleanslice(zrox)
				i--
				continue
			case "(up,", "(cap,", "(low,", "(hex,", "(bin,":
				if i+1 < len(zrox) && len(zrox[i+1]) > 1 {
					zrox[i] = ""
					zrox[i+1] = ""
					zrox = Cleanslice(zrox)
					i--
					continue
				}
			}
		}
		switch zrox[i] {
		case "(cap)":
			if i != 0 {
				zrox[i-1] = Capitalize(zrox[i-1])
				zrox[i] = ""
				zrox = Cleanslice(zrox)
				i--
				continue
			}
		case "(up)":
			if i != 0 {
				zrox[i-1] = strings.ToUpper(zrox[i-1])
				zrox[i] = ""
				zrox = Cleanslice(zrox)
				i--
				continue
			}
		case "(low)":
			if i != 0 {
				zrox[i-1] = strings.ToLower(zrox[i-1])
				zrox[i] = ""
				zrox = Cleanslice(zrox)
				i--
				continue
			}
		case "(hex)":
			if i != 0 {
				num, err := strconv.ParseInt(zrox[i-1], 16, 64)
				if err != nil {
					fmt.Println("error converting hex:", err)
				} else {
					zrox[i-1] = strconv.Itoa(int(num))
					zrox[i] = ""
					zrox = Cleanslice(zrox)
					i--
				}
			}
		case "(bin)":
			if i != 0 {
				num, err := strconv.ParseInt(zrox[i-1], 2, 64)
				if err != nil {
					fmt.Println("error converting bin:", err)
				} else {
					zrox[i-1] = strconv.Itoa(int(num))
					zrox[i] = ""
					zrox = Cleanslice(zrox)
					i--
				}
			}
		case "(cap,":
			if i != 0 && i+1 < len(zrox) && len(zrox[i+1]) > 1 {
				end, err := strconv.Atoi(zrox[i+1][:len(zrox[i+1])-1])
				if err != nil {
					continue
				}
				for k := 1; k <= end; k++ {
					if i-k >= 0 {
						zrox[i-k] = Capitalize(zrox[i-k])
					}
				}
				zrox[i] = ""
				zrox[i+1] = ""
				zrox = Cleanslice(zrox)
				i--
				continue
			} else {
				fmt.Println("Error: (cap,) tag requires a number to specify how many words to capitalize.")
				continue
			}
		case "(low,":
			if i != 0 && i+1 < len(zrox) && len(zrox[i+1]) > 1 {
				end, err := strconv.Atoi(zrox[i+1][:len(zrox[i+1])-1])
				if err != nil {
					continue
				}
				for k := 1; k <= end; k++ {
					if i-k >= 0 {
						zrox[i-k] = strings.ToLower(zrox[i-k])
					}
				}
				zrox[i] = ""
				zrox[i+1] = ""
				zrox = Cleanslice(zrox)
				i--
				continue
			} else {
				fmt.Println("Error: (low,) tag requires a number to specify how many words to lowercase.")
				continue
			}
		case "(up,":
			if i != 0 && i+1 < len(zrox) && len(zrox[i+1]) > 1 {
				end, err := strconv.Atoi(zrox[i+1][:len(zrox[i+1])-1])
				if err != nil {
					continue
				}

				for k := 1; k <= end; k++ {
					if i-k >= 0 && isWord(zrox[i-k]) {
						zrox[i-k] = strings.ToUpper(zrox[i-k])
					}
				}
				zrox[i] = ""
				zrox[i+1] = ""
				zrox = Cleanslice(zrox)
				i--
				continue
			} else {
				fmt.Println("Error: (up,) tag requires a number to specify how many words to uppercase.")
				continue
			}

		}
	}

	return zrox
}

func WriteOutput(filename string, zrox []string) error {
	var slice []byte
	for i, word := range zrox {
		slice = append(slice, []byte(word)...)
		if i != len(zrox)-1 {
			slice = append(slice, ' ')
		}
	}
	return os.WriteFile(filename, slice, 0o644)
}

func isWord(s string) bool {
	for _, r := range s {
		if unicode.IsLetter(r) {
			return true
		}
	}
	return false
}

func IsWordChar(r rune) bool {
	if unicode.IsLetter(r) || unicode.IsDigit(r) {
		return true
	}
	return false
}

func FixSingleQuotes(s string) string {
	var result []rune
	runes := []rune(s)
	i := 0

	for i < len(runes) {
		if runes[i] == '\'' {
			befor := i > 0 && IsWordChar(runes[i-1])
			next := i+1 < len(runes) && IsWordChar(runes[i+1])
			if befor && next {
				result = append(result, '\'')
				i++
				continue
			}
			j := i + 1
			for j < len(runes) {
				if runes[j] == '\'' {
					befor := j > 0 && IsWordChar(runes[j-1])
					next := j+1 < len(runes) && IsWordChar(runes[j+1])
					if !(befor && next) {
						break
					}
				}
				j++
			}
			if j < len(runes) {
				inner := runes[i+1 : j]
				start, end := 0, len(inner)
				for start < end && inner[start] == ' ' {
					start++
				}
				for end > start && inner[end-1] == ' ' {
					end--
				}
				trimmed := inner[start:end]
				result = append(result, '\'')
				result = append(result, trimmed...)
				result = append(result, '\'')
				/* result = append(result, ' ') */
				i = j + 1
			} else {
				result = append(result, '\'')
				i++
			}
		} else {
			result = append(result, runes[i])
			i++
		}
	}
	return string(result)
}

func Gorseloaded(clean string) []string {
	var zrox []string
	
	clean = Handllines(clean)
	clean = normalizePunctuation(clean)
	clean = FixSingleQuotes(clean)
	zrox = StringToSlice(clean)
	zrox = Cleanslice(zrox)
	zrox = vowels(zrox)

	return zrox
}

func Handllines(s string) string {
	lines := strings.Split(s, "\n")
	var result []string
	for _, line := range lines {
		line = normalizeSpaces(line) // Normalize spaces in each line
		words := strings.Split(line, " ")
		words = processTags(words)
		result = append(result, strings.Join(words, " "))
	}
	return strings.Join(result, "\n")
}

func isvoules(s string) bool {
	vowels := "aeiouhAEIOUH"
	for i, b := range s {
		if i == 0 && strings.ContainsRune(vowels, b) {
			return true
		}
	}

	return false
}

func vowels(t []string) []string {
	for i := 0; i < len(t); i++ {
		if i+1 < len(t) && ((t[i] == "a" && isvoules(t[i+1])) || (t[i] == "A" && isvoules(t[i+1]))) {
			t[i] += "n"
			continue
		}
	}
	return t
}

func StringToSlice(strclean string) []string {
	str := strings.Split(strclean, " ")
	return str
}

func Cleanslice(s []string) []string {
	var clean []string
	for i := 0; i < len(s); i++ {
		if s[i] != "" {
			clean = append(clean, s[i])
		}
	}
	return clean
}

func Capitalize(s string) string {
	runes := []rune(s)
	found := false

	for i, r := range runes {
		if !found && unicode.IsLetter(r) {
			runes[i] = unicode.ToUpper(r)
			found = true
		} else if found {
			runes[i] = unicode.ToLower(r)
		}
	}
	return string(runes)
}

func normalizeSpaces(s string) string {
	return strings.Join(strings.Fields(s), " ")
}
