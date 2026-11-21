package src

import (
	"fmt"
	"strconv"
)

func AddBinary(b []byte) string {
	buf := ""
	for _, v := range b {
		buf += "\\x" + strconv.FormatInt(int64(v), 16)
	}
	return buf
}

func AddSingleBinary(b byte) string {
	return fmt.Sprintf("\\x%d", b)
}
