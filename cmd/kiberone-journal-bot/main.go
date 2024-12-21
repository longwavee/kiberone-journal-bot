package main

import (
	"log"

	"github.com/longwavee/kiberone-journal-bot/internal/bot"
	"github.com/longwavee/kiberone-journal-bot/internal/config"
	"github.com/longwavee/kiberone-journal-bot/internal/pkg/logger/zerolog"
)

func main() {
	config := config.MustLoad()

	// storage, err := postgres.New(config.Storage)
	// if err != nil {
	// 	log.Fatalf("failed to create storage: %v", err)
	// }

	logger, err := zerolog.New()
	if err != nil {
		log.Fatalf("failed to create logger: %v", err)
	}

	b, err := bot.New(config.Bot, nil, logger)
	if err != nil {
		log.Fatalf("failed to create bot instance: %v", err)
	}

	if err := b.Run(); err != nil {
		log.Fatalf("error running the bot: %v", err)
	}
}
