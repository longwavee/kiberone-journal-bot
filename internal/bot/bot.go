package bot

import (
	"net/http"

	"github.com/longwavee/kiberone-journal-bot/internal/config"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

type Storager interface {
}

type Logger interface {
	Info(msg string, fields ...any)
	Warn(msg string, fields ...any)
	Error(msg string, fields ...any)
}

type Bot struct {
	config *config.Bot
	client *tgbotapi.BotAPI

	storage Storager
	logger  Logger
}

func New(config *config.Bot, client *tgbotapi.BotAPI, storage Storager, logger Logger) *Bot {
	return &Bot{
		config: config,
		client: client,

		storage: storage,
		logger:  logger,
	}
}

func (b *Bot) Run() error {
	pattern := "/" + b.config.Token

	wh, err := tgbotapi.NewWebhook(b.config.WebhookAddr + pattern)
	if err != nil {
		return err
	}

	_, err = b.client.Request(wh)
	if err != nil {
		return err
	}

	updates := b.client.ListenForWebhook(pattern)
	go http.ListenAndServe(b.config.HostAddr, nil)

	for update := range updates {
		go b.HandleUpdate(&update)
	}

	return nil
}
