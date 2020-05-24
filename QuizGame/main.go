package main

import (
	"encoding/csv"
	"flag"
	"fmt"
	"os"
	"strings"
)

type problem struct {
	question string
	answer   string
}

func main() {

	// Set a flag for CLI use
	var csvFilename = flag.String("csv", "problems.csv", "A csv filename in quotes where the file is structured in the format 'question,answer'")
	flag.Parse()

	// Open the file
	file, err := os.Open(*csvFilename)
	if err != nil {
		errString := fmt.Sprintf("Failed to open file with name: %s ", *csvFilename)
		exit(errString)
	}

	r := csv.NewReader(file)
	lines, err := r.ReadAll()
	if err != nil {
		exit("Could not parse csv file")
	}

	problems := parseLines(lines)
	startQuizWithProblems(problems)
}

// Takes the lines from the CSV and turns them into Problem model objects
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

// Starts the quiz with the problems after parsed from the CSV
func startQuizWithProblems(problems []problem) {
	var numCorrect int
	for i, p := range problems {
		problemString := fmt.Sprintf("Problem #%d: %s =", i+1, p.question)
		var answer string

		fmt.Println(problemString)
		fmt.Scanf("%s\n", &answer)
		isCorrect := checkAnswer(answer, p.answer)
		if isCorrect {
			numCorrect++
		}
	}

	fmt.Printf("You scored %d out of %d.\n", numCorrect, len(problems))
}

// Checks the user answer vs the answer from the CSV for a particular problem
func checkAnswer(userAnswer string, answer string) bool {
	if userAnswer == answer {
		return true
	}

	return false
}

// Exits the program with a custom message.
func exit(msg string) {
	fmt.Printf("Exiting program: %s \n", msg)
	os.Exit(1)
}
