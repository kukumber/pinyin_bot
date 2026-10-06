package telegram

import (
	"context"
	"fmt"
	"log/slog"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"

	"github.com/kukumber/pinyin_bot/internal/config"
	"github.com/kukumber/pinyin_bot/internal/converter"
)

type Bot struct {
	api *tgbotapi.BotAPI
}

var commandHandlers = map[string]func(string) string{
	"py":  converter.ConvertToPinyin,
	"pld": converter.ConvertToPallady,
}

// NewBot initializes a new bot
func NewBot(cfg config.Config) (*Bot, error) {
	api, err := tgbotapi.NewBotAPI(cfg.APIKey)
	if err != nil {
		return nil, fmt.Errorf("failed to create bot: %w", err)
	}
	api.Debug = cfg.Debug
	slog.Info("bot authorized", "username", api.Self.UserName)

	return &Bot{api: api}, nil
}

// Start listens for updates and processes them until ctx is canceled
func (b *Bot) Start(ctx context.Context, cfg config.Config) error {
	u := tgbotapi.NewUpdate(cfg.InitOffset)
	u.Timeout = cfg.UpdateInterval

	updates := b.api.GetUpdatesChan(u)
	defer b.api.StopReceivingUpdates()

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
	return err
}
