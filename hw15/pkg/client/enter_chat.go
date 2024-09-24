package main

import (
	"bufio"
	"chat/internal/domain"
	"context"
	"fmt"
	"log"
	"os"
	"strings"
	"sync"
	"time"
)

func enterChat(ctx context.Context, cl *client) error {
	reader := bufio.NewReader(os.Stdin)
	for {
		fmt.Println("Введите айди чата для начала общения или введите return для выхода в предыдущее меню.:")
		chatID, _ := reader.ReadString('\n')
		chatID = strings.TrimSpace(chatID)

		if chatID == "return" {
			return nil
		}

		if chatID == "" {
			fmt.Println("ID чата не может быть пустым, попробуйте снова.")
			continue
		}

		fmt.Println("Вводите сообщения для отправки:")

		wg := sync.WaitGroup{}
		wg.Add(1)
		go func() {
			defer wg.Done()
			for {
				select {
				case <-ctx.Done():
					return
				default:
					var msg domain.Delivery
					err := cl.ReadJSON(&msg)
					if err != nil {
						log.Fatalf("error reading message: %v", err)
						return
					}
					if msg.Type != domain.DeliveryTypeNewMsg {
						log.Fatalf("invalid delivery type: %v", msg.Type)
						return
					}
					data := msg.Data.(map[string]any)

					date := data["t_date"]
					t, err := time.Parse(time.RFC3339Nano, date.(string))
					if err != nil {
						log.Fatalf("error parsing date: %v", err)
						return
					}

					fmt.Printf("%s %s:\n", data["from_id"], t.Format(time.DateTime))
					fmt.Println(data["body"])
				}
			}
		}()
		wg.Wait()
	}
}
