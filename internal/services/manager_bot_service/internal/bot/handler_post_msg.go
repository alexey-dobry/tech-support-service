package bot

import (
	"log"
	"strconv"
	"strings"

	"github.com/alexey-dobry/tech-support-platform/internal/services/manager_bot_service/internal/session"
	"gopkg.in/telebot.v4"
)

func (b *bot) HandleSendMsg() telebot.HandlerFunc {
	return func(c telebot.Context) error {
		args := c.Args()

		if len(args) < 2 {
			log.Print("Для отправки ответа введите комманду вида: /reply <user_id> <message>")
		}

		clientID, err := strconv.ParseInt(args[0], 10, 64)
		if err != nil {
			return c.Send("ID должен быть числом")
		}
		message := strings.Join(args[1:], " ")

		if !session.IsAuthorized(c.Sender().ID) {
			return c.Send("У вас недостаточно прав доступа")
		}

		// Получить чат по ID клиента
		clientChat := telebot.ChatID(clientID)

		// Ответить клиенту
		_, err = b.client.Send(clientChat, "Ответ от менеджера: "+message)
		if err != nil {
			log.Printf("Error sending message from manager to client: %s", err)
			return c.Send("Ошибка при отправке сообщения клиенту")
		}
		return err
	}
}
