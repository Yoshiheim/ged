package src

import (
	"os"
	"strings"
)

var Prefixses = []string{
	"settag",
	"rep",
	"open",
}

var EndSuffix = "end"

var SwitchFunc = map[string]func(){
	"reverse": func() {
		CurrentText = ReverseString(CurrentText)
	},
	"hex":     func() { IsHexMode = true },
	"hexend":  func() { IsHexMode = false },
	"html":    func() { IsHTMLMode = true },
	"htmlend": func() { IsHTMLMode = false },
	"quit":    func() { IsRunning = false },

	"select":    func() { IsSelectedBinary = true },
	"selectend": func() { IsSelectedBinary = false },
}

// var PrefixsesFunc = map[string]func(v string){
// 	"rep:": func(v string) {
// 		for i := 0; i < 10; i++ {
// 			DoOpinion(v[4:])
// 		}
// 	},
// 	"open:": func(v string) {
// 		//this open special tgs script(tgs means "text generator script"), this tgs script is your input
// 		script, err := os.ReadFile(v[5:])
// 		if err != nil {
// 			return
// 		}
// 		UseOpinion(string(script))
// 	},
// 	"remove:": func(v string) {
// 		n, err := strconv.Atoi(v[7:])
// 		fmt.Println(v[7:])
// 		if err != nil {
// 			return
// 		} else {
// 			if n > len(CurrentText) {
// 				CurrentText = ""
// 			} else {
// 				CurrentText = CurrentText[:len(CurrentText)-(n)]
// 			}
// 		}
// 	},
// }

func ReverseString(s string) string {
	c := ""
	for s != "" {
		c += string(s[len(s)-1])
		s = s[:len(s)-1]
	}
	return c
}

func CheckPrefix(s string, prefix string) (bool, string) {
	prefix += ":"
	if strings.HasPrefix(s, prefix) {
		return true, s[len(prefix):]
	}
	return false, ""
}

func ReadText(v string) bool {
	f, ok := SwitchFunc[v]
	if ok {
		f()
		return true
	}
	for _, prefix := range Prefixses {
		// REMEBER:
		// var wtih name code: text after prefix.
		// var wtih name prefix.
		// and var wtih v is whole text.
		hasPrefix, code := CheckPrefix(v, prefix)
		if hasPrefix {
			switch prefix {
			case "rep":
				for i := 0; i < 10; i++ {
					DoOpinion(code)
				}
				return true
			case "open":
				//this open special tgs script(tgs means "text generator script"), this tgs script is your input
				script, err := os.ReadFile(v[5:])
				if err != nil {
					return false
				}
				UseOpinion(string(script))
				return true
			default:
				CurrentText += v
			}

			// if prefix == "rep" {
			// 	for i := 0; i < 10; i++ {
			// 		DoOpinion(code)
			// 	}
			// } else {
			// 	CurrentText += v
			// }
		}
	}
	return false

}
