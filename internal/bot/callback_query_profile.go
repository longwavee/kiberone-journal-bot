package bot

import (
	"fmt"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

func (b *Bot) CallbackShiftData(update *tgbotapi.Update) {
	fmt.Println("нажатие на кнопку")
}
