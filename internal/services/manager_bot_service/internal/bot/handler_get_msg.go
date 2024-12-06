package bot

import (
	"fmt"
	"log"

	"github.com/alexey-dobry/tech-support-platform/internal/services/manager_bot_service/internal/session"
	"gopkg.in/telebot.v4"
)

var managerID int64 = 549938415

func (b *bot) HandleGetMsg() telebot.HandlerFunc {
	return func(c telebot.Context) error {
		user := c.Sender()

		managerChat := telebot.ChatID(managerID)
		forwardMessage := fmt.Sprintf("Новое сообщение от клиента:\n👤 Клиент: @ %s (%s %s)\nID клиента: %d\n💬 Сообщение: %s", user.Username,
			user.FirstName, user.LastName, user.ID, c.Message().Text)

		if session.IsAuthorized(managerID) {
			_, err := b.client.Send(managerChat, forwardMessage)
			if err != nil {
				log.Printf("Error sending message from client to manager: %s", err)
				return c.Send("Возникла проблема при отправке сообщения, попробуйте отправить поже")
			}
			return nil
		} else {
			return nil
		}
	}

}
