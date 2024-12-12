package main

import (
	"log"

	"github.com/longwavee/kiberone-journal-bot/internal/bot"
	"github.com/longwavee/kiberone-journal-bot/internal/config"
	"github.com/longwavee/kiberone-journal-bot/internal/pkg/logger/zerolog"
	"github.com/longwavee/kiberone-journal-bot/internal/pkg/storage/postgres"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

func main() {
	config := config.MustLoad()

	client, err := tgbotapi.NewBotAPI(config.Bot.Token)
	if err != nil {
		log.Fatal(err)
	}

	storage, err := postgres.New(config.Storage)
	if err != nil {
		log.Fatal(err)
	}

	logger, err := zerolog.New()
	if err != nil {
		log.Fatal(err)
	}

	b := bot.New(config.Bot, client, storage, logger)
	b.Run("/" + config.Bot.Token)
}
