package bot

import (
	"StudyTgServer/internal/api"

	"gopkg.in/telebot.v3"
)

func selectSexKeyboard(key string) *telebot.ReplyMarkup {
	keyboard := &telebot.ReplyMarkup{
		InlineKeyboard: [][]telebot.InlineButton{
			{
				telebot.InlineButton{
					Text:   "Мужской",
					Unique: key,
					Data:   api.StudyStudentSexMan,
				},
				telebot.InlineButton{
					Text:   "Женский",
					Unique: key,
					Data:   api.StudyStudentSexWoman,
				},
			},
		},
	}
	return keyboard
}

func getUpdateFieldKeyboard() *telebot.ReplyMarkup {
	keyboard := &telebot.ReplyMarkup{
		InlineKeyboard: [][]telebot.InlineButton{
			{
				telebot.InlineButton{
					Text:   "Имя",
					Unique: "get_update",
					Data:   "name",
				},
			},
			{
				telebot.InlineButton{
					Text:   "Пол",
					Unique: "get_update",
					Data:   "sex",
				},
			},
			{
				telebot.InlineButton{
					Text:   "Возраст",
					Unique: "get_update",
					Data:   "age",
				},
			},
			{
				telebot.InlineButton{
					Text:   "Курс",
					Unique: "get_update",
					Data:   "course",
				},
			},
		},
	}
	return keyboard
}

func deleteSuccessKeyboard() *telebot.ReplyMarkup {
	keyboard := &telebot.ReplyMarkup{
		InlineKeyboard: [][]telebot.InlineButton{
			{
				telebot.InlineButton{
					Text:   "Удалить",
					Unique: "delete_yes",
				},
				telebot.InlineButton{
					Text:   "Отмена",
					Unique: "delete_no",
				},
			},
		},
	}
	return keyboard
}
