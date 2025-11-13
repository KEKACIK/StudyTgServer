package main

import (
	"StudyTgServer/config"
	"StudyTgServer/internal/api"
	"StudyTgServer/internal/bot"
)

func main() {
	cfg := config.Load()
	studyApi := api.NewStudyApiServer(cfg.ApiHost, cfg.ApiPort)

	bot, err := bot.NewBot(cfg.BotToken, studyApi)
	if err != nil {
		panic(err)
	}

	bot.Start()
}
