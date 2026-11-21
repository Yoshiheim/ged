package src

import "strconv"

func AddBinary(b []byte) string {
	buf := ""
	for _, v := range b {
		buf += "\\x" + strconv.Itoa(int(v))
	}
	return buf
}
