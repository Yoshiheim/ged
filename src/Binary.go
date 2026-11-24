package src

import (
	"strconv"
)

func AsHex(b []byte, sep string) string {
	buf := ""
	for _, v := range b {
		buf += sep + strconv.FormatInt(int64(v), 16)
	}
	return buf
}

func AsSingleHex(b byte, sep string) string {
	return sep + strconv.FormatInt(int64(b), 16)
}
