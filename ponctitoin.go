package goreloaded
import "strings"

func Isponc(s string) (bool,rune) {
	for i := 0; i < len(s); i++ {
		if s[i] == '.' || s[i] == '?' || s[i] == '!' || s[i] == ';' || s[i] == ':' || s[i] == ',' {
			return true ,rune(s[i])
		}
	}
	return false ,' '
}

func Runponc(s rune)bool{

if s== '.' || s == '?' || s == '!' || s == ';' || s == ':' || s == ',' {
			return true 

}
return false 
}


func Index(s string) int {
	index := 0
	for i := 0; i < len(s); i++ {
		if s[i] == '.' || s[i] == '?' || s[i] == '!' || s[i] == ';' || s[i] == ':' || s[i] == ',' {
			index =strings.IndexRune(s, rune(s[i]))
		}
	}
	return index
}
 func Isflags(s string)bool{
	if s == "(up)"{
		return true
	}
	return false
 }

func isSinglePunctuation(s string) bool {
	return len(s) == 1 && strings.ContainsAny(s, ".,!?;:")
}

func isGroupedPunctuation(s string) bool {
	grouped := []string{"...", "!!", "??", "!?", "?!"}
	for _, g := range grouped {
		if s == g {
			return true
		}
	}
	return false
}

// Insert a space after punctuation unless it's followed by another punctuation
func insertSpaceAfterPunctuation(words []string) []string {
	var result []string
	for i := 0; i < len(words); i++ {
		result = append(result, words[i])

		// Add space after if:
		// - not last word
		// - current word ends in punctuation
		// - next word is not punctuation
		if i < len(words)-1 &&
			endsWithPunctuation(words[i]) &&
			!isSinglePunctuation(words[i+1]) &&
			!isGroupedPunctuation(words[i+1]) {
			result = append(result, " ")
		} else if i < len(words)-1 {
			result = append(result, " ")
		}
	}
	return result
}

func endsWithPunctuation(s string) bool {
	if len(s) == 0 {
		return false
	}
	last := s[len(s)-1:]
	return strings.ContainsAny(last, ".,!?;:")
}
func FormatPunctuation(words []string) []string {
	var result []string

	i := 0
	for i < len(words) {
		word := words[i]

		// Case 1: grouped punctuation like "...", "!?"
		if isGroupedPunctuation(word) {
			result = append(result, word)

		// Case 2: single punctuation (.,!?:;)
		} else if isSinglePunctuation(word) {
			// Attach it to the previous word
			if len(result) > 0 {
				result[len(result)-1] += word
			} else {
				result = append(result, word)
			}
		} else {
			result = append(result, word)
		}
		i++
	}

	return insertSpaceAfterPunctuation(result)
}
