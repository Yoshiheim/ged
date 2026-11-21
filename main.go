package main

import (
	"ged/src"
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
	f, err := os.Create(src.Filename)
	if err != nil {
		log.Fatalln(err)
	}
	defer f.Close()
	for src.IsRunning {
		src.PrintItAndSaveFile(f)
		src.UseOpinion(src.ReadInput())
	}
}

func main() {
	GEDInit()
}
