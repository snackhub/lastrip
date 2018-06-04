package lastrip

import (
	"bytes"
	"errors"
)

func ParseSliceToSlices(data []byte) ([]string, error) {
	// trim space
	data = bytes.TrimSpace(data)
	var parameters []string
	var i int
	for i < len(data) {
		value, next := FindFirstStringFromPosition(data[:], i)
		if next == 0 {
			return nil, errors.New("invalid parameters")
		}
		i += next
		parameters = append(parameters, string(value))
	}

	return parameters, nil
}

// return first string and end index
func FindFirstStringFromPosition(data []byte, pos int) (string, int) {
	var i int
	var c byte

	if pos >= len(data) {
		return "", 0
	}
	data = data[pos:]
	for i, c = range data {
		if !isSpace(c) {
			break
		}
	}
	begin := i
	terminator := byte(' ')

	if data[i] == '"' || data[i] == '\'' {
		terminator = data[i]
		i++
	}

	for ; i < len(data); i++ {
		if terminator == ' ' && isSpace(data[i]) {
			return string(data[begin:i]), i + 1
		}
		if !isSpace(data[i]) && data[i] == terminator && (i == len(data)-1 || isSpace(data[i+1])) {
			return string(data[begin : i+1]), i + 1
		}
	}

	return string(data[begin:]), len(data)
}

func isSpace(c byte) bool {
	return c == ' ' || c == '\n' || c == '\t'
}

func isQuotation(c byte) bool {
	return c == '\'' || c == '"'
}
