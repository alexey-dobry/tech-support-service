package bot

import (
	"log"

	"github.com/alexey-dobry/tech-support-platform/internal/services/manager_bot_service/internal/session"
	"gopkg.in/telebot.v4"
)

func (b *bot) HandleSendMsg() telebot.HandlerFunc {
	return func(c telebot.Context) error {
		senderID := c.Sender().ID

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
		err = c.Send(clientChat, c.Text())
		if err != nil {
			log.Println("Ошибка отправки клиенту:", err)
			return c.Send("Не удалось доставить сообщение клиенту.")
		}

		return c.Send("Ваше сообщение отправлено клиенту.")
	}
}
