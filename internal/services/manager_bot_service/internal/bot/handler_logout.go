package bot

import (
	"github.com/alexey-dobry/tech-support-platform/internal/services/manager_bot_service/internal/session"
	"gopkg.in/telebot.v4"
)

func (b *bot) HandleLogut() telebot.HandlerFunc {
	return func(c telebot.Context) error {
		args := c.Args()

		if len(args) != 0 {
			return c.Send("Данная команда не пришимает никаких дополнительных значений")
		}

		senderId := c.Sender().ID

		if session.IsAuthorized(senderId) {
			session.DeauthorizeManager(senderId)
		} else {
			return c.Send("Вы не вошли в аккаунт")
		}
		return nil
	}
}
