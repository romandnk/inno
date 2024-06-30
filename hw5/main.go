package main

import (
	"bufio"
	"context"
	"fmt"
	"log"
	"os"
	"os/signal"
	"strings"
	"sync"
	"syscall"
)

/*1.
1. Напишите 2 функции: 1 функция читает ввод с консоли. Ввод одного значения заканчивается нажатием клавиши enter.
Вторая функция пишет эти данные в файл. Передайте в эти функции контекст.
Используйте context и waitgroup.
*/

func main() {
	ctx, cancel := signal.NotifyContext(context.Background(), syscall.SIGTERM, syscall.SIGINT)
	defer cancel()

	f, err := openFile()
	if err != nil {
		log.Fatal(fmt.Errorf("error opening file: %w", err))
	}
	defer f.Close()

	input := make(chan string, 1)

	wg := &sync.WaitGroup{}
	wg.Add(2)
	go func() {
		defer wg.Done()
		write(ctx, f, input)
	}()
	go func() {
		defer wg.Done()
		read(ctx, input)
	}()
	wg.Wait()
	close(input)
}

func read(ctx context.Context, input chan string) {
	reader := bufio.NewReader(os.Stdin)

	text := make(chan string, 1)

	go func() {
		defer close(text)
		for {
			str, err := reader.ReadString('\n')
			str = strings.TrimRight(str, "\r\n")
			if err != nil {
				log.Printf("error reading console: %v\n", err)
				break
			}
			select {
			case <-ctx.Done():
				break
			case text <- str:
			}
		}
	}()

	for {
		select {
		case <-ctx.Done():
			return
		case str, ok := <-text:
			if !ok {
				return
			}
			input <- str
		}
	}
}

func write(ctx context.Context, f *os.File, input chan string) {
	for {
		select {
		case <-ctx.Done():
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

func openFile() (*os.File, error) {
	return os.OpenFile("./text.txt", os.O_APPEND|os.O_WRONLY, os.ModeAppend)
}
