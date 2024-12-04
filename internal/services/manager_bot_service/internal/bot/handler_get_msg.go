package bot

import (
	"fmt"
	"log"

	"gopkg.in/telebot.v4"
)

func (b *bot) HandleGetMsg() telebot.HandlerFunc {
	return func(c telebot.Context) error {
		user := c.Sender()

		if managerID == 0 {
			return c.Send("Менеджер вне зоны доступа")
		}

		managerChat := telebot.ChatID(managerID)
		forwardMessage := fmt.Sprintf("Новое сообщение от клиента:\n👤 Клиент: @ %s (%s %s)\nID клиента: %d\n💬 Сообщение: %s", user.Username,
			user.FirstName, user.LastName, user.ID, c.Message().Text)

		if user.ID != managerID {
			_, err := b.client.Send(managerChat, forwardMessage)
			if err != nil {
				log.Printf("Error sending message from client to manager: %s", err)
				return c.Send("Возникла проблема при отправке сообщения, попробуйте отправить поже")
			}
		}
		return nil
	}

}
