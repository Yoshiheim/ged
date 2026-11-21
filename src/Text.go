package src

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

var IsRunning = true

var CurrentText = ""

func ReadInput() string {
	scan := bufio.NewScanner(os.Stdin)
	if scan.Scan() {
		return strings.ReplaceAll(scan.Text(), "\\n", "\n")
	}
	return ""
}

func PrintItAndSaveFile(f *os.File) {
	str := ""
	for i := 1; i < 1000; i++ {
		switch {
		case i%80 == 0:
			str += "\n"
		case i < len(CurrentText):
			switch CurrentText[i%len(CurrentText)] {
			case '\n', '\b':
			default:
				str += string(CurrentText[i%len(CurrentText)])
			}
		default:
			str += " "
		}
	}
	if f != nil {
		f.WriteString(CurrentText)
	}
	fmt.Println(str)
}

// Prints With Color
// func ColorString(s string, color int) string {
// 	return string(fmt.Sprintf("\033[%dm%s\033[00m", 40+(color%10), s))
// }
