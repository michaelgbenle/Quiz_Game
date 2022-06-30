package main

import (
	"flag"
	"os"
)

func main() {
	csvFilename := flag.String("csv", "questions.csv", "a csv file for questions and answers")
	flag.Parse()
	file, err := os.Open(*csvFilename)
}
