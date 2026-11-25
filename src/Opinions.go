package src

import (
	"fmt"
	"strings"
)

var HTMLTag = "h1"

var IsHTMLMode = false
var IsHexMode = false

var RepeatTimes = 3

func UseOpinion(s string) {
	op := strings.Split(s, "|")
	for _, v := range op {
		DoOpinion(strings.ReplaceAll(v, "\n", ""))
	}
}

func DoOpinion(v string) {
	if !ReadText(v) {
		buf := v
		switch {
		case IsHTMLMode:
			buf = fmt.Sprintf("<%s>%s</%s>", HTMLTag, buf, HTMLTag)
		case IsHexMode:
			buf = AsNumberString([]byte(buf), "\\x", HEXADECIMAL)
    default:
      CurrentText += buf
	}
}
