package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func enterChat(cl *client) error {
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
	}
}
