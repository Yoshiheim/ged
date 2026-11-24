package src

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"strconv"
	"strings"
)

var HTMLTag = "h1"

var IsHTMLMode = false
var IsBinaryMode = false

var RepeatTimes = 3

func UseOpinion(s string) {
	op := strings.Split(s, "|")
	for _, v := range op {
		DoOpinion(strings.ReplaceAll(v, "\n", ""))
	}
}

func DoOpinion(v string) {
	switch v {
	case "copymode":
		IsCopyMode = !IsCopyMode
		return
	case "copydrop":
		CurrentText += CopyBuffer
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
		case strings.HasPrefix(v, "settag:"):
			HTMLTag = v[7:]
		case strings.HasPrefix(v, "rep:"):
			for i := 0; i < 10; i++ {
				DoOpinion(v[4:])
			}
		case strings.HasPrefix(v, "open:"):
			//this open special tgs script(tgs means "text generator script"), this tgs script is your input
			script, err := os.ReadFile(v[5:])
			if err != nil {
				return
			}
			UseOpinion(string(script))
			return
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
			buf = fmt.Sprintf("<%s>%s</%s>", HTMLTag, buf, HTMLTag)
		} else if IsBinaryMode {
			buf = AsNumberString([]byte(buf), "\\x", HEXADECIMAL)
		}
		if IsCopyMode {
			CopyBuffer += buf
		} else {
			CurrentText += buf
		}
	}
}
