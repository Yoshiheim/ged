package main

import (
	"ged/src"
	"log"
	"os"
)

const Offset = 25

func GEDInit() {

	filename := ""
	// cursor := 2
	if len(os.Args) >= 2 {
		filename = os.Args[1]
	} else {
		filename = "index.txt"
	}
	f, err := os.Create(filename)
	if err != nil {
		log.Fatalln(err)
	}
	defer f.Close()
	for src.IsRunning {
		src.UseOpinion(src.ReadInput())
		src.PrintItAndSaveFile(f)
	}
}

func main() {
	GEDInit()
}
