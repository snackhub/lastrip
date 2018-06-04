package parser

import (
	"bytes"
	"errors"
)

func SliceToSlices(data []byte) ([]string, error) {
	// trim space
	data = bytes.TrimSpace(data)
	var parameters []string
	var i int
	for i < len(data) {
		value, next := FirstStringAt(data[:], i)
		if next == 0 {
			return nil, errors.New("invalid parameters")
		}
		i += next
		parameters = append(parameters, string(value))
	}

	return parameters, nil
}

var numbertable = [...]int{
	-1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1,
	-1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1,
	-1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1,
	0, 1, 2, 3, 4, 5, 6, 7, 8, 9, -1, -1, -1, -1, -1, -1,
	-1, 10, 11, 12, 13, 14, 15, -1, -1, -1, -1, -1, -1, -1, -1, -1,
	-1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1,
	-1, 10, 11, 12, 13, 14, 15, -1, -1, -1, -1, -1, -1, -1, -1, -1,
	-1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1,
	-1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1,
	-1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1,
	-1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1,
	-1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1,
	-1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1,
	-1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1,
	-1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1,
	-1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1,
}

func BytesToInt64(bs []byte) (result int64, err error) {
	sign := int64(1)
	switch bs[0] {
	case '+':
		sign = 1
		bs = bs[1:]
	case '-':
		sign = -1
		bs = bs[1:]
	}
	s := decimalSystem(bs)
	if s == 0 {
		return 0, errors.New("")
	}
	if s == 16 {
		bs = bs[2:]
	}

	for _, v := range bs {
		d := int64(numbertable[v])
		if d == -1 {
			return 0, errors.New("invalid number")
		}
		if d >= s {
			return 0, errors.New("overflow number")
		}
		result = result*s + d
	}
	result *= sign
	return
}

func decimalSystem(bs []byte) int64 {
	length := len(bs)
	if length == 0 {
		return 0
	}
	if length == 1 {
		return 10
	}
	if bs[0] > '1' && bs[0] <= '9' {
		return 10
	}
	if bs[0] != '0' {
		return 0
	}

	if (bs[1] | ' ') != 'x' {
		return 8
	}
	if length > 2 {
		return 16
	}
	return 0
}

// FirstStringAt return first string and end index
func FirstStringAt(data []byte, pos int) (string, int) {
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
