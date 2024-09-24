package main

import (
	"bufio"
	"chat/internal/domain"
	"encoding/json"
	"fmt"
	"os"
	"strings"
)

func createNewChat(cl *client) error {
	reader := bufio.NewReader(os.Stdin)
	for {
		fmt.Println("Введите айди пользователя, с которым вы бы хотели начать чат, или введите 'return' для выхода в предыдущее меню:")
		userID, _ := reader.ReadString('\n')
		userID = strings.TrimSpace(userID)

		if userID == "return" {
			return nil
		}

		if userID == "" {
			fmt.Println("ID пользователя не может быть пустым, попробуйте снова.")
			continue
		}

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
