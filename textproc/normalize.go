package textproc

import "strings"

func Gorseloaded(input string) []string {
	input = Handllines(input)
	input = normalizePunctuation(input)
	input = FixSingleQuotes(input)

	words := strings.Fields(input)
	words = processTags(words)
	words = Cleanslice(words)
	words = vowels(words)
	return words
}

func Handllines(s string) string {
	lines := strings.Split(s, "\n")
	for i, line := range lines {
		lines[i] = normalizeSpaces(line)
	}
	return strings.Join(lines, "\n")
}

func normalizeSpaces(s string) string {
	return strings.Join(strings.Fields(s), " ")
}
