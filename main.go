package main

import (
	"StudyTgServer/api"
	"StudyTgServer/bot"
	"StudyTgServer/config"
)

func main() {
	cfg := config.Load()

	studyApi := api.NewStudyApiServer(cfg.StudyApiHost, cfg.StudyApiPort)
	bot, err := bot.NewBot(cfg.BotToken, studyApi)
	if err != nil {
		panic(err)
	}
	bot.Start()
}
