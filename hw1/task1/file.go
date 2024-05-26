package main

import (
	"encoding/csv"
	"errors"
	"fmt"
	"io"
	"log/slog"
	"math/rand"
	"os"
	"sync"
)

const (
	batchSize  int = 3 // num of rows we can hold in the MEM
	workerNums int = 3 // num of workers that can shuffle batch of rows
)

const testFilePattern string = "test-*.csv"

// OpenCSVFile opens CSV file by given path
func OpenCSVFile(path string) (*os.File, error) {
	f, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	return f, nil
}

func ShuffleCSVFile(f *os.File) (*os.File, error) {
	tmp, err := createTempFile()
	if err != nil {
		return nil, err
	}

	rowsBatch := make([][]string, 0, batchSize)
	rows, err := readCSVFile(tmp, f, &rowsBatch)
	if err != nil {
		return nil, fmt.Errorf("error reading file: %w", err)
	}

	shuffledRowsBatches := make([]chan [][]string, workerNums)
	for i := range workerNums {
		shuffledRowsBatches[i] = shuffleBatch(rows)
	}

	tmp, err = mergeShuffledRows(tmp, shuffledRowsBatches...)
	if err != nil {
		return nil, err
	}

	_, err = tmp.Seek(0, 0)
	if err != nil {
		return nil, fmt.Errorf("error seeking temp file: %w", err)
	}

	return tmp, nil
}

func createTempFile() (*os.File, error) {
	tmp, err := os.CreateTemp("", testFilePattern)
	if err != nil {
		return nil, fmt.Errorf("error creating temp file: %v", err)
	}
	return tmp, nil
}

// readFile reads file by batch and send this batches in the out channel
func readCSVFile(tmp, f *os.File, rowsBatch *[][]string) (<-chan [][]string, error) {
	out := make(chan [][]string)

	r := csv.NewReader(f)
	// read columns names and then write in the temp file
	columns, err := r.Read()
	if err != nil {
		return nil, fmt.Errorf("error reading columns: %v", err)
	}
	if len(columns) != 2 {
		return nil, fmt.Errorf("expected 2 columns, got %d", len(columns))
	}
	w := csv.NewWriter(tmp)
	err = w.Write(columns)
	if err != nil {
		return nil, fmt.Errorf("error writing columns in temp file: %v", err)
	}
	w.Flush()
	err = w.Error()
	if err != nil {
		return nil, fmt.Errorf("error flushing temp file: %v", err)
	}

	go func() {
		defer close(out)
		for {
			row, err := r.Read()
			if errors.Is(err, io.EOF) || err != nil {
				break
			}

			question := row[0]
			answer := row[1]
			if question == "" || answer == "" {
				continue
			}

			if len(*rowsBatch) == batchSize {
				out <- *rowsBatch
				*rowsBatch = [][]string{}
			}
			*rowsBatch = append(*rowsBatch, row)
		}
		out <- *rowsBatch
	}()

	return out, nil
}

// shuffleBatch takes certain num of rows and shuffles them
func shuffleBatch(in <-chan [][]string) chan [][]string {
	out := make(chan [][]string)

	go func() {
		defer close(out)
		for rows := range in {
			rand.Shuffle(len(rows), func(i, j int) { rows[i], rows[j] = rows[j], rows[i] })
			out <- rows
		}
	}()

	return out
}

func mergeShuffledRows(tmpFile *os.File, in ...chan [][]string) (*os.File, error) {
	row := make(chan []string, batchSize)

	wg := sync.WaitGroup{}
	wg.Add(len(in))
	for _, shuffledRowsBatches := range in {
		go func() {
			defer wg.Done()
			for rows := range shuffledRowsBatches {
				for _, r := range rows {
					row <- r
				}
			}
		}()
	}

	go func() {
		wg.Wait()
		close(row)
	}()

	err := writeCSVFile(tmpFile, row)
	if err != nil {
		return nil, err
	}

	return tmpFile, nil
}

func writeCSVFile(f *os.File, row chan []string) error {
	w := csv.NewWriter(f)
	defer w.Flush()
	for r := range row {
		err := w.Write(r)
		if err != nil {
			slog.Error("Writing record to csv", slog.String("error", err.Error()))
		}
	}
	return w.Error()
}
