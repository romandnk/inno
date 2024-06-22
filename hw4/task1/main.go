package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"os/signal"
	"syscall"
)

/*1. Напишите 2 функции:
	Первая функция читает ввод с консоли. Ввод одного значения заканчивается нажатием клавиши enter.
	Вторая функция пишет эти данные в файл. Свяжите эти функции каналом.
Работа приложения должна завершится при нажатии клавиш ctrl+c с кодом 0. */

const filePath string = "./text.txt"

func main() {
	sigs := make(chan os.Signal, 1)
	signal.Notify(sigs, syscall.SIGINT, syscall.SIGTERM)

	f, err := openFile(filePath)
	if err != nil {
		log.Fatal(fmt.Errorf("error opening file: %w", err))
	}
	defer f.Close()

	input := make(chan string, 1)
	done := make(chan struct{})

	go func() {
		<-sigs
		close(done)
	}()

	go write(done, f, input)

	read(done, input)
	close(input)
}

func read(done chan struct{}, input chan string) {
	scanner := bufio.NewScanner(os.Stdin)

	text := make(chan string, 1)

	go func() {
		defer close(text)
		for scanner.Scan() {
			text <- scanner.Text()
		}
	}()

	for {
		select {
		case <-done:
			return
		case t, ok := <-text:
			if !ok {
				return
			}
			input <- t
		}
	}
}

func write(done chan struct{}, f *os.File, input chan string) {
	for {
		select {
		case <-done:
			return
		case text, ok := <-input:
			if !ok {
				return
			}

			_, err := f.WriteString(text + "\n")
			if err != nil {
				log.Printf("error writing to file: %v\n", err)
				return
			}
		}
	}
}

func openFile(name string) (*os.File, error) {
	return os.OpenFile(name, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
}
