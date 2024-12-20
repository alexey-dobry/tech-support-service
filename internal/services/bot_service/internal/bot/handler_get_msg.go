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
		clientID := c.Sender().ID

		// Проверяем, есть ли у клиента назначенный менеджер
		managerID, err := session.GetAssignedManager(clientID)
		if err != nil {
			return c.Send("Произошла ошибка. Попробуйте позже.")
		}

		if managerID == 0 {
			// Менеджер не назначен, назначаем нового
			managerID, err = session.AssignClientToManager(clientID)
			if err != nil {
				return c.Send("Извините, сейчас нет свободных менеджеров. Пожалуйста, подождите.")
			}
		}

		// Отправляем сообщение назначенному менеджеру
		managerChat := telebot.ChatID(managerID)
		msg := fmt.Sprintf("Новое сообщение от клиента %d: %s", clientID, c.Message().Text)
		_, err = b.client.Send(managerChat, msg)
		if err != nil {
			log.Println("Ошибка отправки сообщения менеджеру:", err)
			return c.Send("Не удалось доставить сообщение менеджеру. Попробуйте позже.")
		}

		return c.Send("Ваше сообщение отправлено менеджеру.")
	}

}
