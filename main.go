package main

import (
	"encoding/csv"
	"flag"
	"log"
	"os"
)

func main() {
	csvFilename := flag.String("csv", "questions.csv", "a csv file for questions and answers")
	flag.Parse()
	file, err := os.Open(*csvFilename)
	if err != nil {
		log.Fatalf("unable to open file %v", *csvFilename)
	}
	reader := csv.NewReader(file)
	csvLines, err := reader.ReadAll()
	if err != nil {
		log.Fatal("Cannot read csv")
	}

	problems := parseLines(csvLines)
	for i, p := range problems {

	}
}
func parseLines(lines [][]string) []problem {
	ret := make([]problem, len(lines))
	for i, line := range lines {
		ret[i] = problem{
			question: line[0],
			answer:   line[0],
		}
	}
	return ret
}

type problem struct {
	question string
	answer   string
}
