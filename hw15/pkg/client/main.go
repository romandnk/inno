package main

import (
	"bufio"
	"context"
	"fmt"
	"gopkg.in/yaml.v3"
	"log"
	"net/http"
	"os"
	"os/signal"
	"strings"
	"syscall"
)

const address = "ws://localhost:8001"

type config struct {
	Token string `yaml:"token"`
}

func readToken() (string, error) {
	file, err := os.ReadFile("pkg/client/config.yaml")
	if err != nil {
		return "", err
	}
	var cfg config
	err = yaml.Unmarshal(file, &cfg)
	if err != nil {
		return "", err
	}
	return cfg.Token, nil
}

func main() {
	token, err := readToken()
	if err != nil {
		log.Fatalf("error reading token: %v", err)
	}

	ctx, cancel := signal.NotifyContext(context.Background(), syscall.SIGTERM, syscall.SIGKILL, syscall.SIGINT)
	defer cancel()

	cl, err := newClient(ctx, address, http.Header{
		"Authorization": []string{token},
	})
	if err != nil {
		log.Fatalf("error creating client: %v", err)
	}
	defer cl.Close()

	for {
		fmt.Println("Меню:")
		fmt.Println("1. Создать новый чат с другим пользователем")
		fmt.Println("2. Войти в чат с пользователем")
		fmt.Println("Введите ваш выбор (для выхода введите 'exit' или нажмите Ctrl+C):")

		reader := bufio.NewReader(os.Stdin)
		choice, _ := reader.ReadString('\n')
		choice = strings.TrimSpace(choice)

		switch choice {
		case "1":
			err = createNewChat(cl)
			if err != nil {
				log.Fatalf("error creating chat: %v", err)
			}
		case "2":
			err = enterChat(cl)
			if err != nil {
				log.Fatal(err.Error())
			}
		case "exit":
			fmt.Println("Выход...")
			return
		default:
			fmt.Println("Неверный выбор, попробуйте снова.")
		}
	}
}
