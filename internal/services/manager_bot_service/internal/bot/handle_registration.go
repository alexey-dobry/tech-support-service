package bot

import (
	middleware "github.com/alexey-dobry/tech-support-platform/internal/services/manager_bot_service/internal/bot/Middleware"
	"gopkg.in/telebot.v4"
)

func (b *bot) HandleRegistration() telebot.HandlerFunc {
	return func(c telebot.Context) error {
		args := c.Args()

		if len(args) < 3 {
			return c.Send("Введите команду в формате /register <ключ> <логин> <пароль>")
		}

		hashKey := args[0]
		login := args[1]
		password := args[2]

		err := middleware.Register(hashKey, login, password)
		if err != nil {
			return c.Send("Неудалось зарегестрировать нового менеджера. Причина: %s", err)
		}

		return nil
	}
}
