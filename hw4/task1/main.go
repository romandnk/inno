package main

import (
	"bufio"
	"context"
	"fmt"
	"log"
	"os"
	"os/signal"
	"sync"
	"syscall"
)

/*1. Напишите 2 функции:
	Первая функция читает ввод с консоли. Ввод одного значения заканчивается нажатием клавиши enter.
	Вторая функция пишет эти данные в файл. Свяжите эти функции каналом.
Работа приложения должна завершится при нажатии клавиш ctrl+c с кодом 0. */

func main() {
	ctx, cancel := signal.NotifyContext(context.Background(), syscall.SIGINT, syscall.SIGTERM)
	defer cancel()

	f, err := openFile()
	if err != nil {
		log.Fatal(fmt.Errorf("error opening file: %w", err))
	}

	input := make(chan string)

	go func() {
		<-ctx.Done()
		close(input)
		os.Exit(0)
	}()

	go write(f, input)
	read(ctx, input)
}

func read(ctx context.Context, input chan string) {
	scanner := bufio.NewScanner(os.Stdin)

	wg := sync.WaitGroup{}
	wg.Add(1)
	go func() {
		defer wg.Done()
		for scanner.Scan() {
			select {
			case <-ctx.Done():
				return
			default:
			}
			text := scanner.Text()
			input <- text
		}
	}()
	wg.Wait()
}

func write(f *os.File, input chan string) {
	for text := range input {
		_, err := f.WriteString(text + "\n")
		if err != nil {
			log.Printf("error writing to file: %v\n", err)
		}
	}
}

func openFile() (*os.File, error) {
	return os.OpenFile("./text.txt", os.O_APPEND|os.O_WRONLY, os.ModeAppend)
}
