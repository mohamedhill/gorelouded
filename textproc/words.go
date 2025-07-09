package textproc

import (
	"strings"
	"unicode"
)

func Capitalize(s string) string {
	runes := []rune(s)
	for i, r := range runes {
		if unicode.IsLetter(r) {
			runes[i] = unicode.ToUpper(r)
			for j := i + 1; j < len(runes); j++ {
				runes[j] = unicode.ToLower(runes[j])
			}
			break
		}
	}
	return string(runes)
}

func Cleanslice(words []string) []string {
	var result []string
	for _, w := range words {
		if w != "" {
			result = append(result, w)
		}
	}
	return result
}

func isWord(s string) bool {
	for _, r := range s {
		if unicode.IsLetter(r) {
			return true
		}
	}
	return false
}

func isVowelStart(s string) bool {
	vowels := "aeiouhAEIOUH"
	return len(s) > 0 && strings.ContainsRune(vowels, rune(s[0]))
}

func vowels(words []string) []string {
	for i := 0; i < len(words)-1; i++ {
		if (words[i] == "a" || words[i] == "A") && isVowelStart(words[i+1]) {
			words[i] += "n"
		}
	}
	return words
}
