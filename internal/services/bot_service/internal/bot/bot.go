package bot

import (
	"log"
	"time"

	"github.com/alexey-dobry/tech-support-platform/internal/services/manager_bot_service/internal/config"
	"gopkg.in/telebot.v4"
)

type Bot interface {
	Run()
}

type bot struct {
	client *telebot.Bot
}

// Создание инстанса бота
func New(cfg *config.Config) Bot {
	var b bot
	client, err := telebot.NewBot(telebot.Settings{
		Token:  8114672135:AAE_3GlOCpu_huPnuNbS3B8ooD0ogLCvUm8,
		Poller: &telebot.LongPoller{Timeout: 10 * time.Second},
	})

	if err != nil {
		log.Fatalf("error %s", err)
	}

	b.client = client

	log.Print("Initiating handlers...")

	b.initHandlers()

	log.Print("Handlers intiated")

	return &b
}

// Инициализация функций бота
func (b *bot) initHandlers() {
	// Клиенты отправляют сообщения
	b.client.Handle(telebot.OnText, b.HandleGetMsg())

	// Менеджер отвечает клиенту
	b.client.Handle("/reply", b.HandleSendMsg())

	// Вход в систему(для менеджера)
	b.client.Handle("/login", b.HandleAuth())

	// Выход из системы(для менеджера)
	b.client.Handle("/logout", b.HandleLogut())

	b.client.Handle("/end", b.handleEndTicket())
}

// Запуск бота
func (b *bot) Run() {
	log.Println("Бот запущен!")
	b.client.Start()
}
