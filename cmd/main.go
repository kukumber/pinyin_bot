package main

import (
	"github.com/kukumber/pinyin_bot/cmd/config"
	"github.com/kukumber/pinyin_bot/internal/telegram"
)

func main() {
	// parse flags
	cfg, err := config.ParseFlags()
	if err != nil {
		panic(err)
	}

	// get updates from Telegram
	err = telegram.GetUpdates(&cfg)
	if err != nil {
		panic(err)
	}
}
