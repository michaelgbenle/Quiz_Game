package main

import (
	"encoding/csv"
	"flag"
	"fmt"
	"log"
	"os"
	"strings"
	"time"
)

func main() {
	csvFilename := flag.String("csv", "questions.csv", "a csv file for questions and answers")
	timeLimit := flag.Int("limit", 25, "time limit for your question")
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
	timer := time.NewTimer(time.Duration(*timeLimit) * time.Second)

	var correct int
	for i, p := range problems {
		fmt.Printf("problem #%d: %s= ", i+1, p.question)
		answerCh := make(chan string)
		go func() {
			var ans string
			fmt.Scanf("%s\n", &ans)
			answerCh <- ans
		}()
		select {
		case <-timer.C:
			fmt.Printf("\n You scored %d out of %d.\n", correct, len(problems))
			return
		case ans := <-answerCh:

			if ans == p.answer {
				fmt.Println("Correct!")
				correct++
			}
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
