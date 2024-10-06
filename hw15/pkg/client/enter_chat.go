package main

import (
	"bufio"
	"chat/internal/domain"
	"context"
	"encoding/json"
	"fmt"
	"log"
	"os"
	"strings"
	"time"
)

func enterChat(ctx context.Context, cl *client) error {
	reader := bufio.NewReader(os.Stdin)
	for {
		fmt.Println("Введите айди чата для начала общения или введите return для выхода в предыдущее меню:")
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

		go func() {
			for {
				select {
				case <-ctx.Done():
					return
				default:
					var msg domain.Delivery
					err := cl.ReadJSON(&msg)
					if err != nil {
						log.Printf("error reading message: %v", err)
						return
					}

					if msg.Type != domain.DeliveryTypeNewMsg {
						log.Printf("error reading message: %v", err)
						return
					}

					msgData := msg.Data.(map[string]any)

					t, err := time.Parse(time.RFC3339, msgData["t_date"].(string))
					if err != nil {
						log.Printf("error parsing time: %v", err)
					}

					fmt.Printf(
						"%v %s: \n%s",
						msgData["from_id"],
						t.Format(time.DateTime),
						msgData["body"],
					)
				}
			}
		}()

		for {
			select {
			case <-ctx.Done():
				err := cl.CloseChat()
				if err != nil {
					return err
				}

				return nil
			default:
				body, err := reader.ReadString('\n')
				if err != nil {
					return fmt.Errorf("error reading message from console: %v", err)
				}

				req := domain.Request{
					Type: domain.ReqTypeNewMsg,
				}

				msg := domain.MessageChatRequest{
					Msg:  body,
					Type: domain.MsgTypeAdd,
					ChID: domain.ID(chatID),
				}
				data, err := json.Marshal(msg)
				if err != nil {
					return fmt.Errorf("error marshalling message: %v", err)
				}

				req.Data = data

				if err = cl.WriteJSON(req); err != nil {
					return fmt.Errorf("error writing message: %v", err)
				}
			}
		}
	}
}
