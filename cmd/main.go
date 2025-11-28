package main

import (
	"StudyTgServer/config"
	"StudyTgServer/internal/api"
	"StudyTgServer/internal/bot"
	"fmt"
)

func main() {
	fmt.Println()

	cfg := config.Load()
	studyApi := api.NewStudyApiServer(cfg.ApiHost, cfg.ApiPort, cfg.ApiToken)

	bot, err := bot.NewBot(cfg.BotToken, studyApi)
	if err != nil {
		panic(err)
	}

	bot.Start()
}
