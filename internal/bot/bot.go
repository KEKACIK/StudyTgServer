package bot

import (
	"StudyTgServer/internal/api"
	"fmt"
	"time"

	"gopkg.in/telebot.v3"
)

type States string

type Bot struct {
	*telebot.Bot
	states map[int64]States
	data   map[int64]api.StudyApiStudent
	api    *api.StudyApiServer
}

func NewBot(token string, study_api *api.StudyApiServer) (*Bot, error) {
	bot_settings := telebot.Settings{
		Token:  token,
		Poller: &telebot.LongPoller{Timeout: 10 * time.Second},
	}
	b, err := telebot.NewBot(bot_settings)
	if err != nil {
		return nil, err
	}
	bot := Bot{
		Bot:    b,
		states: make(map[int64]States),
		data:   make(map[int64]api.StudyApiStudent),
		api:    study_api,
	}
	bot.SetCommands()
	bot.RegisterHandlers()

	return &bot, nil
}

func (b *Bot) Start() {
	fmt.Println("Запуск бота")
	b.Bot.Start()
}
