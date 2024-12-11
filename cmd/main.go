package main

import (
	"context"
	"os"
	"os/signal"
	"syscall"

	"github.com/kukumber/pinyin_bot/cmd/config"
	"github.com/kukumber/pinyin_bot/internal/telegram"
)

func main() {
	// parse flags
	cfg, err := config.ParseFlags()
	if err != nil {
		panic(err)
	}

	// create context for graceful shutdown
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	// handle OS signals for graceful shutdown
	go func() {
		c := make(chan os.Signal, 1)
		signal.Notify(c, os.Interrupt, syscall.SIGTERM)
		<-c
		cancel()
	}()

	// get updates from Telegram
	bot, err := telegram.NewBot(cfg.APIKey)
	if err != nil {
		panic(err)
	}

	if err = bot.Start(ctx, &cfg); err != nil {
		panic(err)
	}
}
