package bot

import (
	"StudyTgServer/api"

	"gopkg.in/telebot.v3"
)

func createSexKeyboard() *telebot.ReplyMarkup {
	keyboard := &telebot.ReplyMarkup{
		InlineKeyboard: [][]telebot.InlineButton{
			{
				telebot.InlineButton{
					Text:   "Мужской",
					Unique: "create_sex",
					Data:   api.StudyStudentSexMan,
				},
				telebot.InlineButton{
					Text:   "Женский",
					Unique: "create_sex",
					Data:   api.StudyStudentSexWoman,
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
