package src

import (
	"fmt"
	"io"
	"net/http"
	"strconv"
	"strings"
)

var IsHTMLMode = false
var IsBinaryMode = false

var RepeatTimes = 3

func UseOpinion(s string) {
	op := strings.Split(s, "|")
	for _, v := range op {
		DoOpinion(v)
	}
}
func DoOpinion(v string) {
	switch v {
	case "lessrep":
		if RepeatTimes-1 <= 0 {
			RepeatTimes = 1
		} else {
			RepeatTimes--
		}
	case "morerep":

		if RepeatTimes+1 >= 100 {
			RepeatTimes = 100
		} else {
			RepeatTimes++
		}
	case "html":
		IsHTMLMode = !IsHTMLMode
	case "binary":
		IsBinaryMode = !IsBinaryMode
	case "quit":
		IsRunning = false
	case "clear":
		CurrentText = ""
	case "sizemore":
		Heigth++
	case "sizeless":
		Heigth--
	default:
		switch {
		case strings.HasPrefix(v, "remove:"):
			n, err := strconv.Atoi(v[7:])
			fmt.Println(v[7:])

			if err != nil {
				return
			} else {
				if n > len(CurrentText) {
					CurrentText = ""
				} else {
					CurrentText = CurrentText[:len(CurrentText)-(n)]
				}
			}
			return
		case strings.HasPrefix(v, "rep:"):
			for i := 0; i < 10; i++ {
				DoOpinion(v[4:])
			}
		case strings.HasPrefix(v, "https://"):
			r, err := http.Get(v)
			if err != nil {
				return
			}
			b, err := io.ReadAll(r.Body)
			if err != nil {
				return
			}
			CurrentText += string(b)
		}

		buf := v
		if IsHTMLMode {
			buf = fmt.Sprintf("<%s>%s</%s>", "h1", buf, "h1")
		}
		if IsBinaryMode {
			buf = AsHex([]byte(buf), "\\x")
		}
		CurrentText += buf
	}
}
