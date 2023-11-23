package telegram

import (
	"log"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"

	"github.com/kukumber/pinyin_bot/cmd/config"
	"github.com/kukumber/pinyin_bot/internal/converter"
)

// GetUpdates gets updates from Telegram
func GetUpdates(cfg *config.Config) error {
	// create a new bot
	bot, err := tgbotapi.NewBotAPI(cfg.APIKey)
	if err != nil {
		return err
	}
	bot.Debug = true

	log.Printf("Authorized on account %s", bot.Self.UserName)

	// get updates
	u := tgbotapi.NewUpdate(cfg.InitOffset)
	u.Timeout = 60

	updates := bot.GetUpdatesChan(u)

	for update := range updates {
		// ignore any non-Message updates
		if update.Message == nil {
			continue
		}

		log.Printf("[%s] %s", update.Message.From.UserName, update.Message.Text)

		// ignore any non-command Messages
		if !update.Message.IsCommand() {
			continue
		}

		var result string

		// process commands
		switch update.Message.Command() {
		case "py":
			result = converter.ConvertToPinyin(update.Message.CommandArguments())
		case "pld":
			result = converter.ConvertToPallady(update.Message.CommandArguments())
		default:
			log.Printf("unknown command: %s", update.Message.Command())
		}

		// send the result
		msg := tgbotapi.NewMessage(update.Message.Chat.ID, result)
		msg.ReplyToMessageID = update.Message.MessageID
		bot.Send(msg)
	}

	return nil
}
