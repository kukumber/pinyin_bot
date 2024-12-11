package telegram

import (
	"context"
	"fmt"
	"log/slog"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"

	"github.com/kukumber/pinyin_bot/cmd/config"
	"github.com/kukumber/pinyin_bot/internal/converter"
)

const botTimeout = 60

type Bot struct {
	api *tgbotapi.BotAPI
}

// NewBot initializes a new bot
func NewBot(apiKey string) (*Bot, error) {
	bot, err := tgbotapi.NewBotAPI(apiKey)
	if err != nil {
		return nil, fmt.Errorf("failed to create bot: %w", err)
	}
	bot.Debug = true
	slog.Info("bot authorized", "username", bot.Self.UserName)

	return &Bot{api: bot}, nil
}

// Start listens for updates and processes them
func (b *Bot) Start(ctx context.Context, cfg *config.Config) error {
	u := tgbotapi.NewUpdate(cfg.InitOffset)
	u.Timeout = botTimeout

	updates := b.api.GetUpdatesChan(u)

	for {
		select {
		case <-ctx.Done():
			return nil
		case update := <-updates:
			if err := b.processUpdate(update); err != nil {
				slog.Error("failed to process update", "error", err)
			}
		}
	}
}

// processUpdate processes a single Telegram update
func (b *Bot) processUpdate(update tgbotapi.Update) error {
	if update.Message == nil {
		return nil
	}

	slog.Info("message received", "user", update.Message.From.UserName, "text", update.Message.Text)

	if !update.Message.IsCommand() {
		return nil
	}

	commandHandlers := map[string]func(string) string{
		"py":  converter.ConvertToPinyin,
		"pld": converter.ConvertToPallady,
	}

	handler, exists := commandHandlers[update.Message.Command()]
	if !exists {
		slog.Warn("unknown command", "command", update.Message.Command())
		return b.reply(update, "unknown command.")
	}

	result := handler(update.Message.CommandArguments())
	if result == "" {
		result = "command processing failed"
	}

	return b.reply(update, result)
}

// reply sends a reply to the user
func (b *Bot) reply(update tgbotapi.Update, text string) error {
	msg := tgbotapi.NewMessage(update.Message.Chat.ID, text)
	msg.ReplyToMessageID = update.Message.MessageID

	_, err := b.api.Send(msg)
	if err != nil {
		slog.Error("failed to send message", "chat_id", update.Message.Chat.ID, "error", err)
	}

	return err
}
