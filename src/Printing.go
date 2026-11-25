package src

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

var Heigth = 20
var Width = 80

var Filename = ""

var IsRunning = true

var IsSelectedBinary = false

var CurrentText = ""

var str = ""

func ReadInput() string {
	scan := bufio.NewScanner(os.Stdin)
	if scan.Scan() {
		return strings.ReplaceAll(scan.Text(), "\\n", "\n")
	}
	return ""
}

func PrintItAndSaveFile() {
	if Heigth <= 0 {
		Heigth = 1
	}
	str = strings.Repeat(" ", (80-len(Filename))/2) + "/" + Filename + "\n"
	str += strings.Repeat("_", 80)
	for i := 0; i < Width*Heigth; i++ {
		if i%80 == 0 {
			linenum := strconv.Itoa(i / 80)
			str += "\n" + linenum + "|"
		} else if i-1 < len(CurrentText) {
			if CurrentText[i-1] > 126 || CurrentText[i-1] < 32 {
				if IsSelectedBinary {
					str += AsNumberChar(CurrentText[i-1], "\\x", HEXADECIMAL)

				}
			} else {
				str += string(CurrentText[(i-1)%len(CurrentText)])
			}
		} else {
			str += "."
		}

		// switch {
		// case i%80 == 0:
		// 	linenum := strconv.Itoa(i / 80)
		// 	str += "\n" + linenum + "|"
		// case i-1 < len(CurrentText):
		// 	if CurrentText[i-1] > 126 || CurrentText[i-1] < 32 {
		// 		str += AsNumberChar(CurrentText[i-1], "\\x", HEXADECIMAL)
		// 	} else {
		// 		str += string(CurrentText[(i-1)%len(CurrentText)])
		// 	}
		// default:
		// 	str += "."
		// }
	}

	fmt.Println(string(str))
}
