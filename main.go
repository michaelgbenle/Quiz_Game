package main

import "flag"

func main() {
	csvFilename := flag.String("csv", "questions.csv", "a csv file for questions and answers")
	flag.Parse()
	_
}
