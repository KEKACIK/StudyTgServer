package telegram_bot

import (
	"StudyTgServer/api"
	"StudyTgServer/utils"
	"context"
	"fmt"

	"github.com/go-telegram/bot"
	"github.com/go-telegram/bot/models"
)

// COMMAND START

func (tb *TgBot) startHandler(ctx context.Context, _ *bot.Bot, update *models.Update) {
	tb.client.SendMessage(ctx, &bot.SendMessageParams{
		ChatID: update.Message.Chat.ID,
		Text: utils.MultiLine(
			"Добро пожаловать в <b>StudyTgBot</b>",
			"",
			"Вот список комманд для использования бота",
			"/create - Добавить нового студента",
			"/get - Получить студента по ID",
			"/get_all - Получить всех студентов",
			"/update - Обновить студента по ID",
			"/delete - Удалить студента по ID",
		),
		ParseMode: models.ParseModeHTML,
	})
}

// COMMAND CREATE

func (tb *TgBot) createHandler(ctx context.Context, _ *bot.Bot, update *models.Update) {

}

// COMMAND GET

func (tb *TgBot) getHandler(ctx context.Context, _ *bot.Bot, update *models.Update) {

}

// COMMAND GET_ALL

func (tb *TgBot) getAllHandler(ctx context.Context, _ *bot.Bot, update *models.Update) {
	students, err := tb.api.GetAll()
	fmt.Println(students)
	fmt.Println(err)
	if err != nil {
		tb.client.SendMessage(ctx, &bot.SendMessageParams{
			ChatID:    update.Message.Chat.ID,
			Text:      "Произошла ошибка, обратитесь в Тех. Поддержку",
			ParseMode: models.ParseModeHTML,
		})
		return
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

	tb.client.SendMessage(ctx, &bot.SendMessageParams{
		ChatID:    update.Message.Chat.ID,
		Text:      text,
		ParseMode: models.ParseModeHTML,
	})
}

// COMMAND UPDATE

func (tb *TgBot) updateHandler(ctx context.Context, _ *bot.Bot, update *models.Update) {

}

// COMMAND DELETE

func (tb *TgBot) deleteHandler(ctx context.Context, _ *bot.Bot, update *models.Update) {

}
