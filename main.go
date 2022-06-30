package main

import (
	"encoding/csv"
	"flag"
	"fmt"
	"log"
	"os"
	"strings"
)

func main() {
	csvFilename := flag.String("csv", "questions.csv", "a csv file for questions and answers")
	timeLimit := flag.Int("limit", 25, "")
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
	//fmt.Println(csvLines)

	problems := parseLines(csvLines)
	var correct int
	for i, p := range problems {
		fmt.Printf("problem #%d: %s= \n", i+1, p.question)
		var ans string
		fmt.Scanf("%s\n", &ans)
		if ans == p.answer {
			fmt.Println("Correct!")
			correct++
		}
	}
	fmt.Printf("you scored %d out of %d.\n", correct, len(problems))
}

func parseLines(lines [][]string) []problem {
	ret := make([]problem, len(lines))
	for i, line := range lines {
		ret[i] = problem{
			question: line[0],
			answer:   strings.TrimSpace(line[1]),
		}
	}
	return ret
}

type problem struct {
	question string
	answer   string
}
