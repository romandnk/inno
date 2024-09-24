package main

import (
	"bufio"
	"chat/internal/domain"
	"context"
	"encoding/json"
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
			//enterChat()
		case "exit":
			fmt.Println("Выход...")
			return
		default:
			fmt.Println("Неверный выбор, попробуйте снова.")
		}
	}
}

func createNewChat(cl *client) error {
	reader := bufio.NewReader(os.Stdin)
	for {
		fmt.Println("Введите айди пользователя, с которым вы бы хотели начать чат, или введите 'return' для выхода в предыдущее меню:")
		userID, _ := reader.ReadString('\n')
		userID = strings.TrimSpace(userID)

		if userID == "return" {
			return nil
		}

		if userID != "" {
			req, err := newCreateChatReq(userID)
			if err != nil {
				return fmt.Errorf("error creating new chat request: %v", err)
			}

			err = cl.WriteJSON(req)
			if err != nil {
				return err
			}

			var resp domain.Delivery
			err = cl.ReadJSON(&resp)
			if err != nil {
				return err
			}
			if resp.Type != domain.DeliveryTypeNewChat {
				return fmt.Errorf("invalid response type: %v", resp.Type)
			}

			fmt.Printf("Новый чат: %s\n\n", resp.Data.(string))
			return nil
		} else {
			fmt.Println("ID пользователя не может быть пустым, попробуйте снова.")
		}
	}
}

func newCreateChatReq(userId string) (domain.Request, error) {
	newChat := domain.NewChatRequest{
		UserIDs: []domain.ID{domain.ID(userId)},
	}
	data, err := json.Marshal(&newChat)
	if err != nil {
		return domain.Request{}, err
	}

	req := domain.Request{
		Type: domain.ReqTypeNewChat,
		Data: data,
	}

	return req, nil
}
