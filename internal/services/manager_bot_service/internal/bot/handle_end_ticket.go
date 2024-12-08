package bot

import (
	"log"

	"github.com/alexey-dobry/tech-support-platform/internal/services/manager_bot_service/internal/session"
	"gopkg.in/telebot.v4"
)

func (b *bot) handleEndTicket() telebot.HandlerFunc {
	return func(c telebot.Context) error {
		sender := c.Sender().ID

		if session.IsAuthorized(sender) {
			session.FreeManager(sender)
			log.Printf("Succesfully closed ticket for user %d", sender)
			return c.Send("Тикет успешно закрыт")
		} else {
			log.Printf("Permission denied for command /end user %d", sender)
			return c.Send("Вы не авторизованы для выполнения данной команды")
		}
	}
}
