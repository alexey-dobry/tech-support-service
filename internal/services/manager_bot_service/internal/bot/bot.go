package bot

import (
	"log"
	"time"

	"github.com/alexey-dobry/tech-support-platform/internal/services/manager_bot_service/internal/config"
	"gopkg.in/telebot.v4"
)

var (
	managerID int64 = 549938415 // Замените на Telegram ID менеджера
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
		Token:  "8114672135:AAGX6_mJ_idUcVoxA9Vb_nDhY2BVlyOVD9U", // Замените на токен вашего бота
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
}

// Запуск бота
func (b *bot) Run() {
	log.Println("Бот запущен!")
	b.client.Start()
}
