package src

import "fmt"

const (
	Reset  = "\033[0m"
	Red    = "\033[31m"
	Green  = "\033[32m"
	Yellow = "\033[33m"
	Blue   = "\033[34m"
)

// Prints With Color
func ColorString(s string, color int) string {
	return string(fmt.Sprintf("\033[%dm%s\033[00m", 40+(color%10), s))
}
