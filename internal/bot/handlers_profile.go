package bot

import (
	"errors"
	"fmt"
	"os"
	"path/filepath"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

func (b *Bot) HandleCallbackProfile(update *tgbotapi.Update) {
	d := update.CallbackQuery.Data
	switch {
	case d == "callback_shift_data":
		b.CallbackShiftData(update)
	}
}

func (b *Bot) newMessageProfilePhoto(imgProfile tgbotapi.UserProfilePhotos, update *tgbotapi.Update) tgbotapi.PhotoConfig {
	return tgbotapi.NewPhoto(update.Message.Chat.ID, tgbotapi.FileURL(imgProfile.Photos[0][0].FileID))
}

func (b *Bot) thirdPartyProfilePhotosNewMessage(update *tgbotapi.Update) (tgbotapi.PhotoConfig, error) {
	assetPath := filepath.Join("assets", "photeProfile.jpg")
	if _, err := os.Stat(assetPath); os.IsNotExist(err) {
		return tgbotapi.PhotoConfig{}, errors.New("file not found")
	}
	imgMessage := tgbotapi.NewPhoto(update.Message.Chat.ID, tgbotapi.FilePath(assetPath))
	return imgMessage, nil
}

func (b *Bot) HandleProfile(update *tgbotapi.Update) {
	worker := b.storage.Worker(update.Message.From.ID)
	if worker == nil {
		message := tgbotapi.NewMessage(update.Message.Chat.ID, "вы не зарегистрированы")
		b.client.Send(message)
		return
	}

	text := fmt.Sprintf(
		"%s %s \nusername: %s\nчасы тьютера %b\nчасы ассистента %b\n",
		worker.FirstName,
		worker.LastName,
		update.Message.From.UserName,
		worker.TutorWork,
		worker.AssisWork,
	)

	btn := tgbotapi.NewInlineKeyboardButtonData("рассписания смен", "callback_shift_data")
	inlineKeyboard := tgbotapi.NewInlineKeyboardMarkup(
		tgbotapi.NewInlineKeyboardRow(btn),
	)

	imgMessage := tgbotapi.PhotoConfig{}
	imgProfile, err := b.client.GetUserProfilePhotos(tgbotapi.UserProfilePhotosConfig{UserID: update.Message.From.ID})
	if err == nil && len(imgProfile.Photos) > 0 && len(imgProfile.Photos[0]) > 0 {
		imgMessage = b.newMessageProfilePhoto(imgProfile, update)
	} else {
		imgMessage, err = b.thirdPartyProfilePhotosNewMessage(update)
		if err != nil {
			message := tgbotapi.NewMessage(update.Message.Chat.ID, text)
			message.ParseMode = "Markdown"
			message.ReplyMarkup = inlineKeyboard
			b.client.Send(message)
			return
		}
	}

	imgMessage.Caption = text
	imgMessage.ParseMode = "Markdown"
	imgMessage.ReplyMarkup = inlineKeyboard

	b.client.Send(imgMessage)
}
