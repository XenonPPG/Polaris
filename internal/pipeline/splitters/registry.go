package splitters

import (
	"polaris/internal/db"
)

type Splitter func([]byte) []string

var Splitters = map[db.ContentType]Splitter{
	db.Text: PlainTextSplitter,
}

func Split(data []byte, contentType db.ContentType) []string {
	var splitter Splitter

	if s, ok := Splitters[contentType]; ok {
		splitter = s
	} else {
		// fallback to plain text
		splitter = PlainTextSplitter
	}

	return splitter(data)
}
