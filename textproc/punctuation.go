package textproc


// Runponc returns true if the rune is considered punctuation
func Runponc(r rune) bool {
	return r == '.' || r == '?' || r == '!' || r == ';' || r == ':' || r == ','
}

// normalizePunctuation cleans spacing around punctuation
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
