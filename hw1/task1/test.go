package main

import (
	"bufio"
	"encoding/csv"
	"errors"
	"fmt"
	"io"
	"os"
)

type Answers struct {
	RightCount int
	WrongCount int
}

func StartTest(f *os.File) (Answers, error) {
	ans := Answers{}

	r := csv.NewReader(f)
	// get columns names
	columns, err := r.Read()
	if err != nil {
		return ans, fmt.Errorf("error getting csv file columns names: %w", err)
	}
	if len(columns) != 2 {
		return ans, fmt.Errorf("expected 2 columns, got %d", len(columns))
	}

	scanner := bufio.NewScanner(os.Stdin)

	questionNumber := 1
	for {
		row, err := r.Read()
		if errors.Is(err, io.EOF) {
			break
		}
		if err != nil {
			return ans, fmt.Errorf("error reading csv file row: %w", err)
		}

		testQuestion := row[0]
		testAnswer := formatString(row[1])

		if testQuestion == "" || testAnswer == "" {
			continue
		}

		fmt.Printf("Question %d: %s\n", questionNumber, testQuestion)
		fmt.Printf("Your answer: ")

		scanner.Scan()
		userAnswer := formatString(scanner.Text())

		if userAnswer != testAnswer {
			ans.WrongCount++
		} else {
			ans.RightCount++
		}

		questionNumber++
	}

	return ans, nil
}

func PrintAnswers(a Answers) {
	total := a.RightCount + a.WrongCount
	fmt.Printf("Total: %d, Right: %d, Wrong: %d\n", total, a.RightCount, a.WrongCount)
}
