package main

import (
	"StudyTgServer/api"
	"StudyTgServer/config"
	"StudyTgServer/telegram_bot"
	"context"
	"os"
	"os/signal"
)

func main() {
	ctx, cancel := signal.NotifyContext(context.Background(), os.Interrupt)
	cfg := config.Load()
	defer cancel()

	studyApi := api.NewStudyApiServer(cfg.StudyApiHost, cfg.StudyApiPort)
	tgBot, err := telegram_bot.NewTgBot(cfg.BotToken, studyApi)
	if err != nil {
		panic(err)
	}
	tgBot.Start(ctx)
}
