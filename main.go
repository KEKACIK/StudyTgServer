package main

import (
	"StudyTgServer/api"
	"StudyTgServer/config"
	"StudyTgServer/telegram_bot"
)

func main() {
	cfg := config.Load()

	studyApi := api.NewStudyApiServer(cfg.StudyApiHost, cfg.StudyApiPort)
	bot, err := telegram_bot.NewBot(cfg.BotToken, studyApi)
	if err != nil {
		panic(err)
	}
	bot.Start()
}
