package main

import (
	"flag"
	"log/slog"
	"os"
)

const fileByDefault string = "./problems.csv"

var (
	path      *string = flag.String("file-path", fileByDefault, "The path to the test csv file")
	isShuffle *bool   = flag.Bool("shuffle", false, "Shuffle the test csv file")
)

func main() {
	flag.Parse()

	slog.SetLogLoggerLevel(slog.LevelError)

	slog.Info("Opening file...", slog.String("path", *path))
	csvFile, err := OpenCSVFile(*path)
	if err != nil {
		slog.Error("Opening csv file", slog.String("error", err.Error()))
		return
	}
	slog.Info("File opened", slog.String("path", *path))
	defer func() {
		slog.Info("Closing file...", slog.String("path", *path))
		err = csvFile.Close()
		if err != nil {
			slog.Error("Closing csv file", slog.String("error", err.Error()))
			return
		}
		slog.Info("File closed", slog.String("path", *path))
	}()

	if *isShuffle {
		slog.Info("Shuffling file...", slog.String("path", *path))
		csvFile, err = ShuffleCSVFile(csvFile)
		if err != nil {
			slog.Error("Shuffling file", slog.String("error", err.Error()))
			return
		}
		slog.Info("File shuffled", slog.String("path", *path))
		defer os.Remove(csvFile.Name())
	}

	slog.Info("Starting test...")
	ans, err := StartTest(csvFile)
	if err != nil {
		slog.Error("Starting test", slog.String("error", err.Error()))
		return
	}
	slog.Info("Test ended")

	PrintAnswers(ans)
}
