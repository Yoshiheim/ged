package main

import (
	"ged/src"
	"io"
	"log"
	"os"
)

const Offset = 25

func GEDInit() {
	// cursor := 2
	if len(os.Args) >= 2 {
		src.Filename = os.Args[1]
	} else {
		src.Filename = "index.txt"
	}
	f, err := os.Open(src.Filename)
	if err != nil {
		log.Fatalln(err)
	}
	defer f.Close()
	r, err := io.ReadAll(f)
	if err != nil {
		log.Fatalln(err)
	}
	src.CurrentText = string(r)
	for src.IsRunning {
		src.PrintItAndSaveFile()
		src.UseOpinion(src.ReadInput())
	}
	f.WriteString(src.CurrentText)
}

func main() {
	GEDInit()
}
