package main

import (
	"context"
	"log/slog"
	"os"
	"os/signal"
	"syscall"

	"github.com/kukumber/pinyin_bot/internal/config"
	"github.com/kukumber/pinyin_bot/internal/telegram"
)

func main() {
	// parse flags
	cfg, err := config.ParseFlags()
	if err != nil {
		slog.Error("invalid configuration", "error", err)
		os.Exit(1)
	}

	// create context for graceful shutdown on OS signals
	ctx, stop := signal.NotifyContext(context.Background(), os.Interrupt, syscall.SIGTERM)
	defer stop()

	// get updates from Telegram
	bot, err := telegram.NewBot(cfg)
	if err != nil {
		slog.Error("failed to create bot", "error", err)
		os.Exit(1)
	}

	if err = bot.Start(ctx, cfg); err != nil {
		slog.Error("bot stopped with error", "error", err)
		os.Exit(1)
	}
}
