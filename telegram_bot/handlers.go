package telegram_bot

import (
	"StudyTgServer/api"
	"StudyTgServer/utils"
	"fmt"

	"gopkg.in/telebot.v3"
)

func (b *Bot) RegisterHandlers(api *api.StudyApiServer) {
	b.api = api

	b.Handle("/start", b.startHandler)
	b.Handle("/create", b.createHandler)
	b.Handle("/get", b.getHandler)
	b.Handle("/get_all", b.getAllHandler)
	b.Handle("/update", b.updateHandler)
	b.Handle("/delete", b.deleteHandler)
	// tb.client.RegisterHandler(bot.HandlerTypeMessageText, "/create", bot.MatchTypeExact, tb.createHandler)
	// tb.client.RegisterHandler(bot.HandlerTypeMessageText, "/get", bot.MatchTypeExact, tb.getHandler)
	// tb.client.RegisterHandler(bot.HandlerTypeMessageText, "/get_all", bot.MatchTypeExact, tb.getAllHandler)
	// tb.client.RegisterHandler(bot.HandlerTypeMessageText, "/update", bot.MatchTypeExact, tb.updateHandler)
	// tb.client.RegisterHandler(bot.HandlerTypeMessageText, "/delete", bot.MatchTypeExact, tb.deleteHandler)
}

// COMMAND START

func (b *Bot) startHandler(c telebot.Context) error {
	return c.Send(
		utils.MultiLine(
			"Добро пожаловать в <b>StudyTgBot</b>",
			"",
			"Вот список комманд для использования бота",
			"/create - Добавить нового студента",
			"/get - Получить студента по ID",
			"/get_all - Получить всех студентов",
			"/update - Обновить студента по ID",
			"/delete - Удалить студента по ID",
		),
		&telebot.SendOptions{
			DisableWebPagePreview: false,
			ParseMode:             telebot.ModeHTML,
		},
	)
}

// COMMAND CREATE

func (tb *Bot) createHandler(c telebot.Context) error {
	return nil
}

// COMMAND GET

func (tb *Bot) getHandler(c telebot.Context) error {
	return nil
}

// COMMAND GET_ALL

func (tb *Bot) getAllHandler(c telebot.Context) error {
	students, err := tb.api.GetAll()
	if err != nil {
		return c.Send(
			"Произошла ошибка, обратитесь в Тех. Поддержку",
			&telebot.SendOptions{
				DisableWebPagePreview: false,
				ParseMode:             telebot.ModeHTML,
			},
		)
	}

	text := "Список студентов:\n"
	for _, student := range students {
		sex_ru := ""
		switch student.Sex {
		case api.StudyStudentSexMan:
			sex_ru = "Мужской"
		case api.StudyStudentSexWoman:
			sex_ru = "Женский"
		default:
			sex_ru = "Не определен"
		}
		text = utils.MultiLine(
			text,
			fmt.Sprintf(
				"ID: <b><i>%d</i></b>, Имя: <b><i>%s</i></b>, Пол: <b><i>%s</i></b>, Возраст: <b><i>%d</i></b>, Курс: <b><i>%d</i></b>",
				student.ID, student.Name, sex_ru, student.Age, student.Course,
			),
		)
	}

	return c.Send(
		text,
		&telebot.SendOptions{
			DisableWebPagePreview: false,
			ParseMode:             telebot.ModeHTML,
		},
	)
}

// COMMAND UPDATE

func (tb *Bot) updateHandler(c telebot.Context) error {
	return nil
}

// COMMAND DELETE

func (tb *Bot) deleteHandler(c telebot.Context) error {
	return nil
}
