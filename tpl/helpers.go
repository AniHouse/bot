package tpl

import (
	"bytes"
)

func normalizeSpaces(in []byte) []byte {
	var result []byte

	for _, line := range bytes.Split(in, []byte("\n")) {
		line = append(bytes.TrimSpace(line), byte('\n'))
		if string(line) == "\n" {
			continue
		}
		if string(line) == "[EOL]\n" {
			line = []byte("\n")
		}
		result = append(result, line...)
	}

	return result
}
