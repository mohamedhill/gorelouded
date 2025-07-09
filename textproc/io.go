package textproc

import (
	"os"
	"strings"
)

func WriteOutput(filename string, words []string) error {
	content := strings.Join(words, " ")
	return os.WriteFile(filename, []byte(content), 0o644)
}
