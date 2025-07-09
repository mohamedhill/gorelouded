package textproc

import "unicode"

// FixSingleQuotes fixes quotes surrounding non-word characters
func FixSingleQuotes(s string) string {
	var result []rune
	runes := []rune(s)
	i := 0

	for i < len(runes) {
		if runes[i] == '\'' {
			before := i > 0 && IsWordChar(runes[i-1])
			next := i+1 < len(runes) && IsWordChar(runes[i+1])
			if before && next {
				result = append(result, '\'')
				i++
				continue
			}
			j := i + 1
			for j < len(runes) {
				if runes[j] == '\'' {
					before := j > 0 && IsWordChar(runes[j-1])
					if j+1 >= len(runes) {
						break
					}
					next := IsWordChar(runes[j+1])
					if !(before && next) {
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

func IsWordChar(r rune) bool {
	return unicode.IsLetter(r) || unicode.IsDigit(r)
}
