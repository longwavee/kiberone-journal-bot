package bot

import tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"

func (b *Bot) StartMessage(update *tgbotapi.Update) {
	msg := tgbotapi.NewMessage(update.Message.Chat.ID, "*Hello World!*")
	msg.ParseMode = "Markdown"
	msg.DisableNotification = true
	msg.ReplyMarkup = nil
	b.client.Send(msg)
}
