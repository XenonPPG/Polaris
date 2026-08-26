package splitters

import (
	"strings"
	"unicode/utf8"
)

func PlainTextSplitter(data []byte) []string {
	if !utf8.Valid(data) {
		data = []byte(strings.ToValidUTF8(string(data), ""))
	}

	str := strings.Map(func(r rune) rune {
		if r == 0 || (r < 32 && r != '\n' && r != '\r' && r != '\t') {
			return -1
		}
		return r
	}, string(data))

	return strings.Split(str, ".")
}
