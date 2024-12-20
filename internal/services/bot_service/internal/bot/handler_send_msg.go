package bot

import (
	"log"
	"strings"

	"github.com/alexey-dobry/tech-support-platform/internal/services/manager_bot_service/internal/session"
	"gopkg.in/telebot.v4"
)

func (b *bot) HandleSendMsg() telebot.HandlerFunc {
	return func(c telebot.Context) error {
		senderID := c.Sender().ID

		args := c.Args()
		if len(args) == 1 {
			return c.Send("Используйте команду: /reply <сообщение>")
		}

		// Проверяем, является ли отправитель менеджером
		managerClientID, err := session.GetActiveClientForManager(senderID)
		if err != nil {
			return c.Send("У вас нет активной сессии.")
		}

		if managerClientID == 0 {
			return c.Send("Нет активных клиентов для ответа.")
		}

		// Пересылаем сообщение клиенту
		clientChat := telebot.ChatID(managerClientID)
		msg := strings.Join(args[:], " ")
		_, err = b.client.Send(clientChat, msg)
		if err != nil {
			log.Println("Ошибка отправки клиенту:", err)
			return c.Send("Не удалось доставить сообщение клиенту.")
		}

		return c.Send("Ваше сообщение отправлено клиенту.")
	}
}
