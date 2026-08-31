package splitters

import (
	"polaris/internal/types"
)

type Splitter func([]byte) []string

var Splitters = map[types.ContentType]Splitter{
	types.Text: PlainTextSplitter,
}

func Split(data []byte, contentType types.ContentType) []string {
	var splitter Splitter

	if s, ok := Splitters[contentType]; ok {
		splitter = s
	} else {
		// fallback to plain text
		splitter = PlainTextSplitter
	}

	return splitter(data)
}
