package main

import (
	"encoding/csv"
	"flag"
	"fmt"
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
	fmt.Println(csvLines)
}
func parseLines(lines [][]string) []problem {}

type problem struct {
	question string
	answer   string
}
