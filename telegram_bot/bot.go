package telegram_bot

import (
	"StudyTgServer/api"
	"time"

	"gopkg.in/telebot.v3"
)

type States string

const (
	StudentCreateName   States = "student_create_name"
	StudentCreateSex    States = "student_create_sex"
	StudentCreateAge    States = "student_create_age"
	StudentCreateCourse States = "student_create_course"

	StudentGet States = "student_get"
)

type Bot struct {
	*telebot.Bot
	states map[int64]States
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
		api:    study_api,
	}

	bot.RegisterHandlers(study_api)

	return &bot, nil
}
