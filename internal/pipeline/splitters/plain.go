package splitters

import "strings"

func PlainTextSplitter(data []byte) []string {
	return strings.Split(string(data), ".")
}
