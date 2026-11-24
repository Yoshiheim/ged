package src

import (
	"strconv"
)

const (
	HEXADECIMAL = 16
	OCTAL       = 8
)

func AsNumberString(b []byte, sep string, NumberSystem int) string {
	buf := ""
	for _, v := range b {
		buf += sep + strconv.FormatInt(int64(v), NumberSystem)
	}
	return buf
}

func AsNumberChar(b byte, sep string, NumberSystem int) string {
	return sep + strconv.FormatInt(int64(b), NumberSystem)
}
