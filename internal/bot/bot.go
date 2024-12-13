package bot

import (
	"net/http"

	"github.com/longwavee/kiberone-journal-bot/internal/config"
	"github.com/longwavee/kiberone-journal-bot/internal/model"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

type Storager interface {
	Worker(id int64) *model.Worker
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

func New(config *config.Bot, storage Storager, logger Logger) (*Bot, error) {
	client, err := tgbotapi.NewBotAPI(config.Token)
	if err != nil {
		return nil, err
	}

	return &Bot{
		config: config,

		client: client,

		storage: storage,
		logger:  logger,
	}, err
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
		go func() {
			b.HandleUpdate(&update)
		}()
	}

	return nil
}
