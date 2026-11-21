package src

import (
	"fmt"
	"strings"
)

var IsHTMLMode = false
var IsBinaryMode = false

func UseOpinion(s string) {
	op := strings.Split(s, "|")
	for _, v := range op {
		switch v {

		case "html":
			IsHTMLMode = !IsHTMLMode
		case "binary":
			IsBinaryMode = !IsBinaryMode

		case "quit":
			IsRunning = false
		default:
			fmt.Println(v)
			buf := v
			if IsHTMLMode {
				buf = fmt.Sprintf("<%s>%s</%s>", "h1", buf, "h1")
			}
			if IsBinaryMode {
				buf = AddBinary([]byte(buf))
			}
			CurrentText += buf
		}
	}
}
