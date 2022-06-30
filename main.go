package main

import (
	"flag"
	"log"
	"os"
)

func main() {
	csvFilename := flag.String("csv", "questions.csv", "a csv file for questions and answers")
	flag.Parse()
	file, err := os.Open(*csvFilename)
	if err != nil {
		log.Fatal("unable to open file")
	}
}
