package telegram

import (
	"log/slog"

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

	slog.Info("authorized on account", "account", bot.Self.UserName)

	// get updates
	u := tgbotapi.NewUpdate(cfg.InitOffset)
	u.Timeout = 60

	updates := bot.GetUpdatesChan(u)

	for update := range updates {
		// ignore any non-Message updates
		if update.Message == nil {
			continue
		}

		slog.Info("message received", "user", update.Message.From.UserName, "text", update.Message.Text)

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
			slog.Warn("unknown command", "command", update.Message.Command())
		}

		// check that result is not empty
		if result == "" {
			result = "Command not recognized or processing failed."
		}

		// send the result
		msg := tgbotapi.NewMessage(update.Message.Chat.ID, result)
		msg.ReplyToMessageID = update.Message.MessageID

		if _, err := bot.Send(msg); err != nil {
			slog.Error("Failed to send message", "error", err)
		}
	}

	return nil
}
