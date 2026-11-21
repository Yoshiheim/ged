package src

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

var SizeOfDisplay = 10

var Filename = ""

var IsRunning = true

var CurrentText = ""

func ReadInput() string {
	scan := bufio.NewScanner(os.Stdin)
	if scan.Scan() {
		return strings.ReplaceAll(scan.Text(), "\\n", "\n")
	}
	return ""
}

func PrintItAndSaveFile() {
	if SizeOfDisplay <= 0 {
		SizeOfDisplay = 1
	}
	str := strings.Repeat(" ", (80-len(Filename))/2) + "/" + Filename + "\n"
	str += strings.Repeat("_", 80)
	for i := 0; i < 80*SizeOfDisplay; i++ {
		switch {
		case i/80 == 10:
		case i%80 == 0:
			linenum := strconv.Itoa(i / 80)
			str += "\n" + linenum + "|"
		case i-1 < len(CurrentText):
			switch CurrentText[i%len(CurrentText)] {
			case '\n', '\b':
			default:
				str += string(CurrentText[(i-1)%len(CurrentText)])
			}
		default:
			str += "."
		}
	}
	fmt.Println(str)
}
