package bot

import (
	middleware "github.com/alexey-dobry/tech-support-platform/internal/services/manager_bot_service/internal/bot/Middleware"
	"github.com/alexey-dobry/tech-support-platform/internal/services/manager_bot_service/internal/session"
	"gopkg.in/telebot.v4"
)

func (b *bot) HandleAuth() telebot.HandlerFunc {
	return func(c telebot.Context) error {
		args := c.Args()
		if len(args) < 2 {
			return c.Send("Используйте команду: /login <логин> <пароль>")
		}

		login := args[0]
		password := args[1]

		// Отправляем запрос к микросервису
		if middleware.Authenticate(login, password) {
			managerID := c.Sender().ID
			if session.IsAuthorized(managerID) {
				return c.Send("Вы уже вошли в аккаунт")
			} else {
				session.AddNewManager(managerID)
				return c.Send("Авторизация успешна! Теперь вы можете отвечать клиентам.")
			}
		}

		return c.Send("Неверный логин или пароль.")
	}
}
