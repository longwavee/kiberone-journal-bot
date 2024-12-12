package bot

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

func (b *Bot) HandleUpdate(update *tgbotapi.Update) {
	if update.Message != nil {
		if update.Message.Text != "" {
			b.HandleMessage(update)
		}
	}
	if update.CallbackQuery != nil {
		if update.CallbackQuery.Data != "" {
			b.HandleCallback(update)
		}
	}
}

func (b *Bot) HandleMessage(update *tgbotapi.Update) {
	t := update.Message.Text
	switch {
	case t == "/start":
		b.StartMessage(update)
	}
}

func (b *Bot) HandleCallback(update *tgbotapi.Update) {
	d := update.CallbackQuery.Data
	switch {
	case d == "":
	}
}
